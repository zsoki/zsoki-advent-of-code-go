package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"zsoki/aoc/common"
)

func Day3a() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var regex = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, match := range regex.FindAllStringSubmatch(line, -1) {
			if len(match) != 3 {
				continue
			}
			sum += common.ToInt(match[1]) * common.ToInt(match[2])
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}

func Day3b() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var regex = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	var sum int

	programBytes, err := os.ReadFile("input/day3.txt")
	if err != nil {
		log.Panic("Error reading file:", err)
	}
	line := string(programBytes)

	fmt.Printf("\n\nOriginal + cleaned:\n\n\t%v", line)

	dontIndex := strings.Index(line, "don't()")
	for dontIndex != -1 {
		doIndex := strings.Index(line[dontIndex:], "do()")
		if doIndex != -1 {
			line = line[:dontIndex] + line[dontIndex+doIndex+len("do()"):]
		} else {
			line = line[:dontIndex]
		}
		dontIndex = strings.Index(line, "don't()")
	}

	fmt.Printf("\n\t%v\n", line)

	for _, match := range regex.FindAllStringSubmatch(line, -1) {
		if len(match) != 3 {
			continue
		}
		sum += common.ToInt(match[1]) * common.ToInt(match[2])
	}

	fmt.Println()
	fmt.Printf("\n\nSum: %d\n", sum)
}
