package main

import "fmt"

func fibonnaci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Printf("Inside fibonnaci func, case channel c, value: %d\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Inside Goroutine for statement, index: %d\n", i)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonnaci(c, quit)
}
