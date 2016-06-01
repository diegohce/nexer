package main


var config := '''[[plexor]]
listen="127.0.0.1:9001"
processor="tunnel"
args="target=127.0.0.1:9002"
'''




import (
	"fmt"
)

type BaseProcessor struct {
	InputChan chan string
	OutputChan chan string
	Args map[string]string
}

type Tunnel struct {
	BaseProcessor
	ListenAddress string `toml:listen`
}

type (p *Tunnel) Setup( in, out chan string) {
	p.InputChan = in
	p.OutputChan = out
}


var processorsList := map[string]Processor{}




type Processor interface {

	Setup(in, out chan string)

}

func Register(name string, processor Processor) {

	processorList[name] = processor

}



func main() {
//https://golang.org/pkg/net/#example_Listener


}



