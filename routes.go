package main

import (
	"encoding/json"
	"log"
	"net"
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
			log.Fatalf("There was an error initializing a WebSocket: %v \n", err)
		}

		Clients[conn] = 0

		log.Printf("New client connected... %d clients taking part \n", len(Clients))

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
					pongBack(conn)

				} else if msg != nil {

					log.Printf("Vote given: %v \n", int(msg[1]))

					if IsVote(int(msg[1])) {
						log.Println("Vote given....")
						Clients[conn] = int(msg[1])

						decisionMade, decisionTaken := CheckVotingMap()

						//notiy all clients that vote was made
						if decisionMade {
							broadcastAllReset(decisionTaken)
						} else {
							broadCastVoteStatus()
						}

					} else {
						log.Println("Action key given...")
						executeKeyStroke(int(msg[1]))
					}
				}

			}
			log.Printf("Client disconnected.... %d Clients remaining.... \n", len(Clients))
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

//Methode that shuts the server down after 2 seconds
func gracefulDown() {
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

//Method that broadcasts all clients that a decision was taken and votes have been reset
func broadcastAllReset(decisionTaken string) {
	resetMsg := &Message{
		MessageType:   "reset_vote",
		DecisionTaken: decisionTaken,
	}

	resetMsgJSON, err := json.Marshal(resetMsg)
	if err != nil {
		log.Fatalf("There was an error converting the reset broadcast: %v \n", err)
	}

	for x := range Clients {
		wsutil.WriteServerMessage(x, ws.OpText, resetMsgJSON)
	}
}

//Method that broadcasts the vote status of given votes to total participants
func broadCastVoteStatus() {
	votesMade, parts := getVoteStatus()

	statusMsg := &Message{
		MessageType:  "votemade",
		VotesGiven:   votesMade,
		Participants: parts,
	}

	statusMsgJSON, err := json.Marshal(statusMsg)
	if err != nil {
		log.Fatalf("There was an error getting the vote status: %v \n", err)
	}

	for x := range Clients {
		wsutil.WriteServerMessage(x, ws.OpText, statusMsgJSON)
	}
}

//Method that sends a JSON pong when server is pinged
func pongBack(conn net.Conn) {
	pongMsg := &Message{
		MessageType: "pong",
	}

	pongMsgJSON, err := json.Marshal(pongMsg)
	if err != nil {
		log.Fatalf("There was an error converting the pong to JSON: %v \n", err)
	}

	wsutil.WriteServerMessage(conn, ws.OpText, pongMsgJSON)
}
