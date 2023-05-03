package add

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func addFormatString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt() pos: %d, error: %w", pos+n, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt() pos: %d, error: %w", pos+n, err)
	}

	return a + b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("failed to scan")
	}

	word := scanner.Text()
	num, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert string to int error: %w", err)
	}

	return num, len(word), nil
}

func ReadEq(str string) {
	rst, err := addFormatString(str)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var strErr *strconv.NumError
		if errors.As(err, &strErr) {
			fmt.Printf("NumberError: %v", strErr)
		}
	}
}
