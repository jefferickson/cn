package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var args []string
var printData = false

type cache []string

func init() {
	const usageText = "USAGE: cn [-d] LABEL FILE"

	args = os.Args[1:]
	if len(args) < 2 || len(args) > 3 {
		fmt.Println(usageText)
		os.Exit(1)
	} else if len(args) == 3 && args[0] != "-d" {
		fmt.Println(usageText)
		os.Exit(1)
	} else if len(args) == 3 {
		printData = true
		args = args[1:]
	}
}

func main() {
	var fileCache cache

	err := fileCache.readData(args[1], !printData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if len(fileCache) == 0 {
		os.Exit(1)
	}

	value, ok := findNumber(fileCache[0], args[0], ",")
	if !ok {
		os.Exit(1)
	}

	if printData {
		for _, row := range fileCache[1:] {
			if cell := extractCol(row, value, ","); cell != "" {
				fmt.Println(cell)
			}
		}
	} else {
		fmt.Println(value)
	}
}

func (c *cache) readData(fileName string, firstRowOnly bool) error {
	var scanner *bufio.Scanner

	if len(*c) != 0 {
		return nil
	}

	if fileName == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		*c = append(*c, scanner.Text())
		if firstRowOnly {
			return scanner.Err()
		}
	}

	return scanner.Err()
}

func findNumber(rowData string, label string, delim string) (int, bool) {
	s := strings.Split(rowData, delim)

	for n, header := range s {
		if header == label {
			return n + 1, true
		}
	}

	return 0, false
}

func extractCol(rowData string, col int, delim string) string {
	s := strings.Split(rowData, delim)

	for n, colText := range s {
		if n+1 == col {
			return colText
		}
	}

	return ""
}
