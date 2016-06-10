package directtunnel

import (
	"flag"
	"io"
	"log"
	"net"
	"tunnel"
)

type DirectTunnel struct {
	tunnel.BaseTunnel
	Proto string
	Dest  string
}

func init() {
	tunnel.Register("direct", &DirectTunnel{})
}

func (t *DirectTunnel) New() tunnel.Tunnel {

	return &DirectTunnel{}

}

func (t *DirectTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("direct", flag.ExitOnError)

	fs.StringVar(&t.Dest, "dest", "", "Destination [address]:port")
	fs.StringVar(&t.Proto, "proto", "tcp", "Protocol [tcp/udp]")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	return nil

}

func (t *DirectTunnel) ConnectionHandler(in_conn net.Conn) {

	out_conn, err := net.Dial(t.Proto, t.Dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer out_conn.Close()

	io.Copy(out_conn, in_conn)

	in_conn.Close()

}
