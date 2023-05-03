package main

import (
	"fmt"
	"goprojects/usepkg/custompkg"
	"goprojects/usepkg/exinit"
	"goprojects/usepkg/publicpkg"

	subpkg "goprojects/usepkg/publicpkg/subpkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	student := publicpkg.Student{
		Name: "gylee",
		Age:  32,
	}

	publicpkg.PrintPubpkg()
	publicpkg.PrintStudent(student)

	subpkg.PrintSub()

	exinit.PrintD()
}
