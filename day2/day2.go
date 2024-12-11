package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"zsoki/aoc/common"
)

func Day2ab_bruteforce() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var safeSeries int
	var fixed int
	var lineCount int

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		words := strings.Fields(line)
		nums := make([]int, len(words))
		for i, word := range words {
			nums[i] = common.ToInt(word)
		}

		safe := sequenceSafe(nums)

		if safe {
			safeSeries++
		} else {
			for i := 0; i < len(nums); i++ {
				removed := make([]int, 0, len(nums)-1)
				removed = append(removed, nums[:i]...)
				removed = append(removed, nums[i+1:]...)
				if sequenceSafe(removed) {
					fixed++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Number of safe series: ", safeSeries)
	fmt.Println("Number of corrected series (brute force): ", fixed)
	fmt.Println("Safe + corrected (brute force) series: ", safeSeries+fixed)
}

const (
	decreasing = -1
	increasing = 1
	equals     = 0
)

func calcChange(left int, right int) int {
	switch {
	case left < right:
		return increasing
	case left > right:
		return decreasing
	default:
		return equals
	}
}

func Day2b_logic(nums []int) bool {
	var corrected bool
	var safe = true
	var change int

	// Start with second num, check against previous and next.
	for i := 1; i < len(nums)-1; i++ {
		prev := nums[i-1]
		curr := nums[i]
		next := nums[i+1]

		if i == 1 {
			prevCurrDiffSafe := diffSafe(prev, curr)
			currNextDiffSafe := diffSafe(curr, next)
			prevCurrChange := calcChange(prev, curr)
			currNextChange := calcChange(curr, next)

			if prevCurrDiffSafe && currNextDiffSafe && prevCurrChange == currNextChange {
				// First iteration is OK, can set change and continue
				change = prevCurrChange
				continue
			}

			// If we must remove anything from the beginning, we must assume that the change remains constant until the end -
			// if not, more than one correction would be needed anyway (e.g.: at the end)
			endChange := calcChange(nums[len(nums)-2], nums[len(nums)-1])
			prevNextDiffSafe := diffSafe(prev, next)
			prevNextChange := calcChange(prev, next)

			if endChange == currNextChange && currNextDiffSafe {
				// Remove first
				corrected = true
				change = endChange
			} else if endChange == prevNextChange && prevNextDiffSafe {
				// Remove second
				corrected = true
				change = endChange
				nums[i] = prev
			} else {
				safe = false
				break
			}
			continue
		}

		prevCurrSafe := numsSafe(prev, curr, change)
		currNextSafe := numsSafe(curr, next, change)

		// Determine whether last element must be removed
		if i == len(nums)-2 && prevCurrSafe && !currNextSafe {
			if corrected {
				safe = false
			}
			// Otherwise we hadn't corrected yet, so sequence is safe.
			break
		}

		if !prevCurrSafe || !currNextSafe {
			// Check if removal of curr would be safe
			prevNextSafe := numsSafe(prev, next, change)
			if !prevNextSafe {
				// Cannot fix with removing 1 element from the seq
				safe = false
				break
			}
			if !corrected {
				// Remove the current element.
				nums[i] = nums[i-1]
				corrected = true
				continue
			}

			// Already corrected once, cannot fix further
			safe = false
			break
		}

		// Curr is safe, continue
	}

	return safe
}

func sequenceSafe(nums []int) bool {
	var change int
	var safe = true

	for i, num := range nums {
		if i == 0 {
			continue
		}

		prevNum := nums[i-1]

		if i == 1 {
			if prevNum < num {
				change = 1
			} else if prevNum > num {
				change = -1
			} else {
				safe = false
				break
			}
		}

		safe = numsSafe(prevNum, num, change)
		if !safe {
			break
		}
	}
	return safe
}

func numsSafe(left int, right int, change int) bool {
	return changeSafe(left, right, change) && diffSafe(left, right)
}

func changeSafe(left int, right int, requiredChange int) bool {
	change := calcChange(left, right)
	if change == equals || change != requiredChange {
		return false
	}
	return true
}

func diffSafe(left int, right int) bool {
	diff := left - right
	if diff < 0 {
		diff *= -1
	}
	if diff > 3 {
		return false
	}
	return true
}
