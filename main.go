package main

import (
	"net"

	"github.com/pkg/browser"
)

var Clients = make(map[net.Conn]int)

func main() {
	r := setupRoutes()

	generateClientQR()

	r.Static("/public", "./public")
	r.LoadHTMLGlob("public/*.html")

	initKeyboardMod()

	browser.OpenURL("www.tinder.com")
	browser.OpenURL("http://127.0.0.1:5000")

	r.Run(":5000")
}
