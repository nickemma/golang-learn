package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func generics() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("Cool", 1)
	m1.Insert("Awesome", 2)
}
