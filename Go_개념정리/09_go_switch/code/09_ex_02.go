package main

import "fmt"

type ColorType int

const (
	RED ColorType = iota
	BLUE
	GREEN
	YELLOW
)

func colorToString(color ColorType) string {
	switch color {
	case RED:
		return "RED"
	case BLUE:
		return "BLUE"
	case GREEN:
		return "GREEN"
	case YELLOW:
		return "YELLOW"
	default:
		return "UNKNOWN"
	}
}

func getMyColor() ColorType {
	return RED
}

func main() {
	fmt.Println("my fav color is", colorToString(getMyColor()))
}
