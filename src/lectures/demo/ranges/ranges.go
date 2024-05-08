package main

import "fmt"

func main() {

	ranges := []string{"Hello", "World", "Golang", "Programming", "Language"}

	// Print the range using a for loop
	for i, value := range ranges {
		fmt.Println(i, value, ":")

		for _, char := range value {
			// Print the rune representation character in the string
			fmt.Printf("%q\n", char)
		}
	}
}
