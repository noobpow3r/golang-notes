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

// Entity examples
type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	specialField float64
	Entity
}

type Position struct {
	x, y int
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

	e := Entity{
		name:    "Tim",
		id:      "3333",
		version: "1.0",
	}

	fmt.Printf("Entity is: %+v\n", e)

	se := SpecialEntity{
		specialField: 34.5,
		Entity: Entity{
			name:    "Thors",
			id:      "245",
			version: "2.0",
			Position: Position{
				x: 300,
				y: 400,
			},
		},
	}

	fmt.Printf("Entity is: %+v\n", se)
	fmt.Printf("SpecialEntity is: %+v\n", se)
	fmt.Printf("Entity is: %+v\n", se.Entity)
	fmt.Printf("Entity name is: %+v\n", se.name)
	fmt.Printf("Entity id is: %+v\n", se.id)
	fmt.Printf("Entity version is: %+v\n", se.version)
	fmt.Printf("Entity position is: %+v\n", se.Position)
	fmt.Printf("Entity position x is: %+v\n", se.x)
	fmt.Printf("Entity position y is: %+v\n", se.y)

}
