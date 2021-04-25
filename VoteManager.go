package main

import "net"

var Clients = make(map[net.Conn]int)

//Method that returns the votes made and the number of total participants
func getVoteStatus() (votesMade int, Participants int) {
	for _, y := range Clients {
		if y != 0 {
			votesMade++
		}
	}

	return votesMade, len(Clients)
}
