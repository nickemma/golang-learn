package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanLiness(rooms [4]Room) {
	// when working with a for loop always make sure to make a copy of the array
	for i := 0; i < len(rooms); i++ {
		// make a copy of the room here always
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	// Create a shopping list array with enough room for 4 or more products
	myShoppingList := [...]string{"Milk", "Eggs", "Bread"}

	for i := 0; i < len(myShoppingList); i++ {
		item := myShoppingList[i]
		fmt.Println(item)
	}

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanLiness(rooms)
	fmt.Println("Done! All rooms checked.")

	rooms[1].cleaned = true
	rooms[3].cleaned = true

	checkCleanLiness(rooms)
	fmt.Println("Done! All rooms checked.", rooms)
}
