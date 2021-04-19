package main

import (
	"fmt"
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
		voteCheck[y] += 1
	}

	for _, y := range voteCheck {
		fmt.Println(y)
	}

	_, found := voteCheck[0]

	if found {
		return
	} else {
		MakeKeyDecision(voteCheck)
	}

}

func MakeKeyDecision(check map[int]int) {

	var winnerKey, winnerValue int

	for key, value := range check {
		if value == 0 {
			return
		}
		if winnerValue < value {
			winnerKey = key
			winnerValue = value
		}
	}

	log.Printf("Vote %v won with %d number of votes", winnerKey, winnerValue)

	executeKeyStroke(winnerValue)
}

func executeKeyStroke(value int) {
	log.Println("Deciding on Key")

	switch value {
	case 1:
		kb.SetKeys(keybd_event.VK_RIGHT)
	case 2:
		kb.SetKeys(keybd_event.VK_LEFT)
	case 3:
		kb.SetKeys(keybd_event.VK_RIGHT)
	case 4:
		kb.SetKeys(keybd_event.VK_UP)
	case 5:
		kb.SetKeys(keybd_event.VK_DOWN)
	}

	log.Printf("Pressing key %v...", value)
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
