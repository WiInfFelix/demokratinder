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

		Clients[conn] = 0

		log.Printf("New client connected... %d clients taking part", len(Clients))

		go func() {
			defer conn.Close()
			for {
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					//if connection is close remove from Client list and exit loop
					delete(Clients, conn)
					err = nil
					break
				}

				//if msg is ping, send back pong to client
				if string(msg) == "ping" {
					wsutil.WriteServerMessage(conn, ws.OpPong, []byte("pong"))
				}

				if msg != nil {

					log.Printf("Vote given: %v \n", int(msg[1]))
					if IsVote(int(msg[1])) {
						log.Println("Vote given....")
						Clients[conn] = int(msg[1])

						decisionMade := CheckVotingMap()

						//notiy all clients that vote was made
						if decisionMade {
							broadcastAll("reset_vote")
						}

					} else {
						log.Println("Action key given...")
						executeKeyStroke(int(msg[1]))
					}
				}

			}
			log.Printf("Client disconnected.... %d Clients remaining....", len(Clients))
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
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

func broadcastAll(msg string) {
	for x := range Clients {
		wsutil.WriteServerMessage(x, ws.OpText, []byte(msg))
	}
}
