# influx2proxy

influx2proxy is a simple proxy that forwards incoming UDP packets received on a specific port to an InfluxDB 2.x instance. this is useful for embedded systems on LAN that either does not have internet access or for speed reasons or security reasons. Set this proxy on the same LAN as the embedded systems and configure them to send data to this proxy, then proxy will send the data to InfluxDB. Influx2proxy will take all the complexity of authentication and buffering off the embedded systems.

For now only UDP-server is implemented but soon TCP might be added. For now it listens on 0.0.0.0:<port>, i will add support for binding to specific interface later.

This library adds current time in microseconds to the received line protocol. So do not include timestamp in the line protocol if you are using this proxy.

## Speed
i am using UDP because it is faster than TCP and does not have the overhead of a TCP connection. Because there is no state, clients can send packets to the internet(through wifi or ethernet) and forget the packets. Theoretically, the UDP packets can be lost but in practice, it is very reliable. Theoretically the packets should be sent with same latency to LAN or WAN peers.

## run with arguments
```
go run *.go -udpport 18086 -url http://127.0.0.1:8086 -token "my-token" -org "my-org" -bucket "my-bucket"

where -url, -token, -org, -bucket are mandatory arguments and belong to InfluxDB 2.x instance.
```

## run with env variables
create a .env file inside the same folder as the executable with the following:
```env
INFLUX2_URL=http://127.0.0.1:8086
INFLUX2_TOKEN="my-token"
INFLUX2_ORG="my-org"
INFLUX2_BUCKET="my-bucket"
UDP_LISTEN_PORT=18086
```
then run:
```
go run *.go
```
