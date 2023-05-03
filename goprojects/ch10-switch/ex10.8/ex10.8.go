package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func getMyFavoriteColor() ColorType {
	return Red
}

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func main() {
	color := getMyFavoriteColor()
	fmt.Printf("My favorite color is %s\n", colorToString(color))
}
