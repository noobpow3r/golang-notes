package main

type CustomMap struct {
	data map[string]int
}

func (m *CustomMap) Insert(k string, v int) error {
	m.data[k] = v
	return nil
}

func NewCustomMap() *CustomMap {
	return &CustomMap{
		data: make(map[string]int),
		// data: map[string]int{},
	}
}

func main() {
}
