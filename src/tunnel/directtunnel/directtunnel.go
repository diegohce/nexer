package directtunnel

import (
	"tunnel"
)

type DirectTunnel struct {
	tunnel.BaseTunnel
	ListenAddress string `toml:listen`
}

type (t *DirectTunnel) Setup( in, out chan string) {
	t.InputChan = in
	t.OutputChan = out
}
