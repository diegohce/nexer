package piptunnel

import (
	"bufio"
	//	"errors"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"tunnel"
)

type PipTunnel struct {
	tunnel.BaseTunnel
	Proto string
	Dest  string
}

func init() {
	tunnel.Register("pip", &PipTunnel{})
}

func (t *PipTunnel) New() tunnel.Tunnel {

	return &PipTunnel{}

}

func (t *PipTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("pip", flag.ExitOnError)

	fs.StringVar(&t.Dest, "dest", "pypi.python.org", "Destination pip servername")

	//	if len(tunnel_args) == 0 {
	//		tunnel_args = append(tunnel_args, "--help")
	//	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	t.Proto = "tcp"

	return nil

}

func (t *PipTunnel) ConnectionHandler(in_conn net.Conn) {

	//Search works like a charm!
	//TODO: Install reuses the connection and we're not expecting that here

	defer in_conn.Close()

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	in_req.Host = t.Dest

	//	log.Println(in_req)

	log.Println("Connecting to", t.Dest)

	out_conn, err := tls.Dial(t.Proto, fmt.Sprintf("%s:443", t.Dest), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatalln(err)
	}
	defer out_conn.Close()

	log.Println("Connected to", t.Dest)

	in_req.Write(out_conn)
	io.Copy(in_conn, out_conn)

	log.Println("Closing connection with", t.Dest)

}
