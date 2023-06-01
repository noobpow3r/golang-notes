package main

import "fmt"

// Init module: go mod init <package_name>
// Command to build: go build -o <file_name>

func main() {
	number := getNumber()
  user := User{
    username: "Bruce",
    age: getNumber(),
  }
	fmt.Println("the number is:", number)
	fmt.Println("the user is:", user)
}
