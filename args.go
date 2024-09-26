package main

import (
	"log"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/joho/godotenv"
)

var (
	app                 = kingpin.New("influx2proxy", "InfluxDB2 Proxy")
	influxUrl           = app.Flag("url", "InfluxDB URL.").Required().String()
	influxToken         = app.Flag("token", "InfluxDB token.").Required().String()
	influxOrg           = app.Flag("org", "InfluxDB organization.").Required().String()
	influxBucket        = app.Flag("bucket", "InfluxDB bucket.").Required().String()
	udpListenPortString = app.Flag("udpport", "UDP listen port.").Default("18086").String()
)

func ParseEnv() bool {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
		return false
	}
	log.Println("loading .env file")
	*influxUrl = os.Getenv("INFLUX2_URL")
	*influxToken = os.Getenv("INFLUX2_TOKEN")
	*influxOrg = os.Getenv("INFLUX2_ORG")
	*influxBucket = os.Getenv("INFLUX2_BUCKET")
	*udpListenPortString = os.Getenv("UDP_LISTEN_PORT")
	log.Println(".env file loaded")
	return true
}

func ParseArgs() {
	log.Println("parsing args")
	kingpin.MustParse(app.Parse(os.Args[1:]))

}
