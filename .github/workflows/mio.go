package mio

import (
	"context"
	"errors"
	"net"

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
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
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
	dialer *client.Dialer
}

func RegisterOutbound(registry *outbound.Registry) {
	outbound.Register[outboundOptions](registry, protocolType, newOutbound)
}

func newOutbound(_ context.Context, _ adapter.Router, _ log.ContextLogger, tag string, options outboundOptions) (adapter.Outbound, error) {
	if options.Server == "" || options.ServerPort == 0 || options.Dest == "" {
		return nil, errors.New("mio requires server, server_port and dest")
	}
	key, err := protocol.ParsePublicKey(options.PublicKey)
	if err != nil {
		return nil, err
	}
	return &mioOutbound{
		Adapter: outbound.NewAdapter(protocolType, tag, []string{N.NetworkTCP}, nil),
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
	if N.NetworkName(network) != N.NetworkTCP {
		return nil, N.ErrUnknownNetwork
	}
	return h.dialer.Dial(ctx, destination.String())
}

func (h *mioOutbound) ListenPacket(context.Context, M.Socksaddr) (net.PacketConn, error) {
	return nil, N.ErrUnknownNetwork
}
