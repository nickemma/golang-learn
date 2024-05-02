package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func double(x int) int {
	return x + x
}

func greet() {
	fmt.Println("Hello world from the greet function")
}

func main() {
	greet()

	dozen := double(6)

	result := add(dozen, 1)
	fmt.Println("the result is", result)

	anotherOne := add(double(6), 1)
	fmt.Println("another one is", anotherOne)
}
