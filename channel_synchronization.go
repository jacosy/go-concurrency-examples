package main

import (
	"log"
	"time"
)

func channleSynchronizationExample() {
	done := make(chan bool)

	go channleSynchronizationWorker(done)

	log.Println("waiting for worker to finish on the main goroutine...")

	<-done // Block until the worker sends a signal
}

func channleSynchronizationWorker(done chan bool) {
	// Simulate work
	log.Println("worker is working...")
	time.Sleep(time.Second)
	log.Println("worker was done")
	done <- true // Signal that the work is done
}
