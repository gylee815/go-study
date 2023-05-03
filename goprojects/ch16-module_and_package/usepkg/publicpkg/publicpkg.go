package publicpkg

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	score int
}

type pStudent struct {
	Name  string
	Age   int
	score int
}

func PrintStudent(student Student) {
	fmt.Println(student)
}

func printpStudent(student pStudent) {
	fmt.Println(student)
}
