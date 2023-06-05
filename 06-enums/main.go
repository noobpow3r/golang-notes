package main

import "fmt"

type WeaponType int

// const (
// 	RailGun   WeaponType = 1
// 	Lightning WeaponType = 2
// 	Rocket    WeaponType = 3
// 	Gauntlet  WeaponType = 4
// )

// Using iota
const (
	RailGun WeaponType = iota // increment
	Lightning
	Rocket
	Gauntlet
)

func (w WeaponType) String() string {
	switch w {
	case RailGun:
		return "RAILGUN"
	case Lightning:
		return "LIGHTNING"
	case Rocket:
		return "ROCKET"
	case Gauntlet:
		return "GAUNTLET"
	default:
		return ""
	}
}

func getQuakeDamage(weaponType WeaponType) int {
	switch weaponType {
	case RailGun:
		return 100
	case Lightning:
		return 40
	case Rocket:
		return 80
	case Gauntlet:
		return 10
	default:
		panic("weapon does not exist")
	}
}

// func getQuakeDamage(weaponType string) int {
// 	switch weaponType {
// 	case "railGun":
// 		return 100
// 	case "lightning":
// 		return 40
// 	case "rocket":
// 		return 80
// 	case "gauntlet":
// 		return 10
// 	default:
// 		panic("weapon does not exist")
// 	}
// }

type Color int

type Position struct {
	x, y int
}

type Player struct {
	Position
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "PINK"
	default:
		panic("invalid color given")
	}
}

func (p Position) Move(val int) {
	fmt.Println("The position is moved by:", val)
}

func main() {
	// fmt.Println("damage of weapon:", getQuakeDamage("rocket"))
	// fmt.Println("damage of weapon:", getQuakeDamage(Rocket))
	fmt.Printf("damage of weapon (%d) (%d):\n", RailGun, getQuakeDamage(RailGun))
	fmt.Printf("damage of weapon (%d) (%d):\n", Lightning, getQuakeDamage(Lightning))
	fmt.Printf("damage of weapon (%d) (%d):\n", Rocket, getQuakeDamage(Rocket))
	fmt.Printf("damage of weapon (%d) (%d):\n", Gauntlet, getQuakeDamage(Gauntlet))

	// Using interface
	fmt.Printf("damage of weapon (%s) (%d):\n", RailGun, getQuakeDamage(RailGun))
	fmt.Printf("damage of weapon (%s) (%d):\n", Lightning, getQuakeDamage(Lightning))
	fmt.Printf("damage of weapon (%s) (%d):\n", Rocket, getQuakeDamage(Rocket))
	fmt.Printf("damage of weapon (%s) (%d):\n", Gauntlet, getQuakeDamage(Gauntlet))

	p := Player{}
	p.Move(100)
	fmt.Println("the color is:", ColorBlack)

}
