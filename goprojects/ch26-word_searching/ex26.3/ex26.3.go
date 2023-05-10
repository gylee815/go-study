package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 3 parameters are needed! ex) ex26.1 [word] [filepath]")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllfiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Printf("FileName: %s\n", findInfo.filename)
		fmt.Println("-----------------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-----------------------------------------------")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllfiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("Wrong file path !! err: ", err, "path: ", path)
		return findInfos
	}

	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInAllfile(word, filename))
	}
	return findInfos
}

func FindWordInAllfile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot find file!! ", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}

	return findInfo
}
