package main

import "fmt"

type Players struct {
	HP int
}

func TakeDamage(p *Players, amount int) {
	p.HP -= amount
	fmt.Println("Player is taking damage ->", p.HP)
}

func main() {
	soccer := Players{
		HP: 100,
	}

	TakeDamage(&soccer, 10)
}
