//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//

package main

import "fmt"

//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

// Create a shopping list array with enough room for 4 or more products
type ShoppingList struct {
	//  - Products must include the price and the name
	price int
	name  string
}

func productStat(list [4]ShoppingList) {
	var totalCost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		totalCost += item.price

		// Insert 3 products into the array
		if item.name != "" {
			totalItems++
		}
	}
	fmt.Println("The last item on the list: ", list[totalItems-1].name)
	fmt.Println("The total number of items: ", totalItems)
	fmt.Println("The total cost of the items: ", totalCost)
}
func main() {
	// Add a fourth product to the list
	products := [4]ShoppingList{
		{price: 10, name: "Milk"},
		{price: 20, name: "Eggs"},
		{price: 30, name: "Bread"},
	}
	productStat(products)

	// Add a fourth product to the list
	products[3] = ShoppingList{price: 40, name: "Butter"}

	// Print out the information again
	productStat(products)
	fmt.Println("Done! All products checked.", products)
}
