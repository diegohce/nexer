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
  -version
    	Shows nexer current version
```

# Available tunnels
```
Available tunnel types:
	direct
	echo
	pip
	url
```

## echo tunnel
Echo has no arguments. Implements an echo server.

## direct tunnel

Direct tunnel redirection

```
Usage of direct:
  -dest string
    	Destination [address]:port
  -proto string
    	Protocol [tcp/udp] (default "tcp")
```
## pip tunnel

Tunnel to pip service (python packages index)

```
Usage of pip:
  -dest string
    	Destination pip servername (will use https port 443 always!) (default "pypi.python.org")
```

### Create the pip tunnel

To create a tunnel simply run:
```
# ./nexer -bind :3143 -tunnel pip
```

### Installing python packages using the pip tunnel

```
pip install --index-url "http://yourserver:3143/simple/" <python-package-name>
```

or

```
pip install --index-url "http://yourserver:3143/simple/" -r requirements.txt
```

### Searching for packages using the pip tunnel

```
pip search --index http://yourserver:3143/pypi <python-package-name>
```

## url tunnel
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
[echo tunnel](https://github.com/diegohce/nexer/blob/master/src/tunnel/echotunnel/echotunnel.go) and the 
[direct tunnel](https://github.com/diegohce/nexer/blob/master/src/tunnel/directtunnel/directtunnel.go) as an example.

Also check the ```import``` statement in [nexer.go](https://github.com/diegohce/nexer/blob/master/src/nexer.go)
where tunnels **must** be imported to be available.

## Building 

After setting Go environment values 
([goenv.sh](https://github.com/diegohce/nexer/blob/master/goenv.sh) might help), 
go to ```src``` directory and run from the command line:

```go build nexer.go version.go```





