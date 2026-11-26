package mio

import (
	"context"
	"errors"
	"net"
	"net/netip"

	"github.com/orgmio/mio/client"
	"github.com/orgmio/mio/cmd"
	"github.com/orgmio/mio/config"
	"github.com/orgmio/mio/protocol"
	"github.com/orgmio/mio/server"
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/adapter/inbound"
	"github.com/sagernet/sing-box/adapter/outbound"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/bufio"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/sing/service"
	"github.com/spf13/cobra"
)

const protocolType = "mio"

func RegisterCommand(root *cobra.Command) {
	mio := &cobra.Command{
		Use:   protocolType,
		Short: "Mio utilities",
	}
	mio.AddCommand(&cobra.Command{
		Use:   "keygen",
		Short: "Generate a server key pair",
		Args:  cobra.NoArgs,
		RunE: func(*cobra.Command, []string) error {
			return cmd.Keygen()
		},
	})
	root.AddCommand(mio)
}

type inboundOptions struct {
	Listen     string `json:"listen,omitempty"`
	ListenPort uint16 `json:"listen_port"`
	Dest       string `json:"dest"`
	DestPort   uint16 `json:"dest_port,omitempty"`
	PrivateKey string `json:"private_key"`
}

type mioInbound struct {
	inbound.Adapter
	server *server.Server
	logger log.ContextLogger
}

func RegisterInbound(registry *inbound.Registry) {
	inbound.Register[inboundOptions](registry, protocolType, newInbound)
}

func newInbound(_ context.Context, _ adapter.Router, logger log.ContextLogger, tag string, options inboundOptions) (adapter.Inbound, error) {
	if options.ListenPort == 0 || options.Dest == "" {
		return nil, errors.New("mio requires listen_port and dest")
	}
	srv, err := server.New(config.Inbound{
		Listen:     options.Listen,
		ListenPort: int(options.ListenPort),
		Dest:       options.Dest,
		DestPort:   int(options.DestPort),
		PrivateKey: options.PrivateKey,
	})
	if err != nil {
		return nil, err
	}
	return &mioInbound{Adapter: inbound.NewAdapter(protocolType, tag), server: srv, logger: logger}, nil
}

func (h *mioInbound) Start(stage adapter.StartStage, scope *adapter.Scope) error {
	if stage != adapter.StartStateStart {
		return nil
	}
	ctx, cancel := context.WithCancel(scope.Context())
	if err := h.server.Start(ctx); err != nil {
		cancel()
		return err
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer cancel()
		if err := h.server.Serve(ctx); err != nil {
			h.logger.ErrorContext(ctx, "serve mio: ", err)
		}
	}()
	scope.Add(func() error {
		cancel()
		<-done
		return nil
	})
	return nil
}

type outboundOptions struct {
	option.ServerOptions
	Dest      string `json:"dest"`
	PublicKey string `json:"public_key"`
}

type mioOutbound struct {
	outbound.Adapter
	dialer       *client.Dialer
	dns          adapter.DNSRouter
	dnsTransport adapter.DNSTransportManager
}

func RegisterOutbound(registry *outbound.Registry) {
	outbound.Register[outboundOptions](registry, protocolType, newOutbound)
}

func newOutbound(ctx context.Context, _ adapter.Router, _ log.ContextLogger, tag string, options outboundOptions) (adapter.Outbound, error) {
	if options.Server == "" || options.ServerPort == 0 || options.Dest == "" {
		return nil, errors.New("mio requires server, server_port and dest")
	}
	key, err := protocol.ParsePublicKey(options.PublicKey)
	if err != nil {
		return nil, err
	}
	return &mioOutbound{
		Adapter:      outbound.NewAdapter(protocolType, tag, []string{N.NetworkTCP, N.NetworkUDP}, nil),
		dns:          service.FromContext[adapter.DNSRouter](ctx),
		dnsTransport: service.FromContext[adapter.DNSTransportManager](ctx),
		dialer: client.NewDialer(protocol.Peer{
			Server:     options.Server,
			ServerPort: int(options.ServerPort),
			Dest:       options.Dest,
			PublicKey:  key,
		}),
	}, nil
}

func (h *mioOutbound) Start(stage adapter.StartStage, scope *adapter.Scope) error {
	if stage == adapter.StartStateStart {
		scope.Add(func() error {
			h.dialer.Close()
			return nil
		})
	}
	return nil
}

func (h *mioOutbound) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	switch N.NetworkName(network) {
	case N.NetworkTCP:
		return h.dialer.Dial(ctx, destination.String())
	case N.NetworkUDP:
		conn, ip, err := h.ListenPacketWithDestination(ctx, destination)
		if err != nil {
			return nil, err
		}
		return bufio.NewBindPacketConn(conn, M.SocksaddrFrom(ip, destination.Port)), nil
	default:
		return nil, N.ErrUnknownNetwork
	}
}

func (h *mioOutbound) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	conn, ip, err := h.ListenPacketWithDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	if destination.IsDomain() {
		return bufio.NewNATPacketConn(bufio.NewPacketConn(conn), M.SocksaddrFrom(ip, destination.Port), destination), nil
	}
	return conn, nil
}

func (h *mioOutbound) ListenPacketWithDestination(ctx context.Context, destination M.Socksaddr) (net.PacketConn, netip.Addr, error) {
	ctx, metadata := adapter.ExtendContext(ctx)
	metadata.Outbound = h.Tag()
	metadata.Destination = destination
	ip, err := h.resolveUDP(ctx, destination)
	if err != nil {
		return nil, netip.Addr{}, err
	}
	conn, err := h.dialer.ListenUDP(ctx)
	if err != nil {
		return nil, netip.Addr{}, err
	}
	return &mioPacketConn{PacketConn: conn, realIP: h.realUDPAddress}, ip, nil
}

type mioPacketConn struct {
	net.PacketConn
	realIP func(netip.Addr) bool
}

func (c *mioPacketConn) WriteTo(p []byte, addr net.Addr) (int, error) {
	if addr == nil || !c.realIP(M.SocksaddrFromNet(addr).Addr) {
		return 0, errors.New("mio UDP packet requires a real destination IP")
	}
	return c.PacketConn.WriteTo(p, addr)
}

func (h *mioOutbound) resolveUDP(ctx context.Context, destination M.Socksaddr) (netip.Addr, error) {
	if destination.Port == 0 {
		return netip.Addr{}, errors.New("mio UDP requires a destination port")
	}
	if !destination.IsDomain() {
		if h.realUDPAddress(destination.Addr) {
			return destination.Addr.Unmap(), nil
		}
		return netip.Addr{}, errors.New("mio UDP requires a real destination IP")
	}
	if h.dns == nil {
		return netip.Addr{}, errors.New("mio UDP requires the sing-box DNS router for domains")
	}
	ips, err := h.dns.Lookup(ctx, destination.Fqdn, adapter.DNSQueryOptions{})
	if err != nil {
		return netip.Addr{}, err
	}
	for _, ip := range ips {
		if h.realUDPAddress(ip) {
			return ip.Unmap(), nil
		}
	}
	return netip.Addr{}, errors.New("mio UDP target has no real IP")
}

func (h *mioOutbound) realUDPAddress(ip netip.Addr) bool {
	ip = ip.Unmap()
	if !ip.IsValid() || ip.IsUnspecified() || ip.Zone() != "" {
		return false
	}
	if h.dnsTransport != nil {
		if fake := h.dnsTransport.FakeIP(); fake != nil && fake.Store().Contains(ip) {
			return false
		}
	}
	return true
}
