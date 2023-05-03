package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			num := i
			for {
				DepositAndWithdraw(account, num)
			}
			wg.Done()
		}()
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account, num int) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	fmt.Printf("%d's goroutine\n", num)
	account.Balance += 10000
	time.Sleep(time.Millisecond * 5000)
	account.Balance -= 10000

	fmt.Println(account.Balance)

}
