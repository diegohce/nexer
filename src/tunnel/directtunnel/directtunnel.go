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
	Delay int
}

func init() {
	tunnel.Register("direct", &DirectTunnel{})
}

func (t *DirectTunnel) New() tunnel.Tunnel {

	return &DirectTunnel{}

}

func (t *DirectTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("direct", flag.ExitOnError)

	fs.StringVar(&t.Dest, "dest", "", "Destination address:port")
	fs.StringVar(&t.Proto, "proto", "tcp", "Protocol (tcp/udp)")
	fs.IntVar(&t.Delay, "write-delay", 0, "Write delay in seconds")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	if t.Delay < 0 {
		log.Println("Write delay cannot be less than zero. Setting write-delay = 0")
		t.Delay = 0
	}

	return nil

}

func (t *DirectTunnel) ConnectionHandler(in_conn net.Conn) {

	out_conn, err := net.Dial(t.Proto, t.Dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer out_conn.Close()

	log.Println("Connection to", t.Dest, "established")

	dw := NewDelayedWriter(out_conn, t.Delay)

	io.Copy(dw, in_conn)
	//io.Copy(out_conn, in_conn)

	in_conn.Close()

	log.Println("Connection to", t.Dest, "Closed")

}



