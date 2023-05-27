package main

import "fmt"

// General types
var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "Foo"
	intVar32   int32   = 1
	intVar64   int64   = 48484
	intVar     int     = -1
	uintVar    uint    = 1
	uint32Var  uint32  = 1
	uint64Var  uint64  = 10
	uint8Var   uint8   = 0x1
	byteVar    byte    = 0x2
	runVar     rune    = 'a'
)

func main() {
	fmt.Printf("The type of floatVar32 is %T and has a value of %v\n", floatVar32, floatVar32)
	fmt.Printf("The type of floatVar64 is %T and has a value of %v\n", floatVar64, floatVar64)
	fmt.Printf("The type of name is %T and has a value of %v\n", name, name)
	fmt.Printf("The type of intVar32 is %T and has a value of %v\n", intVar32, intVar32)
	fmt.Printf("The type of intVar64 is %T and has a value of %v\n", intVar64, intVar64)
	fmt.Printf("The type of intVar is %T and has a value of %v\n", intVar, intVar)
	fmt.Printf("The type of uintVar is %T and has a value of %v\n", uintVar, uintVar)
	fmt.Printf("The type of uint32Var is %T and has a value of %v\n", uint32Var, uint32Var)
	fmt.Printf("The type of uint64Var is %T and has a value of %v\n", uint64Var, uint64Var)
	fmt.Printf("The type of uint8Var is %T and has a value of %v\n", uint8Var, uint8Var)
	fmt.Printf("The type of byteVar is %T and has a value of %v\n", byteVar, byteVar)
	fmt.Printf("The type of runVar is %T and has a value of %v\n", runVar, runVar)
}
