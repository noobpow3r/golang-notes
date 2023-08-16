package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// To get the address of a value, use the & operator.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass a copy of the "value of" count (what's in the box)
	// to the increment1 function.
	increment1(count)

	// Print out the "value of" and "address of" count.
	// The value of count will not change after the function call.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

}

// increment1 declares the function to accept its own copy of
// and integer value.
func increment1(inc int) {

	// Increment the local copy of the caller's int value.
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")

}
