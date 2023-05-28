package main

import "fmt"

func main() {
	numbers := []int{}
	// numbers := []int{1, 2, 3}
	otherNumbers := make([]int, 0)

	fmt.Println(numbers)
	fmt.Println(otherNumbers)

	numbers = append(numbers, 1)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
}
