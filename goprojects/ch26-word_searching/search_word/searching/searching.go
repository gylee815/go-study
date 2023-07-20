package searching

import (
	"context"
	"fmt"
	find "goprojects/search_word/findWords"
	"sync"
)

type SearchInfo struct {
	Ctx    context.Context
	Wg     *sync.WaitGroup
	Cancel context.CancelFunc
	Files  []string
	Word   string
	InfoCh chan find.FindInfo
}

func NewSearchInfo(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc, files []string, word string) *SearchInfo {
	return &SearchInfo{
		Ctx:    ctx,
		Wg:     wg,
		Cancel: cancel,
		Files:  files,
		Word:   word,
		InfoCh: make(chan find.FindInfo),
	}
}

// func (s *SearchInfo) GetFileCount() int {
// 	var cnt int
// 	for _, file := range s.Files {
// 		filelist, err := filepath.Glob(file)
// 		if err != nil {
// 			fmt.Println("Error occured!! Err: ", err)
// 		} else {
// 			cnt += len(filelist)
// 		}
// 	}
// 	return cnt
// }

func (s *SearchInfo) GetFileCount() int {
	// var cnt int
	// fmt.Println(s.Files)
	// for _, file := range s.Files {
	// 	filelist, err := find.GetFileList(file)
	// 	if err != nil {
	// 		fmt.Println("Error occured!! Err: ", err)
	// 	} else {
	// 		cnt += len(filelist)
	// 	}
	// }
	filelist, err := find.GetFileList("*.txt")
	if err != nil {
		return 0
	}
	return len(filelist)
}

func (s *SearchInfo) Start() {
	for _, path := range s.Files {
		go find.FindWordInAllfiles(s.Word, path, s.Wg, s.InfoCh)
		// findInfos = append(findInfos, find.FindWordInAllfiles(word, path, ctx, cancel, wg)...)
	}
	s.Wg.Done()
}

func (s *SearchInfo) PrintResult(cnt int) {
	// after := time.After(time.Second * 10)
	for {
		select {
		case info := <-s.InfoCh:
			fmt.Printf("%d => FileName: %s\n", cnt, info.Filename)
			fmt.Println("-----------------------------------------------")
			for _, lineInfo := range info.Lines {
				fmt.Println("\t", lineInfo.LineNo, "\t", lineInfo.Line)
			}
			fmt.Println("-----------------------------------------------")
			fmt.Println()
			cnt--
			if cnt == 0 {
				// s.Wg.Done()
				s.Cancel()
			}
		case <-s.Ctx.Done():
			s.Wg.Done()
			return
		}
	}

	// for _, findInfo := range findInfos {
	// 	fmt.Printf("FileName: %s\n", findInfo.Filename)
	// 	fmt.Println("-----------------------------------------------")
	// 	for _, lineInfo := range findInfo.Lines {
	// 		fmt.Println("\t", lineInfo.LineNo, "\t", lineInfo.Line)
	// 	}
	// 	fmt.Println("-----------------------------------------------")
	// 	fmt.Println()
	// }
}
