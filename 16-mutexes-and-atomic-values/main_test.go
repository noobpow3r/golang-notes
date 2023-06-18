package main

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	state := State{}

	for i := 0; i < 10; i++ {
		state.count = i
	}

	fmt.Printf("%+v\n", state)

}
