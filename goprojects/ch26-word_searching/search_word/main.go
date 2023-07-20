package main

import (
	"context"
	"fmt"
	search "goprojects/search_word/searching"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 3 parameters are needed! ex) ex26.1 [word] [filepath]")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(6)

	word := os.Args[1]
	files := os.Args[2:]
	file := os.Args
	fmt.Println("file: ", file)
	fmt.Printf("files: %s\n", files)
	searchInfo := search.NewSearchInfo(ctx, &wg, cancel, files, word)

	cnt := searchInfo.GetFileCount()

	fmt.Printf("file count: %d\n", cnt)

	go searchInfo.Start()
	go searchInfo.PrintResult(cnt)

	// fmt.Scanln()
	// cancel()

	wg.Wait()
}
