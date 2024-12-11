package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ToInt(word string) int {
	number, err := strconv.Atoi(word)
	if err != nil {
		log.Panicf("Error converting word '%s' to integer: %v\n", word, err)
	}
	return number
}

func Replace(text string, index int, char byte) string {
	lineRunes := []byte(text)
	lineRunes[index] = char
	return string(lineRunes)
}

func ReadLines(name string, lines chan string) {
	defer close(lines)

	file, fileOpenError := os.Open(name)
	if fileOpenError != nil {
		log.Panicf("Error opening file: %v", fileOpenError)
	}
	defer func(file *os.File) {
		fileCloseError := file.Close()
		if fileCloseError != nil {
			log.Panicf("Error closing file: %v", fileCloseError)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Panicf("Error scanning file: %v", err)
	}
}

func AddCoord(left [2]int, right [2]int) [2]int {
	var result [2]int
	for i := 0; i < 2; i++ {
		result[i] = left[i] + right[i]
	}
	return result
}

func SubtractCoord(left [2]int, right [2]int) [2]int {
	var result [2]int
	for i := 0; i < 2; i++ {
		result[i] = left[i] - right[i]
	}
	return result
}
