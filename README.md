# Feel the SF LUV

To build the server:

```
$ cd backend
```

```
$ go install ./...
```

To run the server:

By default the server will install in your `go/bin` directory

```
$ ~/go/bin/luv_server
```

The server will be more configurable in the near future :) 
Currently by default the server works against the Polygon mainnet.
It can be configured to point elsewhere with an environment variable, e.g.:

```
$ CHAIN_URL='https://rpc-mumbai.maticvigil.com' ~/go/bin/luv_server
```