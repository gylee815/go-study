package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	val := q.v.Front()
	if val != nil {
		return q.v.Remove(val)
	} else {
		return nil
	}
}

func (q *Queue) PrintQueue() {
	fmt.Println("Current Queue Elemets: ")
	val := q.v.Front()
	if val != nil {
		for e := val; e != nil; e = e.Next() {
			fmt.Printf("%v -> ", e.Value)
		}
	} else {
		fmt.Println("Queue is Empty")
	}

}

func main() {
	q := NewQueue()

	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	q.PrintQueue()
	fmt.Println()

	for e := q.Pop(); e != nil; e = q.Pop() {
		fmt.Printf("%v is pop\n", e)
	}
	q.PrintQueue()
}
