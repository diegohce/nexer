// +build direct all

package directtunnel

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
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

	defer in_conn.Close()

	out_conn, err := net.Dial(t.Proto, t.Dest)
	if err != nil {
		log.Println(err)
		return
	}
	defer out_conn.Close()

	log.Println("Connection to", t.Dest, "established")

	dw := NewDelayedWriter(out_conn, t.Delay)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(w *sync.WaitGroup) {
		defer w.Done()
		io.Copy(dw, in_conn)
		log.Println("Closing 1")
		out_conn.Close()
	}(wg)

	go func(w *sync.WaitGroup) {
		defer w.Done()
		io.Copy(in_conn, out_conn)
		log.Println("Closing 2")
		in_conn.Close()
	}(wg)

	wg.Wait()

	log.Println("Connection to", t.Dest, "Closed")

}



