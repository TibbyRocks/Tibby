package utils

import (
	"bufio"
	"os"
)

func FileToSlice(inputFile string) []string {
	var lines []string
	file, err := os.Open(inputFile)
	if err != nil {
		log.Error(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Error(err.Error())
	}

	return lines
}
