package main

import (
	"fmt"
	//	"net"
	"flag"
	_ "tunnel/directtunnel"
)

var config = `[[tunnel]]
listen="127.0.0.1:9001"
processor="tunnel"
args="target=127.0.0.1:9002"
`

//https://golang.org/pkg/net/#example_Listener

func main() {

	fmt.Println(config)

	/*
		// Listen on TCP port 2000 on all interfaces.
		l, err := net.Listen("tcp", ":2000")
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		for {
			// Wait for a connection.
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			// Handle the connection in a new goroutine.
			// The loop then returns to accepting, so that
			// multiple connections may be served concurrently.
			go func(c net.Conn) {
				// Echo all incoming data.
				io.Copy(c, c)
				// Shut down the connection.
				c.Close()
			}(conn)
		}
	*/
}
