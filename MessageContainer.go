package main

type Message struct {
	MessageType   string `json:"msg_type"`
	DecisionTaken string `json:"decision_taken"`
	VotesGiven    int    `json:"votes_given"`
	Participants  int    `json:"participants"`
}
