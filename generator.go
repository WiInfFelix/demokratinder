package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/skip2/go-qrcode"
)

func generateClientQR() {
	ip := getIPofHost()
	log.Println("Generating QR Code...")
	err := qrcode.WriteFile("http://"+ip+":5000/", qrcode.Medium, 512, "./public/qr.png")
	if err != nil {
		log.Fatal(err)
	}
}

func getIPofHost() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var currentIP string

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				currentIP = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
				break
			}

		}
	}

	return currentIP
}
