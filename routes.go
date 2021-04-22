package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			log.Fatalf("There was an error initializing a WebSocket: %v", err)
		}

		go func() {
			defer conn.Close()

			for {
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Fatalf("There was an error getting client data: %v", err)
				}
				log.Println(int(msg[1]))
				if IsVote(int(msg[1])) {
					log.Println("Vote given....")
					Clients[conn] = int(msg[1])
					CheckVotingMap()
				} else {
					log.Println("Action key given...")
					executeKeyStroke(int(msg[1]))
				}
			}
		}()

	})

	r.GET("/voting", func(context *gin.Context) {
		context.HTML(200, "voting.html", nil)
	})

	r.GET("/exit", func(context *gin.Context) {
		go gracefulDown()

		context.HTML(200, "exit.html", nil)
	})

	return r
}

func gracefulDown() {
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
