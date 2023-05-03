package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name    string
	Age     int
	Score   int
	PassAcc float64
}

type Players []Player

func (p Players) Len() int           { return len(p) }
func (p Players) Less(i, j int) bool { return p[i].Score < p[j].Score }
func (p Players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	p := []Player{
		{"Donkey Na", 13, 45, 78.4},
		{"Maengtae Oh", 16, 24, 67.4},
		{"Dongdo Oh", 18, 54, 50.8},
		{"Geumsan Hwang", 16, 36, 89.7},
	}

	sort.Sort(sort.Reverse(Players(p)))
	fmt.Println(p)
}
