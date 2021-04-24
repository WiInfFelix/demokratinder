package main

type Message struct {
	MessageType  string `json:"msg_type"`
	VotesGiven   int    `json:"votes_given"`
	Participants int    `json:"participants"`
}
