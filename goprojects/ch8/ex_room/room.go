package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	DiningRoom
)

func setLight(rooms uint8, room uint8) uint8 {
	return rooms | room
}

func resetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func isLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func turnLights(rooms uint8) {
	if isLightOn(rooms, MasterRoom) {
		fmt.Println("Turn on MasterRoom")
	}
	if isLightOn(rooms, LivingRoom) {
		fmt.Println("Turn on LivingRoom")
	}
	if isLightOn(rooms, BathRoom) {
		fmt.Println("Turn on BathRoom")
	}
	if isLightOn(rooms, DiningRoom) {
		fmt.Println("Turn on DiningRoom")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = setLight(rooms, MasterRoom)
	rooms = setLight(rooms, LivingRoom)
	rooms = setLight(rooms, DiningRoom)
	turnLights(rooms)
	fmt.Println()

	rooms = resetLight(rooms, LivingRoom)
	turnLights(rooms)
}
