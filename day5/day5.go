package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"zsoki/aoc/common"
)

func Day5a() {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	beforeArr := make([][]int, 100)
	afterArr := make([][]int, 100)
	firstPart := true
	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstPart = false
			continue
		}

		if firstPart {
			split := strings.Split(line, "|")
			before, after := common.ToInt(split[0]), common.ToInt(split[1])
			beforeArr[before] = append(beforeArr[before], after)
			afterArr[after] = append(afterArr[after], before)
		} else {
			split := strings.Split(line, ",")
			nums := make([]int, len(split))
			for i, word := range split {
				nums[i] = common.ToInt(word)
			}

			correct := true
			for currentIdx, current := range nums {
				for compareIdx, compare := range nums {
					if currentIdx == compareIdx {
						continue
					}

					switch {
					case compareIdx < currentIdx:
						if !slices.Contains(afterArr[current], compare) {
							correct = false
							break
						}
					case compareIdx > currentIdx:
						if !slices.Contains(beforeArr[current], compare) {
							correct = false
							break
						}
					}
				}
				if !correct {
					break
				}
			}
			if correct {
				result += nums[len(nums)/2]
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("\nDay 5 A result: %v", result)
}

func Day5b() {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	beforeArr := make([][]int, 100)
	afterArr := make([][]int, 100)
	firstPart := true
	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstPart = false
			continue
		}

		if firstPart {
			split := strings.Split(line, "|")
			before, after := common.ToInt(split[0]), common.ToInt(split[1])

			beforeArr[before] = append(beforeArr[before], after)
			afterArr[after] = append(afterArr[after], before)

		} else {
			split := strings.Split(line, ",")
			nums := make([]int, len(split))
			for i, word := range split {
				nums[i] = common.ToInt(word)
			}

			correct := true
			addToSum := false
			for leftIdx := 0; leftIdx < len(nums)-1; leftIdx++ {
				left := nums[leftIdx]
				for rightIdx := leftIdx + 1; rightIdx < len(nums); rightIdx++ {
					right := nums[rightIdx]

					if slices.Contains(beforeArr[right], left) {
						correct = false
						corrected := make([]int, len(nums), len(nums))
						corrected = append(nums[:rightIdx], nums[rightIdx+1:]...)
						corrected = append(corrected[:leftIdx], append([]int{right}, corrected[leftIdx:]...)...)
						nums = corrected
						break
					}
				}
				if !correct {
					addToSum = true
					correct = true
					leftIdx = -1 // It will add +1 at the new iteration and we want to start at 0
					continue
				}
			}
			if addToSum {
				result += nums[len(nums)/2]
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("\nDay 5 B result: %v", result)
}
