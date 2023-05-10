package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 3 parameters are needed! ex) ex26.1 [word] [filepath]")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("Word to find: ", word)
	// fmt.Println("files: ", files)
	PrintAllFilesDirect(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	fmt.Println(files)
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Printf("Error occured. Err: %s, Path: %s\n", err, path)
		}
		fmt.Println("File list to find!")
		for _, name := range filelist {
			fmt.Printf("%s\n", name)
		}
	}

}

func PrintAllFilesDirect(files []string) {
	fmt.Println(files)
	fmt.Println("File list to find!")
	for i, file := range files {
		fmt.Printf("%d - %s\n", i, file)
	}

}
