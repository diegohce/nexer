package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tunnel"
	_ "tunnel/directtunnel"
	_ "tunnel/echotunnel"
	_ "tunnel/urltunnel"
)

func split_args() ([]string, []string) {

	var i int
	var nexer_args []string
	var tunnel_args []string
	var sep string

	for i, sep = range os.Args[1:] {
		if sep == "--" {
			break
		}
	}

	if sep == "--" {
		nexer_args = os.Args[1 : i+1]
		tunnel_args = os.Args[i+2:]
	} else {
		nexer_args = os.Args[1:]
	}

	return nexer_args, tunnel_args

}

func main() {

	var address string
	var protocol string
	var tunnel_type string
	var list_tunnels bool

	nexer_args, tunnel_args := split_args()

	log.Println("nexer args:", nexer_args)
	log.Println("tunnel args:", tunnel_args)

	//	os.Exit(0)

	fs := flag.NewFlagSet("nexer", flag.ExitOnError)

	fs.StringVar(&address, "bind", "", "Bind [address]:port")
	fs.StringVar(&protocol, "proto", "tcp", "Protocol [tcp/udp]")
	fs.StringVar(&tunnel_type, "tunnel", "echo", "Tunnel type (see --tunnels)")
	fs.BoolVar(&list_tunnels, "tunnels", false, "Tunnels list")

	if len(nexer_args) == 0 {
		nexer_args = append(nexer_args, "--help")
	}

	err := fs.Parse(nexer_args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if list_tunnels {
		fmt.Println("Available tunnel types:")
		for _, t := range tunnel.TunnelsList() {
			fmt.Printf("\t%s\n", t)
		}
		os.Exit(0)
	}

	if address == "" {
		log.Println("Invalid address in --bind")
		os.Exit(1)
	}

	if err := tunnel.Listener(protocol, address, tunnel_type, tunnel_args); err != nil {
		log.Println(err)
	}

}
