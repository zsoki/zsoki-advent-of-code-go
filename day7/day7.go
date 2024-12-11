package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"zsoki/aoc/common"
)

func Day7a() {
	lines := make(chan string)
	go common.ReadLines("input/day7.txt", lines)

	var calibratedSum uint64 = 0

	for line := range lines {
		var colonIdx = strings.IndexRune(line, ':')
		var numberFields = strings.Fields(line[colonIdx+1:])
		var expected, _ = strconv.ParseUint(line[:colonIdx], 10, 64)

		numbers := make([]uint64, len(numberFields))
		for i, field := range numberFields {
			numbers[i], _ = strconv.ParseUint(field, 10, 64)
		}

		// Operations between numbers. False is +, true is *
		var permutation uint64 = 0
		var operations []bool
		var lastPermutation uint64 = 1

		for i := 0; i < len(numbers)-2; i++ {
			lastPermutation <<= 1
			lastPermutation++
		}

		var actual uint64
		for permutation <= lastPermutation {
			operations = intToBoolArray(permutation, len(numbers)-1)
			//fmt.Printf("\n%v", operations)
			actual = numbers[0]

			for rightOperandIdx := 1; rightOperandIdx < len(numbers); rightOperandIdx++ {
				switch operations[rightOperandIdx-1] {
				case true:
					actual *= numbers[rightOperandIdx]
				case false:
					actual += numbers[rightOperandIdx]
				}
			}

			if actual == expected {
				calibratedSum += expected
				break
			}
			permutation++
		}
		//fmt.Printf("\n--------------------\nexpected=%v\nnumbers=%v\npermutation=%v\nallMults=%v", expected, numbers, permutation, lastPermutation)
	}

	fmt.Printf("\nPossible to calibrate: %v", calibratedSum)
}

func Day7b() {
	lines := make(chan string)
	go common.ReadLines("input/day7.txt", lines)

	var calibratedSum uint64 = 0

	for line := range lines {
		var colonIdx = strings.IndexRune(line, ':')
		var numberFields = strings.Fields(line[colonIdx+1:])
		var expected, _ = strconv.ParseUint(line[:colonIdx], 10, 64)

		numbers := make([]uint64, len(numberFields))
		for i, field := range numberFields {
			numbers[i], _ = strconv.ParseUint(field, 10, 64)
		}

		var possible = calculate(numbers, expected)
		if possible {
			calibratedSum += expected
		}
		//fmt.Printf("\n--------------------\nexpected=%v\nnumbers=%v\npermutation=%v\nallMults=%v", expected, numbers, permutation, lastPermutation)
	}

	fmt.Printf("\nPossible to calibrate: %v", calibratedSum)
}

const (
	plus = iota
	multiply
	concat
)

func calculate(numbers []uint64, expected uint64) bool {
	if calculateRecursive(numbers, plus, 1, numbers[0], expected) ||
		calculateRecursive(numbers, multiply, 1, numbers[0], expected) ||
		calculateRecursive(numbers, concat, 1, numbers[0], expected) {
		return true
	}
	return false
}

func calculateRecursive(numbers []uint64, operation, idx int, accumulator, expected uint64) bool {
	if accumulator > expected {
		return false
	}
	if idx == len(numbers) {
		if accumulator == expected {
			return true
		}
		return false
	}

	switch operation {
	case plus:
		accumulator += numbers[idx]
	case multiply:
		accumulator *= numbers[idx]
	case concat:
		accumulator = concatInts(accumulator, numbers[idx])
	}

	idx++
	if calculateRecursive(numbers, plus, idx, accumulator, expected) ||
		calculateRecursive(numbers, multiply, idx, accumulator, expected) ||
		calculateRecursive(numbers, concat, idx, accumulator, expected) {
		return true
	}
	return false
}

func intToBoolArray(number uint64, desiredLength int) []bool {
	var result []bool

	for number > 0 {
		bit := number & 1
		result = append([]bool{bit == 1}, result...)
		number >>= 1
	}

	// Pad with false values
	if len(result) < desiredLength {
		result = append(make([]bool, desiredLength-len(result)), result...)
	}

	return result
}

func concatInts(left, right uint64) uint64 {
	digits := int(math.Log10(float64(right))) + 1
	return left*uint64(math.Pow(10, float64(digits))) + right
}
