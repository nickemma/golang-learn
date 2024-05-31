package main

import (
	"sync"
)

type State struct {
	mu    sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count = i
}

func mutexes() {

}
