package main

import (
	"log"
	"time"
)

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()

	waitForChannels := 2 // ch1 & ch2
	for waitForChannels > 0 {
		select {
		case msg1 := <-ch1:
			waitForChannels -= 1
			log.Println("Received:", msg1, "waitForChannels:", waitForChannels)
		case msg2 := <-ch2:
			waitForChannels -= 1
			log.Println("Received:", msg2, "waitForChannels:", waitForChannels)
		default:
			time.Sleep(200 * time.Millisecond) // Sleep to avoid busy waiting
			log.Println("No messages received yet, waiting...", "waitForChannels:", waitForChannels)
		}
	}
	log.Println("Done with select example", "waitForChannels:", waitForChannels)
}
