//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

//* Using a slice, create an assembly line that contains type Part
type Part string

//* Create a function to print out the contents of the assembly line
func printParts(lines []Part) {

	for i := 0; i < len(lines); i++ {
		part := lines[i]
		fmt.Println(part)
	}
}

func main() {
	//* Perform the following:
	//  - Create an assembly line having any three parts
	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Nut"

	fmt.Println("First 3 parts")
	printParts(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Screw", "Washer")
	fmt.Println("\nNewly added parts")
	printParts(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	newParts := assemblyLine[3:]
	fmt.Println("\nNewly sliced parts")
	//  - Print out the contents of the assembly line at each step
	printParts(newParts)
}
