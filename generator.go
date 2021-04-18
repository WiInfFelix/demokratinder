package main

import (
	"log"
	"net"

	"github.com/skip2/go-qrcode"
)

func generateClientQR() {
	ip := getIPofHost()
	log.Println("Generating QR Code...")
	err := qrcode.WriteFile("http://"+ip+":5000/voting", qrcode.Medium, 512, "./public/qr.png")
	if err != nil {
		log.Fatal(err)
	}
}

func getIPofHost() string {

	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	currentIP := conn.LocalAddr().(*net.UDPAddr)

	return currentIP.IP.String()
}
