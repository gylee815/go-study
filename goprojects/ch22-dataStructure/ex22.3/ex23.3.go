package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) PrintStack() {
	if s.v.Len() != 0 {
		fmt.Println("Current Stack Elemets: ")
		for e := s.v.Back(); e != nil; e = e.Prev() {
			fmt.Printf("%v -> ", e.Value)
		}
	} else {
		fmt.Println("Stack is Empty")
	}

}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	val := s.v.Back()
	if val != nil {
		return s.v.Remove(val)
	} else {
		return nil
	}
}

func main() {
	s := NewStack()

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	s.PrintStack()
	fmt.Println()

	for e := s.Pop(); e != nil; e = s.Pop() {
		fmt.Printf("%v is pop\n", e)
	}

	s.PrintStack()
}
