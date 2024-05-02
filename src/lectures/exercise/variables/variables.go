package main

import "fmt"

func main() {
	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor = "white"
	fmt.Println("My favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	var birthYear, ageInYears = 1998, 26
	fmt.Println("I was born in", birthYear, "and I am", ageInYears, "years old")

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "Said"
		lastInitial  = "Developer"
	)
	fmt.Println("My initials are", firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	//   variable created earlier
	ageInDays := 365 * ageInYears
	fmt.Println("I am", ageInDays, "days old")
}
