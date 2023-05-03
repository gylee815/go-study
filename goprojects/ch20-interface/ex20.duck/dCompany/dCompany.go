package dCompany

import (
	"fmt"
	"time"
)

type DDatabase struct {
	DBName string
}

func (db *DDatabase) GetData() {
	fmt.Println("Get D Database")
	time.Sleep(time.Second * 3)
}

func (db *DDatabase) SetData() {
	fmt.Println("Set D Database")
	time.Sleep(time.Second * 3)
}
