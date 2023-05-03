package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// Interface 선언
type WriterInterface interface {
	Write(string)
}

// Writer 변수가(function type 별칭) Write라는 method를 구현함
func (w Writer) Write(msg string) {
	// w에는 func(msg string){ fmt.Println(msg) }값이 들어있고
	// msg에는 "Hello World" 값이 들어있고
	// 최종적으로 func("Hello World")가 호출된다.
	w(msg)
}

func writeHello(writer WriterInterface) {
	// wirter interface는 Writer type의 주소를 갖고 있다.
	// Writer type이 implement한 Write함수 실행
	writer.Write("Hello World")

	// Writer라는 function type이 갖고 있는 method 라고 이해하면 되나요?
	writer.(Writer).Write("Hello World2")
}

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	// literal로 function type 변수 생성하고 Writer로 type convertion
	// Writer가 Write method를 구현하고 있기 때문에 interface 사용 가능
	writeHello(Writer(func(msg string) {
		fmt.Fprintln(f, msg)
	}))
}
