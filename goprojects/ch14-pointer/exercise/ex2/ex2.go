package ex2

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func PrintEx2() {
	actor := NewActor("Gold Rabbit", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
