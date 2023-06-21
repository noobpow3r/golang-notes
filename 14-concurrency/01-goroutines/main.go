package main

import (
	"fmt"
	"time"
)

func fetchResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime
	go say("world")
	say("hello")

	// msgch := make(chan string, 128)
	// msgch <- "A"
	// msgch <- "B"
	// msgch <- "C"
	// close(msgch)

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
	// for msg := range msgch {
	// 	fmt.Println("the message -> ", msg)
	// }

	// fmt.Println("reading all message from the channel")

}
