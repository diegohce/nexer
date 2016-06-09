package directtunnel

import (
	"tunnel"
)

type DirectTunnel struct {
	tunnel.BaseTunnel
}

func init() {
	tunnel.Register("direct", &DirectTunnel{})
}

func (t *DirectTunnel) New() *DirectTunnel {

	return &DirectTunnel{}

}

func (t *DirectTunnel) Setup(in chan string, out chan string) {
	t.InputChan = in
	t.OutputChan = out
}
