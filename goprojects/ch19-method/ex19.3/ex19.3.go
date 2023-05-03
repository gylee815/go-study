package main

import (
	"fmt"
)

type Account struct {
	Balance   int
	FirstName string
	LastName  string
}

func (a *Account) withdrawPointer(amount int) {
	a.Balance -= amount
}

func (a Account) withdrawValue(amount int) {
	a.Balance -= amount
}

func (a Account) withdrawReturnValue(amount int) Account {
	a.Balance -= amount
	return a
}

func main() {
	myAccount := &Account{100, "GIYEAN", "LEE"}

	myAccount.withdrawPointer(30)
	fmt.Println("Call withdrawPointer Method")
	fmt.Printf("Reamin balance: %d\n", myAccount.Balance)

	myAccount.withdrawValue(30) // same with (*myAccount).withdrawValue(30)
	fmt.Println("Call withdrawValue Method")
	fmt.Printf("Reamin balance: %d\n", myAccount.Balance)

	*myAccount = myAccount.withdrawReturnValue(30)
	fmt.Println("Call withdrawReturnValue Method")
	fmt.Printf("Reamin balance: %d\n", myAccount.Balance)

	bAccount := Account{200, "GIBAEK", "LEE"}
	(&bAccount).withdrawPointer(30)
	fmt.Println("Call withdrawPointer Method")
	fmt.Printf("Reamin balance: %d\n", bAccount.Balance)

}
