package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"zsoki/aoc/common"
)

const (
	north = iota
	east
	south
	west
)

const (
	north2b   = 0b00001
	east2b    = 0b00010
	south2b   = 0b00100
	west2b    = 0b01000
	obstacled = 0b10000
)

func Day6a() {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	theMap := make([]string, 0)
	var row, col, direction = -1, -1, north

	scanner := bufio.NewScanner(file)
	for lineIdx := 0; scanner.Scan(); lineIdx++ {
		line := scanner.Text()
		theMap = append(theMap, line)
		if row == -1 {
			startCol := strings.IndexRune(line, '^')
			if startCol != -1 {
				row = lineIdx
				col = startCol
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// 'X' characters added around the map.

	var visited = 1
	var inbounds = true
	var next byte
	var nextRow, nextCol = row, col

	var checks int
	for checks = 0; inbounds; checks++ {
		switch direction {
		case north:
			nextRow = row - 1
		case east:
			nextCol = col + 1
		case south:
			nextRow = row + 1
		case west:
			nextCol = col - 1
		}
		next = theMap[nextRow][nextCol]
		switch next {
		case 'X':
			inbounds = false
		case '#':
			if direction < west {
				direction += 1
			} else {
				direction = north
			}
		case '.':
			theMap[nextRow] = common.Replace(theMap[nextRow], nextCol, '+')
			visited++
			row, col = nextRow, nextCol
		case '+':
			row, col = nextRow, nextCol
		}
		nextRow, nextCol = row, col
		//fmt.Printf("\n\n\n")
	}

	for _, line := range theMap {
		fmt.Println(line)
	}

	fmt.Printf("\nDay 6A visited (unique): %v", visited)
	fmt.Printf("\nDay 6A checks: %v", checks)

}

// Logic, does not work
/*func Day6b2() {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	theMap := make([]string, 0)
	var row, col = -1, -1

	scanner := bufio.NewScanner(file)
	for lineIdx := 0; scanner.Scan(); lineIdx++ {
		line := scanner.Text()
		theMap = append(theMap, line)
		if row == -1 {
			startCol := strings.IndexRune(line, '^')
			if startCol != -1 {
				row = lineIdx
				col = startCol
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// 'X' characters added around the map.

	visitedMatrix := make([][]byte, len(theMap))
	for idx := range visitedMatrix {
		visitedMatrix[idx] = make([]byte, len(theMap[0]))
	}

	var loopingObstacle = 0
	var inbounds = true
	var next byte
	var nextRow, nextCol = row, col
	var direction int
	var dirBinary byte

	var checks int
	for checks = 0; inbounds; checks++ {
		nextRow, nextCol = row, col

		switch direction {
		case north:
			dirBinary = north2b
			nextRow = row - 1
		case east:
			dirBinary = east2b
			nextCol = col + 1
		case south:
			dirBinary = south2b
			nextRow = row + 1
		case west:
			dirBinary = west2b
			nextCol = col - 1
		}

		visitedMatrix[row][col] += dirBinary
		next = theMap[nextRow][nextCol]

		switch next {
		case 'X':
			inbounds = false
		case '#':
			direction = changeDirection(direction)
		case '.', '^':

			// Check loop. Simulate #
			theMap[nextRow] = common.Replace(theMap[nextRow], nextCol, '#')
			loopDirection := changeDirection(direction)

			simulateMatrix := make([][]byte, len(theMap))
			for idx := range simulateMatrix {
				simulateMatrix[idx] = make([]byte, len(theMap[idx]))
				copy(simulateMatrix[idx][:], visitedMatrix[idx][:])
			}

			var loopInbounds = true
			var loopNext byte
			var loopNextRow, loopNextCol, loopRow, loopCol = row, col, row, col
			var looping = false

			for loopInbounds {
				loopNextRow, loopNextCol = loopRow, loopCol
				var loopDirBinary byte
				switch loopDirection {
				case north:
					loopDirBinary = north2b
					loopNextRow = loopRow - 1
				case east:
					loopDirBinary = east2b
					loopNextCol = loopCol + 1
				case south:
					loopDirBinary = south2b
					loopNextRow = loopRow + 1
				case west:
					loopDirBinary = west2b
					loopNextCol = loopCol - 1
				}
				if simulateMatrix[loopRow][loopCol]&loopDirBinary == loopDirBinary {
					// We were already here going the same way
					looping = true
					break
				} else {
					// Mark visited
					simulateMatrix[loopRow][loopCol] += loopDirBinary
				}

				loopNext = theMap[loopNextRow][loopNextCol]
				switch loopNext {
				case 'X':
					loopInbounds = false
				case '#':
					loopDirection = changeDirection(loopDirection)
				case '.', '^':
					loopRow, loopCol = loopNextRow, loopNextCol
				}
			}
			if looping {
				if visitedMatrix[nextRow][nextCol]&obstacled == 0 {
					loopingObstacle++
					visitedMatrix[nextRow][nextCol] += obstacled
				}
			}

			theMap[nextRow] = common.Replace(theMap[nextRow], nextCol, '.')
			row, col = nextRow, nextCol
		}
	}

	fmt.Printf("\nDay 6B result: %v", loopingObstacle)
	fmt.Printf("\nDay 6B checks: %v", checks)

}*/

// Brute force, works
func Day6b3() {
	theMap := make([]string, 0)
	var startRow, startCol = -1, -1

	lineIdx := -1
	lines := make(chan string)
	go common.ReadLines("input/day6.txt", lines)
	for line := range lines {
		lineIdx++
		theMap = append(theMap, line)
		if startRow == -1 {
			startCol = strings.IndexRune(line, '^')
			if startCol != -1 {
				startRow = lineIdx
			}
		}
	}

	var possiblePositions int

	for tryRow := 1; tryRow < len(theMap)-1; tryRow++ {
		for tryCol := 1; tryCol < len(theMap[0])-1; tryCol++ {
			found := false
			row, col := startRow, startCol

			// Copy the map
			mapCopy := make([]string, len(theMap))
			for mapRowIdx, line := range theMap {
				mapCopy[mapRowIdx] = strings.Clone(line)
			}

			// Replace the cell with obstacle
			mapCopy[tryRow] = common.Replace(mapCopy[tryRow], tryCol, '#')

			// Track visits by direction
			visitedMatrix := make([][]byte, len(theMap))
			for idx := range visitedMatrix {
				visitedMatrix[idx] = make([]byte, len(theMap[0]))
			}

			var inbounds = true
			var next byte
			var nextRow, nextCol = row, col
			var direction int
			var dirBinary byte

			var checks int
			for checks = 0; inbounds; checks++ {
				nextRow, nextCol = row, col

				switch direction {
				case north:
					dirBinary = north2b
					nextRow = row - 1
				case east:
					dirBinary = east2b
					nextCol = col + 1
				case south:
					dirBinary = south2b
					nextRow = row + 1
				case west:
					dirBinary = west2b
					nextCol = col - 1
				}

				if visitedMatrix[row][col]&dirBinary == dirBinary {
					found = true
					break
				} else {
					visitedMatrix[row][col] += dirBinary
				}
				next = mapCopy[nextRow][nextCol]

				switch next {
				case 'X':
					inbounds = false
				case '#':
					direction = changeDirection(direction)
				case '.', '^', 'O':
					row, col = nextRow, nextCol
				}
			}

			if found {
				theMap[tryRow] = common.Replace(theMap[tryRow], tryCol, 'O')
				possiblePositions++
			}
		}
	}

	for _, line := range theMap {
		fmt.Println(line)
	}

	fmt.Printf("\nDay 6B possible positions: %v", possiblePositions)

}

func changeDirection(direction int) int {
	if direction < west {
		return direction + 1
	} else {
		return north
	}
}
