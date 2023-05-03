package cCompany

import (
	"fmt"
	"time"
)

type CDatabase struct {
	DBName string
}

func (db *CDatabase) Get() {
	fmt.Println("Get C Database")
	time.Sleep(time.Second * 2)
}

func (db *CDatabase) Set() {
	fmt.Println("Set C Database")
	time.Sleep(time.Second * 2)
}
