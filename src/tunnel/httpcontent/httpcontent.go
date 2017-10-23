// +build httpcontent all

package httpcontent

import (
	"strings"
	"bufio"
	"bytes"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"tunnel"
)

type HttpContentTunnel struct {
	tunnel.BaseTunnel
	RulesFile string
	Rules [][]string
}

func init() {
	tunnel.Register("httpcontent", &HttpContentTunnel{})
}

func (t *HttpContentTunnel) New() tunnel.Tunnel {

	return &HttpContentTunnel{}

}

func (t *HttpContentTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("httpcontent", flag.ExitOnError)

	fs.StringVar(&t.RulesFile, "rules-file", "", "File with routing rules")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	if t.RulesFile == "" {
		return errors.New("--rules-file not specified")
	}

	if err := t.readRules(); err != nil {
		return err
	}

	t.setupSignaling()

	return nil

}

func (t *HttpContentTunnel) ConnectionHandler(in_conn net.Conn) {

	defer in_conn.Close()

	remote_addr := "[" + in_conn.RemoteAddr().String() + "]"

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(in_req.Body)
	b, _ := url.QueryUnescape(string(body))
	//log.Println("BODY ONE", string(b))

	in_req2, err := http.NewRequest(in_req.Method, in_req.URL.String(), bytes.NewReader(body))
	in_req2.Header = in_req.Header
	in_req2.Header.Set("X-Forwarded-For", strings.SplitN(remote_addr, ":", 2)[0])

	/*body2, _ := ioutil.ReadAll(in_req2.Body)
	b2, _ := url.QueryUnescape(string(body2))
	log.Println("BODY TWO", string(b2))*/

	//log.Println("HEADERS ONE", in_req.Header)
	//log.Println("HEADERS TWO", in_req2.Header)

	ws_function, ws_target, ws_terminalid, err := t.xmlParse(b)

	log.Println(remote_addr, "RULES", "FUNCTION", ws_function, "TARGET", ws_target, "TERMINALID", ws_terminalid)

	hostbyrule, err := t.getHostByRules(ws_function, ws_target, ws_terminalid)

	log.Println(remote_addr, "RULES", hostbyrule)

	log.Println(remote_addr, in_req2.URL)

	in_req2.Host = hostbyrule.host
	//log.Println(remote_addr, "Dest Host header value", in_req2.Host)

	if hostbyrule.rewrite != "" {
		in_req2.URL, _ = in_req2.URL.Parse(hostbyrule.rewrite)
		log.Println(remote_addr, "URL Changed to", in_req2.URL)
	}


	out_conn, err := net.Dial("tcp", hostbyrule.hostport)
	if err != nil {
		log.Println(remote_addr, err)
		//REMOVE!!!!!!
		//REMOVE!!!!!!
		//in_conn.Write([]byte("200 OK\r\n\r\n"))
		//REMOVE!!!!!!
		//REMOVE!!!!!!
		return
	}
	defer out_conn.Close()

	in_req2.Write(out_conn)
	io.Copy(in_conn, out_conn)

}







