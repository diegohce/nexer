package tunnel


type BaseTunnel struct {
	InputChan chan string
	OutputChan chan string
	Args map[string]string
}


type Tunnel interface {

	Setup(in, out chan string)

}


var tunnelList := map[string]Tunnel{}


func Register(name string, tunnel Tunnel) {

	tunnelList[name] = tunnel

}




