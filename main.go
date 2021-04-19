package main

import (
	"github.com/pkg/browser"
	"gopkg.in/olahol/melody.v1"
)

var Clients = make(map[*melody.Session]int)

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
