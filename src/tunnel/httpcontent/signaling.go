// +build httpcontent all
// +build !windows

package httpcontent


import (
	"log"
	"os"
	"os/signal"
	"syscall"
)


func (t *HttpContentTunnel)setupSignaling() error {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGUSR1)


	go func() {

		for {
			s := <-sigs
			if s == syscall.SIGUSR1 {
				log.Println("Reloading RULES")
				t.readRules()
			}
		}
	}()

	return nil
}





