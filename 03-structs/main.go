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
	// player := Player{} // Empty struct
	player := Player{
		name:        "Captain Jack",
		health:      100,
		attackPower: 45.1,
	}

	fmt.Printf("This is the player: %v\n", player)
	fmt.Printf("This is the player: %+v\n", player)
	fmt.Printf("This is the player health: %d\n", getHealth(player))
	fmt.Printf("This is the player health: %d\n", player.getHealth())
	fmt.Printf("This is the player health: %d\n", player.health)
}
