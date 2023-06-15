package main

import (
	"fmt"
	"time"
)

func fetchResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	// msg := <-msgch
	// fmt.Println("the msg is:", msg)
	// msg = <-msgch
	// fmt.Println("the msg is:", msg)
	// msg = <-msgch
	// fmt.Println("the msg is:", msg)

	// for {
	// 	msg, ok := <-msgch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("the message -> ", msg)
	// }

	// This piece of code is our consumer
	for msg := range msgch {
		fmt.Println("the message -> ", msg)
	}

	fmt.Println("reading all message from the channel")

}
