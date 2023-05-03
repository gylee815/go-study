package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)
	wg.Add(3)

	go makeBody(tireCh)
	go intstallTire(tireCh, paintCh)
	go paintBody(paintCh)

	wg.Wait()
	fmt.Println("Close Car Factory")
}

func makeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			car := &Car{
				Body: "Hyundai IONIQ5",
			}
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}

}

func intstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		car.Tire = "Hankook Tire"
		time.Sleep(time.Second)
		paintCh <- car
	}
	close(paintCh)
	wg.Done()
}

func paintBody(paintCh chan *Car) {
	for car := range paintCh {
		car.Color = "Gray"
		time.Sleep(time.Second)
		// duration := time.Now().Sub(startTime)
		duration := time.Since(startTime)

		fmt.Printf("After %.2f Completed make car => Body: %s / Tire: %s / Color: %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
