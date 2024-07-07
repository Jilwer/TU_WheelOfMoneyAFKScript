package main

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fpid, err := robotgo.FindIds("Tower-Win64-Shipping.exe")
	if err != nil {
		log.Fatal(err)
	}
	if len(fpid) > 0 {
		log.Println("Found process with pid:", fpid[0])
		err := scriptLoop(fpid[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Process not found")
	}
}

func scriptLoop(pid int) error {
	for {
		err := pressSpace(pid)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
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
