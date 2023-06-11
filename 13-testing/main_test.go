package main

// Usefull command
// go test -v ./...
// go test -v ./... -run TestEqualPlayers
// without cache
// go test -v ./... -run TestEqualPlayers -count=1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	// t.Fail()
	fmt.Println("hello from our test")
	var (
		expected int = 10
		a        int = 3
		b        int = 5
	)
	have := CalculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
}

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Bruce",
		hp:   100,
	}
	have := Player{
		name: "Tim",
		hp:   50,
	}
	// have := Player{
	// 	name: "Bruce",
	// 	hp:   100,
	// }

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}

}
