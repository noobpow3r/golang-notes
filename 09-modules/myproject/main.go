package main

import (
	"fmt"
	"myproject/types"
	"myproject/util"
)

// Init module: go mod init <package_name>
// Command to build: go build -o <file_name>

func main() {
	user := types.User{
		Username: util.GetUsername(),
		Age:      util.GetAge(),
	}
	fmt.Printf("the user is %v:\n", user)
	fmt.Printf("the user is %+v:", user)
}
