### grpc-loadbalance 

* grpc lb with etcd. 

* before run: 

``` 
# start a etcd 
docker run -it -p 2379:2379  walwong/etcd-amd64:3.2.18 etcd 
    --data-dir "./data" \ 
    --listen-client-urls 'http://0.0.0.0:2379' \ 
    --advertise-client-urls 'http://0.0.0.0:2379' 
    --debug

tips: the params addr must be 0.0.0.0 and can`t be localhost. 
```

* run 

```
start servers:
    go run server.go -p 5001
    go run server.go -p 5002
    go run server.go -p 5003
    go run server.go -p 5004
    go run server.go -p 5005
    go run server.go -p 5006

start a client:
    go run client.go
```

* outputs:

 ``` 
☁  lb [master] ⚡  go run client.go        
rpc error: code = Unavailable desc = there is no address available
{"time":"2019-03-21T13:41:04.116617+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"<nil> rpc error: code = Unavailable desc = there is no address available"}
<nil>
<nil>
{"time":"2019-03-21T13:41:05.121261+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5001\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:06.126386+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5002\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:07.130571+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5003\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:08.135061+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5004\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:09.138967+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5005\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:10.143493+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5006\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:11.148733+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5001\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:12.154+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5002\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:13.158553+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5003\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:14.16313+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5004\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:15.167441+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5005\"  <nil>"}
<nil>
{"time":"2019-03-21T13:41:16.172231+08:00","level":"INFO","prefix":"-","file":"client.go","line":"45","message":"addr:\":5006\"  <nil>"}
^Csignal: interrupt
```