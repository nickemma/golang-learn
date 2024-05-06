package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	said := Passenger{"SAID", 2, false}
	fmt.Println(said)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 3}
		ella = Passenger{Name: "Ella", TicketNumber: 5}
	)
	fmt.Println(bill, ella)

	var techie Passenger
	techie.Name = "Emma"
	techie.TicketNumber = 4
	fmt.Println(techie)

	said.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println(bill.Name, "has boarded the bus")
	}
	if said.Boarded {
		fmt.Println(said.Name, " has boarded the bus")
	}

	techie.Boarded = true
	bus := Bus{techie}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "Is the one on the front seat")
}
