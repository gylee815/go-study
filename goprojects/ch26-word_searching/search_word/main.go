package main

import (
	"context"
	"fmt"
	search "goprojects/search_word/searching"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 3 parameters are needed! ex) ex26.1 [word] [filepath]")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(10)

	word := os.Args[1]
	files := os.Args[2:]
	cnt := 0
	searchInfo := search.NewSearchInfo(ctx, &wg, cancel, files, word)

	for _, file := range files {
		filelist, err := filepath.Glob(file)
		if err != nil {
			fmt.Println("Error occured!! Err: ", err)
		} else {
			cnt += len(filelist)
		}
	}

	fmt.Printf("file count: %d\n", cnt)

	go searchInfo.Start()
	go searchInfo.PrintResult(cnt)

	// fmt.Scanln()
	// cancel()

	wg.Wait()
}
