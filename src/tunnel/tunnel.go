package tunnel

import (
	"errors"
	"log"
	"net"
)

type BaseTunnel struct {
	In_Conn net.Conn
	Args    map[string]string
}

type Tunnel interface {
	ConnectionHandler(in_conn net.Conn)
}

var tunnelList = map[string]Tunnel{}

func Register(name string, t Tunnel) {

	tunnelList[name] = t

}

func GetTunnel(name string) Tunnel {

	return tunnelList[name]

}

func Listener(proto string, address string, tunnel_type string) error {

	t := GetTunnel(tunnel_type)
	if t == nil {
		return errors.New("Invalid tunnel type '" + tunnel_type + "'")
	}

	l, err := net.Listen(proto, address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go t.ConnectionHandler(conn)
	}

	return nil
}
