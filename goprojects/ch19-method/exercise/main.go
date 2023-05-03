package main

import (
	"fmt"
	"time"
)

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	Id    int
}

func NewProduct(n string, p int, id int) *Product {
	pdt := &Product{
		Name:  n,
		Price: p,
		Id:    id,
	}
	return pdt
}

func NewCourier(n string) *Courier {
	cour := &Courier{n}
	return cour
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func NewParcel(p *Product) *Parcel {
	par := &Parcel{
		Pdt:         p,
		ShippedTime: time.Now(),
	}
	return par
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	par := NewParcel(p)
	return par
}

func (par *Parcel) Delivered() *Product {
	par.DeliveredTime = time.Now()
	return par.Pdt
}

func main() {
	myProduct := NewProduct("rebok", 10000, 1)

	c := NewCourier("gyl")

	myParcel := c.SendProduct(myProduct)
	fmt.Printf("%s(%d) is sended at %v\n", myParcel.Pdt.Name, myParcel.Pdt.Id, myParcel.ShippedTime)

	time.Sleep(time.Second * 5)

	myDelivery := myParcel.Delivered()
	fmt.Printf("%s(%d) is delivered at %v\n", myParcel.Pdt.Name, myParcel.Pdt.Id, myParcel.DeliveredTime)
	fmt.Printf("%s(%d) has delivered correctly\n", myDelivery.Name, myDelivery.Id)
}
