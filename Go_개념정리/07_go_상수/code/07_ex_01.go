package main

import "fmt"

const (// bit flag 에 대한 설명
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}
func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}
func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}
func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("MasterRoom is on")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("LivingRoom is on")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("BathRoom is on")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("SmallRoom is on")
	}
}
func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)
	// fmt.Printf("rooms = %08b\n", rooms)
	rooms = ResetLight(rooms, BathRoom)
	// fmt.Printf("rooms = %08b\n", rooms)
	TurnLights(rooms)
}
