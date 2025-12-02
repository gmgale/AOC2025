package aocutils

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadInput reads the input.txt file for a given day and returns a slice of strings
func ReadInput(day string) []string {
	path := filepath.Join(day, "input.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
