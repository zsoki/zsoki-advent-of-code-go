package day8

import (
	"fmt"
	"zsoki/aoc/common"
)

func Day8a() {
	lines := make(chan string)
	go common.ReadLines("input/day8.txt", lines)

	charMap := make(map[rune][][2]int)
	antiNodeHashSet := make(map[[2]int]struct{})

	maxColIdx := 0
	row := 0
	for line := range lines {
		if maxColIdx == 0 {
			maxColIdx = len(line) - 1
		}
		for col, char := range line {
			if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
				posArray, ok := charMap[char]
				if !ok {
					posArray = make([][2]int, 0)
				}
				posArray = append(posArray, [2]int{row, col})
				charMap[char] = posArray
			}
		}
		row++
	}
	maxRowIdx := row - 1

	for _, posArray := range charMap {
		for posIdx, _ := range posArray {
			if posIdx == len(posArray)-1 {
				break
			}
			for compIdx := posIdx + 1; compIdx < len(posArray); compIdx++ {
				antiNodes := calcAntinodes(posArray[posIdx], posArray[compIdx])
				for _, antiNode := range antiNodes {
					nodeRow := antiNode[0]
					nodeCol := antiNode[1]
					if nodeRow > maxRowIdx || nodeRow < 0 || nodeCol > maxColIdx || nodeCol < 0 {
						continue
					}
					antiNodeHashSet[antiNode] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("\nAntinodes: %v", len(antiNodeHashSet))
}

func Day8b() {
	lines := make(chan string)
	go common.ReadLines("input/day8.txt", lines)

	charMap := make(map[rune][][2]int)
	antiNodeHashSet := make(map[[2]int]struct{})

	maxColIdx := 0
	row := 0
	for line := range lines {
		if maxColIdx == 0 {
			maxColIdx = len(line) - 1
		}
		for col, char := range line {
			if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
				posArray, ok := charMap[char]
				if !ok {
					posArray = make([][2]int, 0)
				}
				posArray = append(posArray, [2]int{row, col})
				charMap[char] = posArray
			}
		}
		row++
	}
	maxRowIdx := row - 1

	for _, posArray := range charMap {
		for posIdx, _ := range posArray {
			if posIdx == len(posArray)-1 {
				break
			}
			for compIdx := posIdx + 1; compIdx < len(posArray); compIdx++ {
				antiNodes := calcAntinodesAll(posArray[posIdx], posArray[compIdx], maxRowIdx, maxColIdx)
				for _, antiNode := range antiNodes {
					antiNodeHashSet[antiNode] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("\nAntinodes: %v", len(antiNodeHashSet))
}

func calcAntinodes(current, compare [2]int) [2][2]int {
	diff := common.SubtractCoord(compare, current)
	return [2][2]int{common.AddCoord(compare, diff), common.SubtractCoord(current, diff)}
}

func calcAntinodesAll(current, compare [2]int, maxColIdx, maxRowIdx int) [][2]int {
	diff := common.SubtractCoord(compare, current)

	results := make([][2]int, 0)

	forwardCompare := compare
	for {
		nextANodeForward := common.AddCoord(forwardCompare, diff)
		nodeRow := nextANodeForward[0]
		nodeCol := nextANodeForward[1]
		if nodeRow > maxRowIdx || nodeRow < 0 || nodeCol > maxColIdx || nodeCol < 0 {
			// out of bounds
			break
		}
		results = append(results, [2]int{nodeRow, nodeCol})
		forwardCompare = nextANodeForward
	}

	backwardsCompare := current
	for {
		nextANodeForward := common.SubtractCoord(backwardsCompare, diff)
		nodeRow := nextANodeForward[0]
		nodeCol := nextANodeForward[1]
		if nodeRow > maxRowIdx || nodeRow < 0 || nodeCol > maxColIdx || nodeCol < 0 {
			// out of bounds
			break
		}
		results = append(results, [2]int{nodeRow, nodeCol})
		backwardsCompare = nextANodeForward
	}

	// The antennas itself are antinodes
	results = append(results, current, compare)
	return results
}
