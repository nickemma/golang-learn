package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)
	fmt.Println(counter)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	replace(&phrase[1], "GoLang", &counter)
	fmt.Println(phrase)
	fmt.Println(counter)

}
