#!/usr/bin/env bash
# Patch sing-box so it builds and runs on Go 1.27.

set -euo pipefail

root="$(cd "$(dirname "$0")" && pwd)"
cd "$root"

die() { echo "error: $*" >&2; exit 1; }

test -f go.mod && grep -q 'module github.com/sagernet/sing-box' go.mod \
  || die "run this from a sing-box source tree"

force_close="$root/transport/v2rayhttp/force_close.go"
test -f "$force_close" || die "missing $force_close"

cat > "$force_close" <<'EOF'
package v2rayhttp

import (
	"net/http"
	"reflect"

	E "github.com/sagernet/sing/common/exceptions"

	"golang.org/x/net/http2"
)

func ResetTransport(rawTransport http.RoundTripper) http.RoundTripper {
	switch transport := rawTransport.(type) {
	case *http.Transport:
		transport.CloseIdleConnections()
		return transport.Clone()
	case *http2.Transport:
		// Go 1.27+ x/net/http2 wraps net/http and no longer has
		// (*Transport).connPool; CloseIdleConnections is the public equivalent.
		transport.CloseIdleConnections()
		return transport
	default:
		panic(E.New("unknown transport type: ", reflect.TypeOf(transport)))
	}
}
EOF
echo "patched $force_close"

strip_tlsunsafeekm() {
	local f="$1"
	test -f "$f" || return 0
	if grep -q 'tlsunsafeekm' "$f"; then
		# drop ",tlsunsafeekm=1" or "tlsunsafeekm=1," or standalone
		sed -i \
			-e 's/,tlsunsafeekm=1//g' \
			-e 's/tlsunsafeekm=1,//g' \
			-e 's/tlsunsafeekm=1//g' \
			"$f"
		echo "patched $f"
	else
		echo "skip $f (already without tlsunsafeekm)"
	fi
}

strip_tlsunsafeekm "$root/release/LDFLAGS"
strip_tlsunsafeekm "$root/cmd/internal/build_shared/flags.go"
strip_tlsunsafeekm "$root/experimental/libbox/ffi.json"

echo
echo "done."
