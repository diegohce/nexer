// +build teetunnel all

package teetunnel

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"os"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"tunnel"
)

type TeeTunnel struct {
	tunnel.BaseTunnel
	Main      string
	ForwardTo string
	LogFile   string
	file      *os.File

}

func init() {
	tunnel.Register("tee", &TeeTunnel{})
}

func (t *TeeTunnel) New() tunnel.Tunnel {

	return &TeeTunnel{}

}

func (t *TeeTunnel) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("tee", flag.ExitOnError)

	fs.StringVar(&t.Main, "main", "", "Real request/endpoint destination")
	fs.StringVar(&t.ForwardTo, "forward-to", "", "Where to forward requests to")
	fs.StringVar(&t.LogFile, "logfile", "(stdout)", "Where to log the forwarded responses")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	if t.Main == "" {
		return errors.New("--main not specified")
	}
	if t.ForwardTo == "" {
		return errors.New("--forward-to not specified")
	}

	if err := t.openLogFile(); err != nil {	
		return err
	}

	return nil

}

func (t *TeeTunnel) ConnectionHandler(in_conn net.Conn) {

	defer in_conn.Close()

	remote_addr := "[" + in_conn.RemoteAddr().String() + "]"

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(in_req.Body)

	in_req2, err := http.NewRequest(in_req.Method, in_req.URL.String(), bytes.NewReader(body))
	in_req2.Header = in_req.Header


	main_conn, err := net.Dial("tcp", t.Main)
	if err != nil {
		log.Println(remote_addr, err)
		return
	}
	defer main_conn.Close()

	fwd_conn, err := net.Dial("tcp", t.ForwardTo)
	if err != nil {
		log.Println(remote_addr, err)
		return
	}
	defer fwd_conn.Close()

	go func() {
		in_req.Write(main_conn)
		io.Copy(in_conn, main_conn)
	}()

	in_req2.Write(fwd_conn)
	io.Copy(t.file, fwd_conn)

}


func (t *TeeTunnel) openLogFile()  error {

	if t.LogFile == "(stdout)" {
		t.file = os.Stdout
		return nil
	}

	f, err := os.OpenFile(t.LogFile, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	t.file = f

	return nil

}




