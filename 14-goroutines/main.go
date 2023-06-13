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
	// 1 unbuffered channel
	// 2 buffered channel
	resultch := make(chan string) // -> unbuffered channel

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "foo"

}
