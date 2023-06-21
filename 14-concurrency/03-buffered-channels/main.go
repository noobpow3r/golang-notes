package main

import "fmt"

func main() {
	// Channels can be buffered. Provide the buffer length as the second argument to 'make' to initialize a buffered channel
	// ch := make(chan int, 100)
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
