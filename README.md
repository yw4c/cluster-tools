# Cluster Tools
This is a troubleshooting box for validating cluster function including:

* gRPC Load Balance (Observing pod ID)
* [B3 Tracing](https://github.com/openzipkin/b3-propagation) : request ID, trace ID, span ID, etc.
* X-Forwarded-For and NAT Egress address
* Websocket
* HPA (Burn memory)
* PV (Generate file)
* MySQL, Redis connection
* Fault Injection (Mock internal error)

## Content
[In Container Tools](in-container-tools)

## In Container Tools
* mysql-client 
````sh
mysql -hmyhost --port 3306 -uuser -ppassword  --database mydb
````
* redis
````sh
redis-cli -h myhost -p 6379 ping
````

* grpcurl
````sh
grpcurl -plaintext myhost:9090 list
grpcurl -plaintext -rpc-header X-Request-ID:test -plaintext -d '{"ParamOne": "1", "ParamTwo": "1"}' localhost:8081 observe.ObserveService/GetStatus
grpcurl -rpc-header x-request-id:example-request-id -plaintext -d '{"ParamOne": "1", "ParamTwo": "1"}' myhost:7002 pingpong.PingPongService/PingPongEndpoint
````

* nc
````sh
nc myhost myport -v -z
````

# Service
* http 
```
curl http://localhost:8081/observe/info  --header "X-Request-ID: test"
```
* websocket: 
```
ws://your-host/cluster-tool/ping
```

