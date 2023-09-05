package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID   int
	Name string
}

func main() {

	// Retrieve the user profile.
	u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile.
	fmt.Printf("%+v\n", *u)

}

func retrieveUser(name string) (*user, error) {

	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	var u user
	err = json.Unmarshal([]byte(r), &u)
	return &u, err

}

func getUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}
