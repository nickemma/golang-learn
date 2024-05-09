package main

import "fmt"

type Preparer interface {
	Prepare()
}

type Chicken string
type Salad string

func (c Chicken) Prepare() {
	fmt.Println("Wash the chicken")
	fmt.Println("Season the chicken")
	fmt.Println("Cook the chicken")
}

func (s Salad) Prepare() {
	fmt.Println("Wash the salad")
	fmt.Println("Cut the salad")
	fmt.Println("Toss and dress the salad")
}

func PrepareMeal(dishes []Preparer) {
	fmt.Println("Preparing a meal:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--- Dish %v --- \n", dish)
		dish.Prepare()
	}
	fmt.Println()
	fmt.Println("Meal is ready!")
}

func main() {
	dishes := []Preparer{
		Chicken("chicken wings"),
		Salad("salad bowl"),
	}
	PrepareMeal(dishes)
}
