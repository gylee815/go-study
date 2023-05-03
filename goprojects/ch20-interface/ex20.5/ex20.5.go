package main

import (
	"fmt"
)

//empty interface -> fmt는 empty interface를 받는 함수이다.
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("type is int , value: %d\n", int(t))
	case float64:
		fmt.Printf("type is float64 , value: %f\n", float64(t))
	case string:
		fmt.Printf("type is string , value: %s\n", string(t))
	default:
		fmt.Printf("unsupported type : %T , value: %v\n", t, t)
	}
}

func PrintVal2(v interface{}) {
	switch t := fmt.Sprintf("%T", v); t {
	case "int":
		fmt.Printf("type is int , value: %v\n", v)
	case "float64":
		fmt.Printf("type is float64 , value: %v\n", v)
	case "string":
		fmt.Printf("type is string , value: %v\n", v)
	default:
		fmt.Printf("unsupported type: %T , value: %v\n", v, v)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(1.1)
	PrintVal("gylee")
	PrintVal(Student{20})

	PrintVal2(10)
	PrintVal2(1.1)
	PrintVal2("gylee")
	PrintVal2(Student{20})
}
