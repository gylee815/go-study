package main

import (
	"fmt"
)

type Attacker interface {
	Attack()
}

type Monster struct {
	Name   string
	Lv     int
	Damage int
	HP     int
}

func (m *Monster) Attack() {
	fmt.Printf("Monster Name: %s", m.Name)
	fmt.Println("Monster Attack")
}

type Hero struct {
	Name   string
	Lv     int
	Damage int
	HP     int
}

func (h *Hero) Attack() {
	fmt.Printf("HEro Name: %s\n", h.Name)
	fmt.Println("Hero Attack")
}

func HeroAttack(h Attacker, m Attacker) {
	h.Attack()
	hero := h.(*Hero)
	monster := m.(*Monster)
	fmt.Printf("%s attack %s. %d Damaged\n", hero.Name, monster.Name, hero.Damage)
}

func main() {
	myHero := &Hero{"gylee", 99, 500, 100000}
	myMonster := &Monster{"slime", 1, 1, 10}

	HeroAttack(myHero, myMonster)
}
