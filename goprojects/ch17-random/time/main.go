package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	rand.Seed(t.UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", rand.Intn(100))
	}
	fmt.Println()
	var d time.Duration = time.Duration(1)
	time.Sleep(d)

	fmt.Println(getDurationSec(t))
	fmt.Println(getDurationDay(t))
}

func getDurationSec(t time.Time) time.Duration {
	// loc, _ := time.LoadLocation("Asia/Seoul")
	// const longForm = "Jan 2, 2006 at 3:04pm"
	// t1, err := time.ParseInLocation(longForm, "Sep 1, 2022 at 0:00am", loc)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t1, t1.Location(), t1.UTC())

	return time.Now().Sub(t)
}

func getDurationDay(t time.Time) time.Duration {
	const shortForm = "Jan-02-2006"
	t2, err := time.Parse(shortForm, "Sep-18-2022")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t2)
	return t2.Sub(t)
}
