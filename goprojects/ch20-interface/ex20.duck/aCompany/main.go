package main

import (
	"fmt"
	"goprojects/ex20.duck/bCompany"
	"goprojects/ex20.duck/cCompany"
	"goprojects/ex20.duck/dCompany"
	"time"
)

type Database interface {
	Get()
	Set()
}

// Adapter 패턴
// dCompany에서는 bCompany, cCompany와 다르게 Get(), Set()이 아닌 GetData(), SetData()라는 function을 제공한다.
// 그러므로, dCompany는 위의 Database라는 interface를 구현 할 수 없다. 이런경우 Adapter pattern을 통해서 해결 할 수 있다.
// 1. struct type의 Adapter를 생성하고 해당 struct의 field로 dCompany.DDatabase struct를 추가한다.
// 2. Adapter struct가 Get(), Set() method를 갖도록 생성한다.
// 3. Adapter의 method에서 이제 struct의 filed인 dCompany의 method에 접근 할 수 있고 dCompany의 GetData(), SetData()를 호출한다.
type DBAdapter struct {
	adaterDB *dCompany.DDatabase
}

func (ad *DBAdapter) Get() {
	ad.adaterDB.GetData()
}

func (ad *DBAdapter) Set() {
	ad.adaterDB.SetData()
}

func main() {
	Compare()
}

func Compare() {
	BDB := &bCompany.BDatabase{"BDB"}
	fmt.Println(TotalTime(BDB))

	CDB := &cCompany.CDatabase{"CDB"}
	fmt.Println(TotalTime(CDB))

	// Apater 패턴
	DDB := &dCompany.DDatabase{"DDB"}
	AdapterDB := &DBAdapter{DDB}
	fmt.Println(TotalTime(AdapterDB))
}

func TotalTime(db Database) time.Duration {
	start := time.Now()
	db.Get()
	db.Set()
	end := time.Now()

	return end.Sub(start)
}
