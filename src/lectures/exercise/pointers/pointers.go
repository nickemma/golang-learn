//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//  - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

//* Create a structure to store items and their security tag state
type securityTag bool

type item struct {
	name string
	tag  securityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *securityTag) {
	*tag = Active
}

func deactivate(tag *securityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []item) {
	fmt.Println("Checking out items...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirt := item{"Shirt", Active}
	pants := item{"Pants", Active}
	shoes := item{"Shoes", Active}
	hat := item{"Hat", Active}

	//  - Store them in a slice or array
	items := []item{shirt, pants, shoes, hat}
	fmt.Println("Initial items...", items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[1].tag)
	fmt.Println("item 1 Deactivated...", items)

	// - Activate any one security tag in the array/slice
	activate(&items[1].tag)
	fmt.Println("item 1 Activated...", items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)

	//  - Print out the array/slice after each change
	fmt.Println("Checked out...", items)
}
