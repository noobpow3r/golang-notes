package main

import "fmt"

// Structs
type Player struct {
	name        string
	health      int
	attackPower float64
}

func getHealth(player Player) int {
	return player.health
}

func (player Player) getHealth() int {
	return player.health
}

func main() {
	users := map[string]int{
		"Bruce": 20,
		"Clark": 20,
	}

	fmt.Printf("users: %+v\n", users)

	cars := map[string]int{} // Empty map
	// cars := make(map[string]int)
	cars["Ferrari"] = 100
	cars["Tesla"] = 110

	fmt.Printf("cars: %+v\n", cars)

	// cost, ok := cars["Ferrari"] // Exist
	cost, ok := cars["Porsche"] // Not exist
	if !ok {
		fmt.Println("Porsche not exist in the map and value by default is:", cost)
	} else {
		fmt.Println("Porsche exists in the map:", cost)
	}

	cost = cars["Ferrari"]
	// Get a value from map
	fmt.Printf("Cost Ferrari: %d\n", cost)

	// Delete
	fmt.Printf("Delete Ferrari from cars map\n")
	delete(cars, "Ferrari")
	fmt.Printf("cars: %+v\n", cars)

	cars["Cupra"] = 120
	cars["Lamborghini"] = 200
	// Range over maps
	for k, v := range cars {
		fmt.Printf("The key %s and the value %d\n", k, v)
	}

}
