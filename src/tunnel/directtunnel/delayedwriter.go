// +build direct all

package directtunnel


import (
	"io"
	"time"
)


type DelayedWriter struct {
	delay int
	w     io.Writer
}


func NewDelayedWriter(w io.Writer, delay int) *DelayedWriter {

	return &DelayedWriter{delay: delay, w: w}
}


func (dw *DelayedWriter) Write(p []byte) (int, error) {

	time.Sleep(time.Duration(dw.delay)*time.Second)

	return dw.w.Write(p)
}




