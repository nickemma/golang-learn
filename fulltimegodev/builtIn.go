package main

import "fmt"

// Builtin and Custom types

// - General types
// - Structs
// - Maps
// - Slices
// - Arrays
// - Custom types

// =========== GENERAL TYPES ============
var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	intVar32   int32   = 1
	intVar64   int64   = 45464
	intVar     int     = -1
	unitVar    uint    = 1
	unitVar32  uint32  = 1
	unitVar64  uint64  = 10
	unitVar8   uint8   = 0x2
	byteVar    byte    = 0x1
	runeVar    rune    = 'a'
	coolName   string  = "awesome"
)

// ============ STRUCTS ====================

type Player struct {
	name        string
	health      int
	attachPower float64
}

func getHealth2(player Player) int {
	return player.health
}

func (player Player) getHealth() int {
	return player.health
}

// ============ CUSTOM TYPES =================
type Weapon string

func getWeapon(w Weapon) string {
	return string(w)
}

func builtin() {
	player := Player{
		name:        "C. Ronaldo",
		health:      99,
		attachPower: 86.9,
	}
	fmt.Printf("player's health %d\n", getHealth2(player))
	fmt.Printf("player's health %d\n", player.getHealth())

	// ============ MAPS =================
	// users := map[string]int{}
	users := make(map[string]int)

	users["cool"] = 20
	users["nice"] = 10

	key := "cool" // the key you want to check
	age, ok := users[key]

	if !ok {
		fmt.Printf("User '%s' does not exist in the map\n", key)
	} else {
		fmt.Printf("User '%s' exists in the map with age %d\n", key, age)
	}

	// ============ SLICES =================
	number := []int{}
	number = append(number, 1, 2, 3, 4, 5)
	fmt.Println(number)

	// ============ ARRAYS =================
	arr := [3]int{}
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

}
