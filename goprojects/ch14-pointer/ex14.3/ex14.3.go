package main

import "fmt"

type Data struct {
	Value int
	Data  [200]int
}

func ChangeData(arg Data) {
	arg.Value = 100
	arg.Data[100] = 100
}

func ChangeData2(arg Data) Data {
	arg.Value = 100
	arg.Data[100] = 100
	return arg
}

func ChangeData3(arg *Data) {
	// fmt.Println(arg)
	fmt.Printf("Before changing data...\n")
	fmt.Printf("Value: %d, Data[100]: %d\n", arg.Value, arg.Data[100])
	(*arg).Value = 200     // eqauls to arg.Value
	(*arg).Data[100] = 200 // equals to arg.Data[100]
}

func main() {
	var data Data

	ChangeData(data)
	fmt.Println("call ChangeData")
	fmt.Printf("Value: %d, Data: %d\n", data.Value, data.Data[100])
	fmt.Println()

	data = ChangeData2(data)
	fmt.Println("call ChangeData2")
	fmt.Printf("Value: %d, Data: %d\n", data.Value, data.Data[100])
	fmt.Println()

	ChangeData3(&data)
	fmt.Println("call ChangeData3")
	fmt.Printf("Value: %d, Data: %d\n", data.Value, data.Data[100])
	fmt.Println()
}
