package main

import "fmt"

func printSlices(title string, slice []string) {
	fmt.Print()
	fmt.Println("_ _ _", title, "_ _ _")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println("elements here", element)
	}
}

func main() {
	route := []string{"Home", "About", "Projects"}

	printSlices("Initial Route", route)
	route = append(route, "Contact", "Services")

	printSlices("Updated Route", route)

	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")
	fmt.Println(route[2], "Visited")

	newRoute := route[3:5]
	printSlices("New Route", newRoute)
}
