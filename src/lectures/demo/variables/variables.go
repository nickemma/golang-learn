package main

import "fmt"

func main() {
	var myName = "John Doe"
	fmt.Println("My name is", myName)

	var name string = "Techie Emma"
	fmt.Println("My name is", name)

	userName := "admin"
	fmt.Println("UserName =", userName)

	var sum int
	fmt.Println("the sum is", sum)

	var (
		firstInit = "mason"
		lastInit  = "jones"
	)
	fmt.Println("My initials are", firstInit, lastInit)
}
