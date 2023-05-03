package main

import (
	"container/list"
	"fmt"
)

const M = 10

func hash(d int) int {
	return d % M
}

type data struct {
	Key   int
	Value int
}

type hashMap struct {
	HashMap [M]*list.List
}

func (m *hashMap) insert(k int, v int) {
	d := data{k, v}
	if m.HashMap[hash(k)] == nil {
		m.HashMap[hash(k)] = list.New()
	}
	m.HashMap[hash(k)].PushBack(d)
}

func (m *hashMap) getVal(k int) (int, bool) {
	elements := m.HashMap[hash(k)]
	for e := elements.Front(); e != nil; e = e.Next() {
		if e.Value.(data).Key == k {
			return e.Value.(data).Value, true
		}
	}
	return 0, false
}

func (m *hashMap) remove(k int) (interface{}, bool) {
	elements := m.HashMap[hash(k)]
	for e := elements.Front(); e != nil; e = e.Next() {
		if e.Value.(data).Key == k {
			elements.Remove(e)
			return e.Value, true
		}
	}
	return nil, false
}

func NewHashMap(m [M]*list.List) *hashMap {
	newMap := &hashMap{
		HashMap: m,
	}
	return newMap
}

func main() {
	myMap := NewHashMap([M]*list.List{})
	myMap.insert(23, 10)
	myMap.insert(33, 20)

	myMap.insert(50, 100)
	myMap.insert(60, 100)

	for i := 0; i < M; i++ {
		if elements := myMap.HashMap[i]; elements != nil {
			for e := elements.Front(); e != nil; e = e.Next() {
				fmt.Printf("%d -> ", e.Value)
			}
		} else {
			continue
		}
		fmt.Println()
	}

	fmt.Println(myMap.getVal(33))
	fmt.Println(myMap.remove(63))
	fmt.Println(myMap.getVal(33))

	// m[hash(23)] = list.New()
	// m[hash(23)].PushBack(10)
	// m[hash(33)].PushBack(11)
	// m[hash(43)].PushBack(12)

	// m[hash(50)] = list.New()
	// m[hash(50)].PushBack(50)
	// m[hash(60)].PushBack(51)
	// m[hash(70)].PushBack(52)

	// fmt.Println("m[hash(23)]")
	// for e := m[hash(23)].Front(); e != nil; e = e.Next() {
	// 	fmt.Printf("%d -> ", e.Value)
	// }
	// fmt.Println()

	// fmt.Println("m[hash(50)]")
	// for e := m[hash(50)].Front(); e != nil; e = e.Next() {
	// 	fmt.Printf("%d -> ", e.Value)
	// }
	// fmt.Println()

}
