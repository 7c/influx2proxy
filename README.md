# influx2proxy

influx2proxy is a simple proxy that forwards UDP packets received on a specific port to an InfluxDB 2.x instance. this is useful for embedded systems on LAN that either does not have internet access or for speed reasons or security reasons. Set this proxy on the same LAN as the embedded systems and configure them to send data to this proxy, then proxy will send the data to InfluxDB. Influx2proxy will take all the complexity of authentication and buffering off the embedded systems.

for now UDP server is implemented but soon TCP might be added. For now it listens on 0.0.0.0:<port>, i will add support for binding to specific interface later.

## run with arguments
```
go run *.go -udpport 18086 -url http://127.0.0.1:8086 -token "my-token" -org "my-org" -bucket "my-bucket"
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
