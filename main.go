package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readRow(fileName string, row int) (string, error) {
	var scanner *bufio.Scanner

	if fileName == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			return "", err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	rowRead := 0
	for scanner.Scan() {
		if rowRead == row {
			return scanner.Text(), scanner.Err()
		}
		rowRead++
	}

	return "", nil
}

func findNumber(rowData string, label string, delim string) string {
	s := strings.Split(rowData, delim)

	for n, header := range s {
		if header == label {
			return strconv.Itoa(n + 1)
		}
	}

	return ""
}

func main() {
	usageText := "USAGE: cn LABEL FILE"

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println(usageText)
		os.Exit(1)
	}

	data, err := readRow(args[1], 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	value := findNumber(data, args[0], ",")
	fmt.Println(value)
}
