package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("squrae should be postive number. f: %g", f)
		return 0, errors.New("square should be positive number")
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-1.5)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("rst = %v\n", sqrt)
	}
}
