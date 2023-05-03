package bCompany

import (
	"fmt"
	"time"
)

type BDatabase struct {
	DBName string
}

func (db *BDatabase) Get() {
	fmt.Println("Get B Database")
	time.Sleep(time.Second * 1)
}

func (db *BDatabase) Set() {
	fmt.Println("Set B Database")
	time.Sleep(time.Second * 1)
}
