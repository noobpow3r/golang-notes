package main

import "fmt"

// global variables
var x = 10
var name = "Foo" // infer string
var firstName string = "Foo"
var lastName string

// global variables encapsulate ()
var (
	country        = "Bar"
	address string = "Bar"
	phone   string
)

// constants
const oldVersion = 1

const (
	lastVersion = 1
	LastVersion = 1 // Export constant use uppercase first letter
	keyLen      = 10
	KeyLen      = 10 // Export constant use uppercase first letter
)

func main() {
	// local variables
	var newVersion int // default value is 0
	fmt.Println(newVersion)

	version := 1           // infer int
	otherVersion := "Bar"  // infer string
	anotherVersion := 10.1 // infer float32 a float64
	fmt.Println(version)
	fmt.Println(otherVersion)
	fmt.Println(anotherVersion)

	// constants
	fmt.Println(oldVersion)
	fmt.Println(lastVersion)
	fmt.Println(keyLen)

}
