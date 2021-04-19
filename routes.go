package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/voting", func(context *gin.Context) {
		context.HTML(200, "voting.html", nil)
	})

	r.GET("/exit", func(context *gin.Context) {
		go gracefulDown()

		context.HTML(200, "exit.html", nil)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		Clients[s] = int(msg[1])
		CheckVotingMap()
	})

	m.HandleConnect(func(session *melody.Session) {
		Clients[session] = 0
		log.Printf("Client connected! %d Clients taking part\n", len(Clients))
	})

	m.HandleDisconnect(func(session *melody.Session) {
		delete(Clients, session)
		log.Printf("Client disconnected: %v.... %d Clients remaining\n", *session, len(Clients))
	})

	return r
}

func gracefulDown() {
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
