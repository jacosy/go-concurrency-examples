package main

import "fmt"

func channelExample() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(s, c)

	// receive from c
	// NOTE: the order of receiving from the channel is from left to right
	x, y, z := <-c, <-c, <-c

	fmt.Println(x, y, z)
	fmt.Println("Total sum:", x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
