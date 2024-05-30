package main

import "fmt"

// Enums

type WeaponType int

const (
	Axe WeaponType = iota // increments everything below
	Sword
	WoodStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodStick:
		return 10
	case Knife:
		return 70
	default:
		panic("weapon does not exist")
	}
}

// Advance enum

type Color int

// fmt string method
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "The First Color Is Blue"
	case ColorBlack:
		return "The Second Color Is Black"
	case ColorRed:
		return "The Third Color Is Red"
	case ColorYellow:
		return "The Fourth Color Is Yellow"
	case ColorWhite:
		return "The Fifth Color Is White"
	default:
		panic("Invalid color given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorRed
	ColorYellow
	ColorWhite
)

func enums() {
	fmt.Println("weapon", getDamage(Axe))
	fmt.Println(ColorBlack)
}
