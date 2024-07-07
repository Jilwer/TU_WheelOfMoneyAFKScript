package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fpid, err := robotgo.FindIds("Tower-Win64-Shipping.exe")
	if err != nil {
		log.Fatal(err)
	}

	if len(fpid) > 0 {
		log.Println("Found process with pid:", fpid[0])
		var isRunning bool
		setupScriptHook(fpid[0], &isRunning)
	} else {
		log.Fatal("Process not found")
	}
}

func setupScriptHook(pid int, isRunning *bool) {
	fmt.Println("CTRL + S to start/stop the script")
	hook.Register(hook.KeyDown, []string{"s", "ctrl"}, func(e hook.Event) {
		*isRunning = !*isRunning
		if *isRunning {
			fmt.Println("Script started")
			go scriptLoop(pid, isRunning)
		} else {
			fmt.Println("Script stopped")
		}

		fmt.Println("Running:", *isRunning)
	})

	s := hook.Start()
	<-hook.Process(s)
}

func scriptLoop(pid int, isRunning *bool) {
	spaceTicker := time.NewTicker(1 * time.Second)
	afkTicker := time.NewTicker(1 * time.Minute)
	defer spaceTicker.Stop()
	defer afkTicker.Stop()

	for *isRunning {
		select {
		case <-spaceTicker.C:
			if err := pressSpace(pid); err != nil {
				log.Printf("Error pressing space: %v", err)
				return
			}
			log.Println("Pressed space")
		case <-afkTicker.C:
			if err := antiAfk(pid); err != nil {
				log.Printf("Error performing anti-afk: %v", err)
				return
			}
			log.Println("Performed anti-afk")
		}
	}
	log.Println("Script loop exited")
}

func pressSpace(pid int) error {
	err := robotgo.KeyTap("space", pid)
	if err != nil {
		return err
	}
	return nil
}

func antiAfk(pid int) error {

	return nil
}
