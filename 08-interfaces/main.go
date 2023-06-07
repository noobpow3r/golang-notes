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

// Second Example Advanced interface
type Storage interface {
	Putter
	Get(id int) (any, error)
	// Put(id int, val any) error
}

type Putter interface {
	Put(id int, val any) error
}

type Server struct {
	store Storage
}

type FooStorage struct{}

func (s FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s FooStorage) Put(id int, val any) error {
	return nil
}

func updateValue(id int, val any, p Putter) error {
	// store := FooStorage{}
	// return store.Put(id, val)
	return p.Put(id, val)
}

type SimplePutter struct{}

func (s SimplePutter) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s BarStorage) Put(id int, val any) error {
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

	fmt.Printf("Numbers: %v\n\n", numbers)

	// Second example Advanced interface
	fmt.Println("Second example Advanced Interface")
	s := Server{
		store: FooStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "foo")

	sputter := SimplePutter{}
	updateValue(1, "bar", sputter)

}
