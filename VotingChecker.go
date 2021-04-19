package main

import (
	"github.com/micmonay/keybd_event"
	"log"
	"runtime"
	"time"
)

var kb keybd_event.KeyBonding

func CheckVotingMap() {
	log.Println("Checking map for decisions....")

	var voteCheck = make(map[int]int)

	for _, y := range Clients {
		voteCheck[y] = voteCheck[y] + 1
	}
	_, found := voteCheck[0]

	if found {
		return
	} else {
		MakeKeyDecision(voteCheck)
		for x := range Clients {
			Clients[x] = 0
		}
		log.Println("Votes reset...")
	}

}

func MakeKeyDecision(check map[int]int) {

	var winnerKey, winnerValue int

	for key, value := range check {

		if winnerValue < value {
			winnerKey = key
			winnerValue = value
		}
	}

	log.Printf("Vote %v won with %d number of votes", winnerKey, winnerValue)

	executeKeyStroke(winnerKey)
}

func executeKeyStroke(value int) {
	log.Println("Deciding on Key")
	var k string

	switch value {
	case 49:
		k = "right"
		kb.SetKeys(keybd_event.VK_RIGHT)
	case 50:
		k = "left"
		kb.SetKeys(keybd_event.VK_LEFT)
	case 51:
		k = "space"
		//The function loops if this is triggered on a desktop: can be disabled for desktop in future
		kb.SetKeys(keybd_event.VK_SPACE)
	case 52:
		k = "up"
		kb.SetKeys(keybd_event.VK_UP)
	case 53:
		k = "down"
		kb.SetKeys(keybd_event.VK_DOWN)
	}

	log.Printf("Pressing key %v...", k)
	err := kb.Launching()
	if err != nil {
		log.Fatalf("There was an error when pressing the keys: %v", err)
	}
}

func initKeyboardMod() {
	var err error
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatalf("There was an error initialising the keyboard module: %v", err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
}

func IsVote(message int) (isVote bool) {
	if message == 49 || message == 50 {
		isVote = true
	} else {
		isVote = false
	}
	return isVote
}
