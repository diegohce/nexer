// +build apt all

package apttunnel

import (
	"bufio"
	//	"errors"
	"crypto/tls"
	"flag"
	"fmt"
	//"io"
	"log"
	"net"
	"net/http"
	"strings"
	"tunnel"
)

type AptTunnelExp struct {
	tunnel.BaseTunnel
	Dest        string
	Proto       string
	LogRequests bool
}

func init() {
	tunnel.Register("apt-experimental", &AptTunnelExp{})
}

func (t *AptTunnelExp) New() tunnel.Tunnel {

	return &AptTunnelExp{}

}

func (t *AptTunnelExp) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("apt-experimental", flag.ExitOnError)

	fs.BoolVar(&t.LogRequests, "log-requests", false, "Show http requests")

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

func (t *AptTunnelExp) ConnectionHandler(in_conn net.Conn) {

	defer in_conn.Close()
	var out_conn net.Conn

	in_buffer := bufio.NewReader(in_conn)
	in_req, err := http.ReadRequest(in_buffer)
	if err != nil {
		log.Println(err)
		return
	}

	t.Dest = in_req.Host

	if t.LogRequests {
		log.Printf("%+v\n", in_req)
	}

	if strings.Index(t.Dest, ":443") > -1 {

		log.Println("Connecting with", t.Dest, "with TLS")

		out_conn, err = tls.Dial(t.Proto, fmt.Sprintf("%s", t.Dest), &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			log.Println(err)
			return
		}

	} else {

		log.Println("Connecting with", t.Dest)

		out_conn, err = net.Dial(t.Proto, fmt.Sprintf("%s:80", t.Dest))
		if err != nil {
			log.Println(err)
			return
		}
	}
	defer out_conn.Close()
	log.Println("Connected to", t.Dest)

	log.Println("Sending request to", t.Dest)
	in_req.Write(out_conn)
	log.Println("Request sent to", t.Dest)

	for {

		log.Println("--------", "Into loop")

		response, err := http.ReadResponse(bufio.NewReader(out_conn), in_req)
		if err != nil {
			log.Println(err)
			break
		}
		response.Write(in_conn)

		in_buffer := bufio.NewReader(in_conn)
		in_req, err := http.ReadRequest(in_buffer)
		if err != nil {
			log.Println("ReadRequest(2)", err)
			break
		}

		if t.LogRequests {
			log.Printf("%+v\n", in_req)
		}

		log.Println("Sending request to", t.Dest)
		in_req.Write(out_conn)
		log.Println("Request sent to", t.Dest)

	}

	log.Println("Closing connection with", t.Dest)

}
