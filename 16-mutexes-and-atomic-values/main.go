package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

type State struct {
	mu    sync.Mutex
	count int
	// count int32
}

func (s *State) setState(i int) {
	// atomic.AddInt32(&s.count, int32(i))
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
}

func main() {
	state := State{} // Init field with default value

	// for i := 0; i < 10; i++ {
	// 	state.count = i
	// }

	fmt.Println(state)

}
