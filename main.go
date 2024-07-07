package main

import (
	"log"

	"github.com/go-vgo/robotgo"
)

func main() {
	fpid, err := robotgo.FindIds("Tower-Win64-Shipping.exe")
	if err != nil {
		log.Fatal(err)
	}
	if len(fpid) > 0 {
		log.Println("Found process with pid:", fpid[0])
	} else {
		log.Fatal("Process not found")
	}
}

func pressSpace() {

}

func antiAfk() {

}
