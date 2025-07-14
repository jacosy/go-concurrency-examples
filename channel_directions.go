package main

import (
	"fmt"
	"log"
)

func ping(ch chan<- string, msg string) {
	log.Println("ping: Sending message to channel...")
	ch <- msg // Send a message to the channel
}

func pong(inCh <-chan string, outChan chan<- string) {
	log.Println("pong: Receiving message from channel...")
	msg := <-inCh  // Receive a message from the input channel
	outChan <- msg // Send it to the output channel
}

func channelDirectionsExample() {
	inCh := make(chan string, 1)
	outCh := make(chan string, 1)
	go ping(inCh, "Hello")
	go pong(inCh, outCh)
	msg := <-outCh
	fmt.Println(msg)
}
