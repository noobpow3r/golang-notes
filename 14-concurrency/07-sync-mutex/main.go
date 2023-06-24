package main

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

}
