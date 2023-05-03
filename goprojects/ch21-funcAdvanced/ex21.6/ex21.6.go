package main

import (
	"fmt"
)

// type CallFn func([]func())

func CaptureLoop(call func([]func())) {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	call(f)
}

func CaptureLoop2(call func([]func())) {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	call(f)
}

func main() {
	callFunction := func(f []func()) {
		for i := 0; i < 3; i++ {
			f[i]()
		}
	}
	CaptureLoop(callFunction)
	CaptureLoop2(callFunction)
}
