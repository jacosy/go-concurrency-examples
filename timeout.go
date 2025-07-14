package main

import (
	"log"
	"time"
)

func timeoutExample() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "from channel 1"
	}()
	select {
	case msg := <-ch1:
		log.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		log.Println("Timeout occurred, no message received from channel 1")
	}

	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "from channel 2"
	}()
	select {
	case msg := <-ch2:
		log.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		log.Println("Timeout occurred, no message received from channel 2")
	}
}
