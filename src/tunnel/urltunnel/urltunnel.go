package urltunnel

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"tunnel"
)

type UrlTunnel struct {
	tunnel.BaseTunnel
	Proto string
	Prod  string
	Debug string
}

func init() {
	tunnel.Register("url", &UrlTunnel{})
}

func (t *UrlTunnel) New() tunnel.Tunnel {

	return &UrlTunnel{}

}

func (t *UrlTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("url", flag.ExitOnError)

	fs.StringVar(&t.Prod, "prod", "", "Destination [address]:port")
	fs.StringVar(&t.Debug, "debug", "", "Destination [address]:port")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	if t.Prod == "" || t.Debug == "" {
		return errors.New("--prod and --debug params are required")
	}

	t.Proto = "tcp"

	return nil

}

func (t *UrlTunnel) ConnectionHandler(in_conn net.Conn) {

	var out_conn_addr string

	defer in_conn.Close()

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	query := in_req.URL.Query()

	if _, ok := query["debug"]; ok {
		out_conn_addr = t.Debug
	} else {
		out_conn_addr = t.Prod
	}

	log.Println(out_conn_addr)

	out_conn, err := net.Dial(t.Proto, out_conn_addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer out_conn.Close()

	in_req.Write(out_conn)
	io.Copy(in_conn, out_conn)

}
