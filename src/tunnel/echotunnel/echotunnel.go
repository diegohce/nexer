package echotunnel

import (
	"io"
	"net"
	"tunnel"
)

type EchoTunnel struct {
	tunnel.BaseTunnel
}

func init() {
	tunnel.Register("echo", &EchoTunnel{})
}

func (t *EchoTunnel) New() *EchoTunnel {

	return &EchoTunnel{}

}

func (t *EchoTunnel) ConnectionHandler(in_conn net.Conn) {
	t.In_Conn = in_conn

	io.Copy(in_conn, in_conn)

	in_conn.Close()

}
