package findWords

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type LineInfo struct {
	LineNo int
	Line   string
}

type FindInfo struct {
	Filename string
	Lines    []LineInfo
}

// func GetFileList(path string) ([]string, error) {
// 	return filepath.Glob(path)
// }

func FindWordInAllfile(word, filename string, infoCh chan FindInfo, wg *sync.WaitGroup) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot find file!! ", filename)
		infoCh <- findInfo
		wg.Done()
		return
	}

	defer wg.Done()
	defer file.Close()

	lineNo := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.Lines = append(findInfo.Lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	infoCh <- findInfo
}

func FindWordInAllfiles(word, path string, wg *sync.WaitGroup, infoCh chan FindInfo) {
	// findInfos := []FindInfo{}

	// filelist, err := filepath.Glob(path)
	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("Wrong file path !! err: ", err, "path: ", path)
		return
	}

	// ch := make(chan FindInfo)
	// cnt := len(filelist)
	// recvCnt := 0

	for _, filename := range filelist {
		go FindWordInAllfile(word, filename, infoCh, wg)
	}

	wg.Done()
	// for findInfo := range ch {
	// 	findInfos = append(findInfos, findInfo)
	// 	recvCnt++
	// 	if recvCnt == cnt {
	// 		cancel()
	// 		break
	// 	}
	// }
	// return findInfos
}

func GetFileList(pattern string) ([]string, error) {
	filelist := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return filelist, nil
}
