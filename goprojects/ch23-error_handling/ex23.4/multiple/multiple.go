package multiple

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func multipleFormatString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d, err: %w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d, err: %w", pos, err)
	}

	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word: %s, err: %w", word, err)
	}
	return number, len(word), nil
}

func ReadEq(eq string) {
	rst, err := multipleFormatString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}
