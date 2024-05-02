package main

import "fmt"

//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.

func greetPerson(name string) {
	fmt.Println("Hello and welcome", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func anyMessage() string {
	return "This is any message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers(x, y, z int) int {
	return x + y + z
}

//* Write a function that returns any number
func anyNumber() int {
	return 10
}

//* Write a function that returns any two numbers
func anyTwoNumbers() (int, int) {
	return 10, 20
}

func main() {
	greetPerson("John Doe")
	fmt.Println(anyMessage())

	//* Add three numbers together using any combination of the existing functions.
	a, b := anyTwoNumbers()
	answer := addThreeNumbers(anyNumber(), a, b)
	fmt.Println("The answer is", answer)

	//  * Print the result
	//* Call every function at least once
	x, y := anyTwoNumbers()
	z := anyNumber()
	result := addThreeNumbers(x, y, z)
	fmt.Println("The result is", result)

}
