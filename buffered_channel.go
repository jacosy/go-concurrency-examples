package main

import (
	"fmt"
	"time"
)

func unbufferedChannelExample() {
	// Unbuffered channel
	unbufferedChan := make(chan int)

	go func() {
		fmt.Println("Sending to unbuffered channel...")
		unbufferedChan <- 1 // Blocks until a receiver is ready
		// Interstingly, the following Println won't be executed
		fmt.Println("Sent to unbuffered channel")
	}()

	time.Sleep(time.Millisecond) // Give the goroutine time to start

	fmt.Println("Receiving from unbuffered channel...")
	<-unbufferedChan
	fmt.Println("Received from unbuffered channel")

	// Buffered channel
	bufferedChan := make(chan int, 1)

	fmt.Println("Sending to buffered channel...")
	bufferedChan <- 2 // Does not block because the buffer is not full
	fmt.Println("Sent to buffered channel")

	fmt.Println("Receiving from buffered channel...")
	<-bufferedChan
	fmt.Println("Received from buffered channel")
}
