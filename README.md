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
	tee
	connectionpool
	direct
	echo
	pip
	apt
	apt-experimental
	url
```

## tee tunnel

Sends the same request to both ```-main``` and ```-forward-to``` hosts. Responses from ```-forward-to``` host are logged into ```-logfile``` or to standard output.

Usage of tee:
  -forward-to string
    	Where to forward requests to
  -logfile string
    	Where to log the forwarded responses (default "(stdout)")
  -main string
    	Real request/endpoint destination

## connectionpool

Connection pooling between nexer and -dest
```
Usage of connectionpool:
  -dest string
    	Destination address:port
  -pool-size int
    	Connection pool size (default 1)
```

## echo tunnel
Echo has no arguments. Implements an echo server.

## direct tunnel

Direct tunnel redirection

```
Usage of direct:
  -dest string
    	Destination address:port
  -proto string
    	Protocol [tcp/udp] (default "tcp")
  -write-delay int
      	Write delay in seconds (default 0)

```

## apt / apt-experimental tunnel

The apt-experimental tunnel is expected to be faster and more reliable than the plain apt tunnel.

```
Usage of apt:
  -log-requests
    	Show http requests (default false)
```

To create the tunnel run:
```
# ./nexer -bind :3142 -tunnel apt
or
# ./nexer -bind :3142 -tunnel apt-experimental
```

### To use apt-get with the tunnel

```
# apt-get -o "Acquire::http::proxy=http://yourserver:3142" update
```

or

```
# apt-get -o "Acquire::http::proxy=http://yourserver:3142" install <package name>
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
    	Destination address:port
  -prod string
    	Destination address:port
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

```go build -tags all nexer.go version.go```





