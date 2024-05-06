package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("sum is", sum)
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("sum is now", sum)
	}
}
