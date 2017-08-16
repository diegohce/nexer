// +build connectionpool all

package connectionpool

import (
	"fmt"
	"time"
	"crypto/md5"
)

func (t *ConnectionPool) connectionHash() string {

	b := []byte(time.Now().String())
	s := fmt.Sprintf("%x", md5.Sum(b))

	return s
}

