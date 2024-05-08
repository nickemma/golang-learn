package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 4
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	// delete an item from the map
	delete(shoppingList, "milk")
	fmt.Println("Deleting from the list", shoppingList)

	// Iterate over the map using a for loop
	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal not found in the list")
	} else {
		fmt.Println("Cereal found in the list", cereal)
	}

	totalItems := 0
	for _, quantity := range shoppingList {
		fmt.Println(quantity)
		totalItems += quantity
	}
}
