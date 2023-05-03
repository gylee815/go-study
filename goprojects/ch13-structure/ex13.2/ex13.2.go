package main

import (
	"fmt"
)

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Level int
	Price int
}

func main() {
	user := User{"giyean.lee", "1", 32, 1}
	vip := VIPUser{user, 1001, 10}

	fmt.Println(vip)

	fmt.Printf("Name: %s, Level: %d\n", vip.Name, vip.User.Level)
	fmt.Printf("Name: %s, VIPLevel: %d\n", vip.Name, vip.Level)
}
