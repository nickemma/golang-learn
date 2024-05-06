package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	// Declare the variables
	var score1, score2, score3 int

	// Get the scores from the user
	fmt.Print("Enter score 1: ")
	fmt.Scan(&score1)
	fmt.Print("Enter score 2: ")
	fmt.Scan(&score2)
	fmt.Print("Enter score 3: ")
	fmt.Scan(&score3)

	// Calculate the average
	avg := average(score1, score2, score3)

	// Print the average
	fmt.Printf("The average is %.2f\n", avg)

	// Check if the average is greater than 70
	if avg >= 70 {
		fmt.Println("Congratulations! You passed the test.")
	} else {
		fmt.Println("Sorry! You failed the test.")
	}
}
