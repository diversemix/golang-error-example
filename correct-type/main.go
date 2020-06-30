package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func validCSVContent(content string) bool {
	lines := strings.Split(content, "\n")
	// Map of "number of commas on line" -> "number of lines with that amount of commas on"
	commaCount := make(map[int]int, 0)
	for _, line := range lines {
		if len(line) > 0 {
			commas := strings.Count(line, ",")
			commaCount[commas] = commaCount[commas] + 1
		}
	}
	// only valid if all lines have the same number of commas on
	return len(commaCount) == 1
}

func ReadCSV(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	content := string(b)
	if !validCSVContent(content) {
		return "", errors.New("Invalid CSV content")
	}
	return content, nil
}

func main() {
	fmt.Println(ReadCSV("main.go"))
}
