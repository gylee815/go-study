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

type FileWriter struct {
	File *os.File
}

type StdioWriter struct {
}

func (f *FileWriter) Write(msg string) {
	fmt.Fprintln(f.File, msg)
}

func (s *StdioWriter) Write(msg string) {
	fmt.Println(msg)
}

func writeHello(writer WriterInterface) {
	writer.Write("Hello World")
}

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	w := &FileWriter{f}
	writeHello(w)

	s := &StdioWriter{}
	writeHello(s)
}
