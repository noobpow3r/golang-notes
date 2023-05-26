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

// Structs
type Player struct {
	name        string
	health      int
	attackPower float64
}

func getHealth(player Player) int {
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
}
