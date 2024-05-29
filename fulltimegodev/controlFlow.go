package main

import "fmt"

// Control Structures
// - for loops
// - range
// - break / continue

func controlFlow() {

	numbers := []int{1, 2, 3, 4, 5}
	// ========== FOR LOOPS ===========
	for i := 0; i < len(numbers); i++ {
		fmt.Println("for loop ", i)
	}
	// ========== Range ===========

	names := []string{"Peter", "John", "Doe", "Jane"}

	for _, name := range names {
		fmt.Println("range ", name)
	}

	// ========== Break ===========
	alphabet := []string{"a", "b", "c", "d"}

	for _, alp := range alphabet {
		if alp == "b" {
			break
		}
	}
	fmt.Println("break")

	users := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 5,
		"e": 7,
		"f": 9,
	}

	for key, value := range users {
		fmt.Printf("key %s value %d\n", key, value)
	}
}
