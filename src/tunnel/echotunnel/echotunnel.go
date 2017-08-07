// +build echo all

package echotunnel

import (
	"io"
	"log"
	"net"
	"tunnel"
)

type EchoTunnel struct {
	tunnel.BaseTunnel
}

func init() {
	tunnel.Register("echo", &EchoTunnel{})
}

func (t *EchoTunnel) New() tunnel.Tunnel {

	return &EchoTunnel{}

}

func (t *EchoTunnel) Setup(tunnel_args []string) error {

	//	fs := flag.NewFlagSet("nexer", flag.ExitOnError)

	//	fs.StringVar(&address, "bind", "", "Bind [address]:port")
	//	fs.StringVar(&protocol, "proto", "tcp", "Protocol [tcp/udp]")
	//	fs.StringVar(&tunnel_type, "tunnel", "echo", "Tunnel type (see --tunnels)")
	//	fs.BoolVar(&list_tunnels, "tunnels", false, "Tunnels list")

	//	if len(nexer_args) == 0 {
	//		nexer_args = append(nexer_args, "--help")
	//	}

	//	err := fs.Parse(nexer_args)
	//	if err != nil {
	//		log.Println(err)
	//		os.Exit(1)
	//	}

	return nil

}

func (t *EchoTunnel) ConnectionHandler(in_conn net.Conn) {

	log.Println("Received connection from", t.In_Conn.RemoteAddr().String())

	io.Copy(in_conn, in_conn)

	in_conn.Close()

}
