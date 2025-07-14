package main

import "log"

func nonBlockingChannelWithSelect() {
	msgCh := make(chan string)
	signalCh := make(chan struct{})

	select {
	case msg := <-msgCh:
		log.Println("Received message:", msg)
	default:
		log.Println("No message received, proceeding without blocking")
	}

	msg := "Hello, World!"
	select {
	case msgCh <- msg:
		log.Println("Message sent:", msg)
	default:
		log.Println("Channel is full, message not sent")
	}

	select {
	case msg := <-msgCh:
		log.Println("Received message after sending:", msg)
	case <-signalCh:
		log.Println("Received signal")
	default:
		log.Println("No message received or signal, proceeding without blocking")
	}
}
