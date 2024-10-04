package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func Influx2Writer(serverURL string, token string, org string, bucket string) (client influxdb2.Client, writeAPI api.WriteAPI) {
	clientOptions := influxdb2.DefaultOptions().
		SetPrecision(time.Microsecond). // second > millisecond > microsecond > nanosecond
		SetLogLevel(1)

	client = influxdb2.NewClientWithOptions(serverURL, token, clientOptions)
	writeAPI = client.WriteAPI(org, bucket)
	return
}

func DefaultInterface() (*net.Interface, error) {
	_, defaultRoute, err := net.ParseCIDR("0.0.0.0/0")
	if err != nil {
		return nil, err
	}

	// Get all network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Error getting network interfaces:", err)
		return nil, err
	}

	// Find the interface associated with the default route
	for _, iface := range ifaces {
		// lets skip the 'lo' interface
		if iface.Name == "lo" {
			continue
		}
		// fmt.Println("Interface:", iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("Error getting addresses for interface", iface.Name, ":", err)
			continue
		}
		for _, addr := range addrs {
			ip, _, _ := net.ParseCIDR(addr.String())
			// fmt.Println("IP:", ip, "Default route:", defaultRoute)
			if defaultRoute.Contains(ip) {
				// fmt.Println("Default route interface:", iface.Name)
				return &iface, nil
			}
		}
	}

	return nil, errors.New("default route interface not found")
}

func ListInterfaceIPs() ([]net.Interface, error) {
	color.Blue("Listening on:")
	_, err := DefaultInterface()
	if err != nil {
		return nil, err
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("iface: %s IP: %s\n", color.GreenString(iface.Name), color.YellowString(addr.String()))
		}
	}
	return interfaces, nil
}
