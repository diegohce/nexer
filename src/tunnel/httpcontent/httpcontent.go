package httpcontent

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	//"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"tunnel"
)

type HttpContentTunnel struct {
	tunnel.BaseTunnel
	Proto string
	Prod  string
	Debug string
}

func init() {
	tunnel.Register("httpcontent", &HttpContentTunnel{})
}

func (t *HttpContentTunnel) New() tunnel.Tunnel {

	return &HttpContentTunnel{}

}

func (t *HttpContentTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("httpcontent", flag.ExitOnError)

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

func (t *HttpContentTunnel) ConnectionHandler(in_conn net.Conn) {

	//var out_conn_addr string

	defer in_conn.Close()

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(in_req.Body)
	b, _ := url.QueryUnescape(string(body))
	log.Println("BODY ONE", string(b))

	in_req2, err := http.NewRequest(in_req.Method, in_req.URL.String(), bytes.NewReader(body))
	in_req2.Header = in_req.Header
	body2, _ := ioutil.ReadAll(in_req2.Body)

	b2, _ := url.QueryUnescape(string(body2))

	log.Println("BODY TWO", string(b2))

	ws_function, ws_target, ws_terminalid, err := xmlParse(b)

	log.Println("HEADERS ONE", in_req.Header)
	log.Println("HEADERS TWO", in_req2.Header)

	log.Println("FUNCTION", ws_function, "TARGET", ws_target, "TERMINALID", ws_terminalid)


	in_conn.Write([]byte("200 OK\r\n\r\n"))

	/*query := in_req.URL.Query()

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
	io.Copy(in_conn, out_conn)*/

}







