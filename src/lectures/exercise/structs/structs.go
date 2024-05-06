package main

import "fmt"

//--Requirements:
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides
type Coordinate struct {
	x, y int
}

//* Create a rectangle structure containing a length and width field
type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

//* Using functions, calculate the area and perimeter of a rectangle,
func length(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	// return 2 * (length(rect) + width(rect))
	return (width(rect) * 2) + (length(rect) * 2)
}

//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func printInfo(rect Rectangle) {
	fmt.Println("Rectangle: ", rect)
	fmt.Println("Area is: ", area(rect))
	fmt.Println("Perimeter is: ", perimeter(rect))
}

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	//  - Print the results to the terminal
	printInfo(rect)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.b.x *= 2
	rect.a.y *= 2
	printInfo(rect)
	//  - Print the new results to the terminal
	//
}
