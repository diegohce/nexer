// +build connectionpool all

package connectionpool

import (
	"flag"
	"io"
	"log"
	"net"
	"errors"
	"tunnel"
)

type ConnectionPool struct {
	tunnel.BaseTunnel
	Proto    string
	Dest     string
	PoolSize int
	Pool     chan int
}

func init() {
	tunnel.Register("connectionpool", &ConnectionPool{})
}

func (t *ConnectionPool) New() tunnel.Tunnel {

	return &ConnectionPool{}

}

func (t *ConnectionPool) Setup(tunnel_args []string) error {

	fs := flag.NewFlagSet("connectionpool", flag.ExitOnError)

	fs.StringVar(&t.Dest, "dest", "", "Destination address:port")
	fs.IntVar(&t.PoolSize, "pool-size", 1, "Connection pool size")

	if len(tunnel_args) == 0 {
		tunnel_args = append(tunnel_args, "--help")
	}

	err := fs.Parse(tunnel_args)
	if err != nil {
		return err
	}

	if t.Dest == "" {
		return errors.New("Missing -dest argument for connection pool")
	}

	if t.PoolSize < 1 {
		log.Println("Connection pool size cannot be less than one. Setting pool-size = 1")
		t.PoolSize = 1
	}

	t.Proto = "tcp"

	t.Pool = make(chan int, t.PoolSize)
	for i := 0; i < t.PoolSize; i++ {
		t.Pool <- i
	}

	return nil

}

func (t *ConnectionPool) ConnectionHandler(in_conn net.Conn) {

	chash := t.connectionHash()

	log.Println(chash, "Waiting for spot")

	spot := <-t.Pool

	log.Println("Pool state:", len(t.Pool) , "/", cap(t.Pool) )

	out_conn, err := net.Dial(t.Proto, t.Dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer out_conn.Close()

	log.Println(chash, spot, "Connection to", t.Dest, "established")

	go func() {
		io.Copy(in_conn, out_conn)
	}()

	io.Copy(out_conn, in_conn)

	in_conn.Close()

	log.Println(chash, spot, "Connection to", t.Dest, "Closed")

	log.Println(chash, spot, "Releasing spot")

	t.Pool <-spot

	log.Println("Pool state:", len(t.Pool) , "/", cap(t.Pool) )
}






