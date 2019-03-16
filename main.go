package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var args config

type config struct {
	d          bool
	headerFile string
	label      string
	dataFile   string
	delim      string
}

type cache struct {
	data   []string
	header string
}

func init() {
	d := flag.Bool("d", false, "Print the data in the column instead of the index")
	headerFile := flag.String("h", "", "File to use for header")
	delim := flag.String("delim", ",", "Delimiter")

	flag.Parse()

	tail := flag.Args()
	if len(tail) != 2 {
		fmt.Println("cn: A tool to find the index (1-based) of a header in a CSV or data in a CSV based on the header.\n")
		fmt.Println("Usage:\n\n\tcn [flags] label datafile\n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	args = config{
		d:          *d,
		headerFile: *headerFile,
		label:      tail[0],
		dataFile:   tail[1],
		delim:      *delim,
	}

	if args.headerFile == args.dataFile {
		args.headerFile = ""
	}
}

func main() {
	var fileCache cache

	fileCache.fillCache(args.dataFile, args.headerFile)
	if len(fileCache.header) == 0 {
		os.Exit(1)
	}

	value, ok := findNumber(fileCache.header, args.label, args.delim)
	if !ok {
		os.Exit(1)
	}

	if args.d {
		for _, row := range fileCache.data {
			if cell := extractCol(row, value, args.delim); cell != "" {
				fmt.Println(cell)
			}
		}
	} else {
		fmt.Println(value)
	}
}

func (c *cache) fillCache(dataFile string, headerFile string) {
	// if we already have our data, return
	if len(c.data) != 0 && len(c.header) != 0 {
		return
	}

	// data first
	if dataFile != "" {
		data, err := readData(dataFile, !args.d)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		c.data = data
	}

	// then the header
	if headerFile == "" {
		c.header = c.data[0]
		c.data = c.data[1:]
	} else {
		data, err := readData(headerFile, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		c.header = data[0]
	}

	return
}

func readData(fileName string, firstRowOnly bool) ([]string, error) {
	var scanner *bufio.Scanner
	var result []string

	if fileName == "" {
		return nil, nil
	}

	if fileName == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		result = append(result, scanner.Text())
		if firstRowOnly {
			return result, scanner.Err()
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("file %s is empty", fileName)
	}

	return result, scanner.Err()
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
