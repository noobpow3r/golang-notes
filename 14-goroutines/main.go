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
  // resultch <- " foo" // -> is no FULL -> IT WILL BLOCK -> BLOCK HERE

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "foo"

}
