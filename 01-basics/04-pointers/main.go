package main

import "fmt"

// First example

type Player struct {
	HP int
}

type Person struct {
	HP int
}

// Get a copy of the player
func TakeDamage(player Player, amount int) {
	player.HP -= amount
	fmt.Println("Player is taking damage: ", player.HP)
}

// Get a pointer of the person
func TakePower(person *Person, amount int) {
	person.HP -= amount
	fmt.Println("Person is taking damage: ", person.HP)
}

// Function receiver
func (p *Person) TakePower(amount int) {
	p.HP -= amount
	fmt.Println("Person is taking damage: ", p.HP)
}

// Second Example
type Database struct {
	user string
}

type Server struct {
	db *Database
}

func (s *Server) GetUserFromDB() string {
	if s.db == nil {
		// panic("database == nil. Is not initialized")
		return ""
	}
	return s.db.user
}

func main() {
	// First Example
	player := Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Println("Outside function Player is taking damage: ", player.HP)

	// person := Person{
	person := &Person{
		HP: 500,
	}
	// TakePower(&person, 50)
	// TakePower(person, 50)
	person.TakePower(50)
	fmt.Println("Outside function Person is taking damage: ", person.HP)

	// Second Example
	s := &Server{}
	s.GetUserFromDB()

}
