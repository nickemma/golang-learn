package main

import (
	"fmt"
)

type Stuff struct {
	Values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.Values) {
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.Values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
