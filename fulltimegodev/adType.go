package main

import "fmt"

// Advance struct Embedding

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string

	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func adType() {
	e := SpecialEntity{
		specialField: 88.0,
		Entity: Entity{
			name:    "John",
			id:      "id_1",
			version: "version 1.2",
			Position: Position{
				x: 100,
				y: 100,
			},
		},
	}

	fmt.Printf("the value is %+v\n", e)
}
