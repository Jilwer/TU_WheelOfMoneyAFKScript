package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
	"log"
	"math/rand/v2"
	"time"
)

var stopped bool

func main() {
	stopStartHotkey := hotkey.New(nil, hotkey.KeyF6)
	err := stopStartHotkey.Register()

	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}
	log.Printf("hotkey: %v is registered\n", stopStartHotkey)

	start()
	for {
		<-stopStartHotkey.Keydown()
		log.Printf("hotkey: %v is down\n", stopStartHotkey)
		if !stopped {
			stopped = true
			fmt.Println("Stopped")
		} else {
			// Introduce a delay here to prevent immediate restart
			time.Sleep(2 * time.Second)
			start()
			stopped = false
			fmt.Println("Started")
		}
	}
}

func start() {
	go spamSpace()
	go moveMouseRandom()
}

func spamSpace() {
	for {
		if !stopped {
			robotgo.KeyTap("space")
			fmt.Println("Pressed space")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func moveMouseRandom() {
	for {
		if !stopped {
			heightX, heightY := robotgo.GetScreenSize()
			fmt.Println("Screen size: ", heightX, heightY)
			screenWidth := rand.IntN(heightX)
			screenHeight := rand.IntN(heightY)
			fmt.Println("Random screen position: ", screenWidth, screenHeight)
			robotgo.MoveSmooth(screenWidth, screenHeight, 0.5, 1.0, 0)
			fmt.Println("Moved mouse")
			time.Sleep(5000 * time.Millisecond)
		}
	}
}
