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

func enums() {
	fmt.Println("weapon", getDamage(Axe))
}
