package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println("it:", i)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value", numbers[i])
	}

	// Slice for loop
	names := []string{"a", "b", "c", "d"}
	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, value := range names {
		if value == "b" {
			return
		}
		if value == "a" {
			break
		}
	}

	fmt.Println("break out of loop")

	// Map for loop
	users := map[string]int{
		"foo":   1,
		"bar":   2,
		"baz":   3,
		"Bruce": 33,
	}

	for key, value := range users {
		fmt.Printf("key %s value %d\n", key, value)
	}

	// Switch
	name := "foo"
	switch name {
	case "Bruce":
		fmt.Println("The name is Bruce")
	case "Drake":
		fmt.Println("The name is Drake")
	default:
		fmt.Println("The name default is", name)
	}

}
