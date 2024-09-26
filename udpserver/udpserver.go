package udpserver

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var logger *log.Logger = log.New(os.Stdout, color.CyanString("UDPSERVER "), log.LstdFlags)

type UDPServer struct {
	listenPort    int
	influx2Writer *api.WriteAPI
	udpConn       *net.UDPConn
}

func NewUDPServer(listenPort int, influx2Writer *api.WriteAPI) *UDPServer {
	return &UDPServer{
		listenPort:    listenPort,
		influx2Writer: influx2Writer,
	}
}

func (u *UDPServer) Start() {
	logger.Printf("Starting UDP server on port %d", u.listenPort)
	// start udp server
	udpAddr := &net.UDPAddr{
		Port: u.listenPort,
		IP:   net.ParseIP("0.0.0.0"),
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	u.udpConn = udpConn
	defer udpConn.Close()
	logger.Printf("UDP server started on port %d", u.listenPort)

	// read from udpConn
	buffer := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		logger.Printf("Received %d bytes from %s: '%s'", n, addr, buffer[:n])
		// send response
		udpConn.WriteToUDP([]byte("Received\n"), addr)
		// we expect line protocol
		// we need to split the buffer into lines
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			// u.influx2Writer.WritePoint(line)
			logger.Printf("received line: %s", line)
			(*u.influx2Writer).WriteRecord(line)
		}
	}
}

func (u *UDPServer) Stop() {
	logger.Println("Stopping UDP server")
	if u.udpConn != nil {
		logger.Println("Closing UDP connection")
		u.udpConn.Close()
		logger.Println("UDP connection closed")
	}
}
