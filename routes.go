package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "voting.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(string(msg))
		Clients[s] = int(msg[1])

	})

	m.HandleConnect(func(session *melody.Session) {
		Clients[session] = 0
	})

	m.HandleDisconnect(func(session *melody.Session) {
		delete(Clients, session)
	})

	return r
}
