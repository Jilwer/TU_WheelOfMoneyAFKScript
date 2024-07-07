package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand/v2"
	"time"
)

func main() {
	go spamSpace()
	go moveMouseRandom()
	go antiAFK()
	select {}
}

func spamSpace() {
	for {
		robotgo.KeyTap("space")
		fmt.Println("Pressed space")
		time.Sleep(500 * time.Millisecond)
	}
}

func moveMouseRandom() {
	for {
		// Generate random screen width and height
		heightX, heightY := robotgo.GetScreenSize()
		fmt.Println("Screen size: ", heightX, heightY)
		screenWidth := rand.IntN(heightX)
		screenHeight := rand.IntN(heightY)
		fmt.Println("Random screen position: ", screenWidth, screenHeight)

		// Move the mouse smoothly to the random position
		robotgo.MoveSmooth(screenWidth, screenHeight, 0.5, 1.0, 0)
		fmt.Println("Moved mouse")
		time.Sleep(5000 * time.Millisecond)
	}
}

func antiAFK() {
	for {
		// send every letter key besides y, x, and b
		var keys = []string{"a", " c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "z"}
		for _, key := range keys {
			robotgo.KeyTap(key)
			fmt.Println("Pressed key: ", key)
		}
		fmt.Println("Pressed all keys")
		time.Sleep(10 * time.Second)
	}
}
