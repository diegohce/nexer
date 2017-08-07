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

func (t *BaseTunnel) SetConn(conn net.Conn) {
	t.In_Conn = conn
}

type Tunnel interface {
	SetConn(net.Conn)
	Setup(tunnel_args []string) error
	ConnectionHandler(in_conn net.Conn)
	New() Tunnel
}

var tunnelList = map[string]Tunnel{}

func Register(name string, t Tunnel) {

	tunnelList[name] = t

}

func TunnelsList() []string {
	var l []string
	for t, _ := range tunnelList {
		l = append(l, t)
	}
	return l
}

func GetTunnel(name string) Tunnel {

	t := tunnelList[name]
	if t == nil {
		return nil
	}

	return t.New()

}

func Listener(proto string, address string, tunnel_type string, tunnel_args []string) error {

	t := GetTunnel(tunnel_type)
	if t == nil {
		return errors.New("Invalid tunnel type '" + tunnel_type + "'")
	}
	if err := t.Setup(tunnel_args); err != nil {
		return err
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
		t.SetConn(conn)
		go t.ConnectionHandler(conn)
	}

	return nil
}


