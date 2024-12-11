package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"zsoki/aoc/common"
)

func Day1a() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var col1 []int
	var col2 []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		col1 = append(col1, common.ToInt(words[0]))
		col2 = append(col2, common.ToInt(words[1]))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	slices.Sort(col1)
	slices.Sort(col2)

	sum := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			sum += diff * -1
			continue
		}
		sum += diff
	}

	fmt.Println(sum)
}

func Day1b() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var col1 []int
	var col2 []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		col1 = append(col1, common.ToInt(words[0]))
		col2 = append(col2, common.ToInt(words[1]))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	simScore := 0
	for _, col1Value := range col1 {
		occurrence := 0
		for _, col2Value := range col2 {
			if col1Value == col2Value {
				occurrence++
			}
		}
		simScore += col1Value * occurrence
	}

	fmt.Println(simScore)
}
