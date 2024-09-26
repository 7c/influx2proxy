package main

import (
	"context"
	"influx2proxy/udpserver"
	"log"
	"strconv"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var Influx2Client influxdb2.Client
var Influx2WriteAPI api.WriteAPI
var logDebug *log.Logger = log.Default()
var listenPort int

func main() {

	if !ParseEnv() {
		ParseArgs()
	}
	if _listenPort, err := strconv.Atoi(*udpListenPortString); err != nil {
		log.Fatal(err)
	} else {
		listenPort = _listenPort
		logDebug.Println("listen port :", listenPort)
	}
	logDebug.Println("influx2-url :", *influxUrl)
	logDebug.Println("influx2-token :", *influxToken)
	logDebug.Println("influx2-org :", *influxOrg)
	logDebug.Println("influx2-bucket :", *influxBucket)
	ListInterfaceIPs()

	Influx2Client, Influx2WriteAPI = Influx2Writer(*influxUrl, *influxToken, *influxOrg, *influxBucket)
	// ping influx2 server - this does not validate the token
	_, err := Influx2Client.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("pinging %s succeed", *influxUrl)

	udpserver := udpserver.NewUDPServer(listenPort, &Influx2WriteAPI)
	go udpserver.Start()

	select {}
}
