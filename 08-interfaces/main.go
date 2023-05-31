package main

import "fmt"

// Create interface
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

// Create struct that use interface
type ApiServer struct {
	numberStore NumberStorer
}

type MongoDBNumberStore struct {
	// fill
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

type PostgresNumberStore struct {
	// fill
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}

func (m PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the postgres storage")
	return nil
}

func main() {
	// apiServer := ApiServer{
	// 	numberStore: MongoDBNumberStore{},
	// }

	apiServer := ApiServer{
		numberStore: PostgresNumberStore{},
	}

	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(numbers)

}
