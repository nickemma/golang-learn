package main

// Advance Interface mechanism
// typed functions

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type Server struct {
	store Storage
}

func adInterface() {
	s := &Server{}
	s.store.Get(1)
	s.store.Put(2, "cool")
}
