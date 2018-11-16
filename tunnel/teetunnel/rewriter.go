// +build teetunnel all

package teetunnel


import (
	"log"
	"net"
)


type ReconnectWriter struct {
	remote string
	c      net.Conn
}

func (r *ReconnectWriter) Write(p []byte) (int, error) {

	if r.c == nil {
		var err error
		r.c, err = net.Dial("tcp", r.remote)
		if err != nil {
			log.Println("ReconnectWriter:", err)
			return len(p), nil
		}
	}

	if bc, err := r.c.Write(p); err != nil {
		log.Println("ReconnectWriter:", err)
		if c , cerr := net.Dial("tcp", r.remote); cerr != nil {
			log.Println("ReconnectWriter:", err)
			return len(p), nil
		} else {
			r.c = c
			bc, _ := r.c.Write(p)
			return bc, nil
		}
	} else {
		return bc, nil
	}
}



