package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"zsoki/aoc/common"
)

func Day11a() {
	lines := make(chan string)
	go common.ReadLines("input/day11.txt", lines)

	var stones = make([]uint64, 0)

	for line := range lines {
		for _, word := range strings.Fields(line) {
			num, _ := strconv.ParseUint(word, 10, 64)
			stones = append(stones, num)
		}
	}

	var start time.Time
	for blinkIdx := 0; blinkIdx < 25; blinkIdx++ {
		start = time.Now()
		for stoneIdx := len(stones) - 1; stoneIdx >= 0; stoneIdx-- {
			stone := stones[stoneIdx]
			if stone == 0 {
				stones[stoneIdx] = 1
			} else {
				digits := calcDigits(stone)
				if digits%2 == 0 {
					leftStone, rightStone := splitStone(stone, digits)
					stones[stoneIdx] = leftStone
					stones = insertStone(&stones, stoneIdx+1, rightStone)
				} else {
					stones[stoneIdx] = stone * 2024
				}
			}
		}
		fmt.Printf("\nBlink %d - Took %v", blinkIdx, time.Since(start))
	}

	fmt.Printf("No. of stones: %d\n", len(stones))
}

func insertStone(stones *[]uint64, insertIdx int, insertStone uint64) []uint64 {
	newStones := make([]uint64, len(*stones)+1)
	copy(newStones[:insertIdx], (*stones)[:insertIdx])
	newStones[insertIdx] = insertStone
	copy(newStones[insertIdx+1:], (*stones)[insertIdx:])
	return newStones
}

type memo struct {
	number    uint64
	remBlinks int
}

func Day11b() {
	lines := make(chan string)
	go common.ReadLines("input/day11.txt", lines)

	cache := make(map[memo]int)
	startNums := make([]uint64, 0)

	for line := range lines {
		for _, word := range strings.Fields(line) {
			num, _ := strconv.ParseUint(word, 10, 64)
			startNums = append(startNums, num)
		}
	}

	allSplits := 0
	remainingBlinks := 75
	for _, num := range startNums {
		allSplits += calculateSplits(&cache, num, remainingBlinks)
	}

	fmt.Printf("\nNo. of stones: %d\n", len(startNums)+allSplits)
}

func calculateSplits(cache *map[memo]int, num uint64, remainingBlinks int) int {
	if remainingBlinks == 0 {
		return 0
	}
	cachedSplits, contains := (*cache)[memo{num, remainingBlinks}]
	if contains {
		return cachedSplits
	} else {
		splits := 0
		if num == 0 {
			splits = calculateSplits(cache, 1, remainingBlinks-1)
		} else {
			digits := calcDigits(num)
			if digits%2 == 0 {
				leftStone, rightStone := splitStone(num, digits)
				splits++
				splits += calculateSplits(cache, leftStone, remainingBlinks-1) + calculateSplits(cache, rightStone, remainingBlinks-1)
			} else {
				splits = calculateSplits(cache, num*2024, remainingBlinks-1)
			}
		}
		(*cache)[memo{num, remainingBlinks}] = splits
		return splits
	}
}

func calcDigits(stone uint64) int {
	return int(math.Log10(float64(stone))) + 1
}

func splitStone(stone uint64, digits int) (uint64, uint64) {
	divisor := uint64(math.Pow10(digits / 2))
	left := stone / divisor
	right := stone - left*divisor
	return left, right
}
