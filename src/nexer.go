package main

import (
	"log"
	//	"flag"
	"tunnel"
	_ "tunnel/echotunnel"
)

func main() {

	if err := tunnel.Listener("tcp", ":2000", "echo"); err != nil {
		log.Println(err)
	}

}
