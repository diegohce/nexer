# nexer
An easily extensible and content based, network multiplexer.

# Usage 

## Nexer core arguments
```
Usage of nexer:
  -bind string
    	Bind [address]:port
  -proto string
    	Protocol [tcp/udp] (default "tcp")
  -tunnel string
    	Tunnel type (see --tunnels) (default "echo")
  -tunnels
    	Tunnels list
```

# Available tunnels
```
Available tunnel types:
	direct
	echo
	url
```

## echo
Echo has no arguments. Implements an echo server.

## direct

Direct tunnel redirection

```
Usage of direct:
  -dest string
    	Destination [address]:port
  -proto string
    	Protocol [tcp/udp] (default "tcp")
```

## url
```
Usage of url:
  -debug string
    	Destination [address]:port
  -prod string
    	Destination [address]:port
```
Redirects to -prod. If ```debug=1``` is present in the querystring, redirects to -debug

# Extending nexer (new tunnel dev)

The easiest way to commence is to look at the 
[echo tunnel](https://github.com/diegohce/nexer/blob/master/src/tunnel/echotunnel/echotunnel.go) and the [direct tunnel](https://github.com/diegohce/nexer/blob/master/src/tunnel/directtunnel/directtunnel.go) as an example.




