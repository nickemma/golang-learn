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
	name  string
	price float64
}

func main() {
	// Insert 3 products into the array
	products := [4]ShoppingList{
		{name: "Milk", price: 2.50},
		{name: "Eggs", price: 1.50},
		{name: "Bread", price: 1.00},
	}

	fmt.Println("The last item on the list: ", products[len(products)-2].name)
	fmt.Println("The total number of items: ", len(products))

	// The total cost of the items
	totalCost := 0.0
	for i := 0; i < len(products); i++ {
		totalCost += products[i].price
	}
	fmt.Println("The total cost of the items: ", totalCost)

	// Add a fourth product to the list
	products[3] = ShoppingList{name: "coffee", price: 1.25}

	// Recalculate the total cost of the items after adding the fourth product
	totalCost += products[3].price

	// Print out the information again
	fmt.Println("The last item on the list: ", products[len(products)-1].name)
	fmt.Println("The total number of items: ", len(products))
	fmt.Println("The total cost of the items: ", totalCost)
}
