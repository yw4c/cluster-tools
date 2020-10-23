# Tools
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
grpcurl -rpc-header x-request-id:example-request-id -plaintext -d '{"ParamOne": "1", "ParamTwo": "1"}' myhost:7002 pingpong.PingPongService/PingPongEndpoint
````

* nc
````sh
nc myhost myport -v -z
````