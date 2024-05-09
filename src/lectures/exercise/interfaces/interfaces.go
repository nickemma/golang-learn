//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//

package main

import (
	"fmt"
)

// --Requirements:
// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int // Lift type

// Define a lift picker interface
type LiftPicker interface {
	PickLift() Lift
}

type Motorcycles string
type Cars string
type Trucks string

// Define a function for all lifts
func (m Motorcycles) String() string {
	return fmt.Sprintf("Motorcycles: %v", string(m))
}

func (c Cars) String() string {
	return fmt.Sprintf("Cars: %v", string(c))
}

func (t Trucks) String() string {
	return fmt.Sprintf("Trucks: %v", string(t))
}

func (m Motorcycles) PickLift() Lift {
	return SmallLift
}
func (c Cars) PickLift() Lift {
	return StandardLift
}
func (t Trucks) PickLift() Lift {
	return LargeLift
}

//   - Write a single function to handle all of the vehicles
//     that the shop works on.
func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to Small Lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to Standard Lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to Large Lift\n", p)
	}
}

// func sendToLift(p LiftPicker) {
// 	switch p.PickLift() {
// 	case SmallLift:
// 		fmt.Printf("Send %s to Small Lift\n", p) // Utilize String() method here
// 	case StandardLift:
// 		fmt.Printf("Send %s to Standard Lift\n", p) // Utilize String() method here
// 	case LargeLift:
// 		fmt.Printf("Send %s to Large Lift\n", p) // Utilize String() method here
// 	}
// }

func main() {
	//* Vehicles have a model name in addition to the vehicle type:
	//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name

	car := Cars("Lamborghini")
	truck := Trucks("Crusher")
	motorcycle := Motorcycles("Harley")

	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	//
	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
