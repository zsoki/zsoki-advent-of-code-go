package day10

import (
	"fmt"
	"zsoki/aoc/common"
)

const (
	startByte byte = '0'
	endByte   byte = '9'
)

type coord struct {
	row, col int
}

func (left coord) add(right coord) coord {
	return coord{left.row + right.row, left.col + right.col}
}

func (left coord) sub(right coord) coord {
	return coord{left.row - right.row, left.col - right.col}
}

func (left coord) gt(right coord) bool {
	return left.row > right.row || left.col > right.col
}

func (left coord) lt(right coord) bool {
	return left.row < right.row || left.col < right.col
}

func (left coord) eq(right coord) bool {
	return left.row == right.row && left.col == right.col
}

var (
	north = coord{-1, 0}
	east  = coord{0, 1}
	south = coord{1, 0}
	west  = coord{0, -1}
)

var directions = [4]coord{
	north,
	east,
	south,
	west,
}

func Day10a() {
	lines := make(chan string)
	go common.ReadLines("input/day10.txt", lines)

	var theMap [][]byte
	var startCoords []coord = make([]coord, 0)

	row := 0
	for line := range lines {
		theMap = append(theMap, []byte(line))
		for col, curByte := range theMap[row] {
			if curByte == startByte {
				startCoords = append(startCoords, coord{row, col})
			}
		}
		row++
	}

	mapLengths := coord{row: len(theMap), col: len(theMap[0])}

	reachableEnds := 0

	// Traverse the map
	for _, start := range startCoords {
		reachableEnds += startTraversal(start, mapLengths, theMap)
	}

	fmt.Printf("Reachable Ends: %d\n", reachableEnds)
}

func startTraversal(start coord, mapLengths coord, theMap [][]byte) int {
	result := make([]coord, 0)
	for _, direction := range directions {
		result = append(result, traverseRecursive(start, start.add(direction), mapLengths, theMap, createVisitMap(mapLengths))...)
	}

	// Deduplicate coordinates
	coordSet := make(map[coord]struct{})
	for _, resultCoord := range result {
		coordSet[resultCoord] = struct{}{}
	}
	return len(coordSet)
}

func traverseRecursive(prev coord, current coord, mapLengths coord, theMap [][]byte, visitMap [][]bool) []coord {
	if current.gt(mapLengths.sub(coord{1, 1})) || current.lt(coord{0, 0}) {
		// Out of bounds
		return nil
	}
	if current.eq(prev) {
		// Do not visit backwards
		return nil
	}
	if visitMap[current.row][current.col] {
		// Already visited in another branch
		return nil
	}
	curByte := theMap[current.row][current.col]
	prevByte := theMap[prev.row][prev.col]
	if curByte-prevByte != 1 {
		// Not increasing by one
		return nil
	}

	visitMap[current.row][current.col] = true
	if curByte == endByte {
		// Found the end of trail, add one
		return []coord{current}
	}

	result := make([]coord, 0)
	// Visit all directions
	for _, direction := range directions {
		result = append(result, traverseRecursive(current, current.add(direction), mapLengths, theMap, visitMap)...)
	}
	return result
}

func createVisitMap(mapLengths coord) [][]bool {
	result := make([][]bool, mapLengths.row)
	for i := range result {
		result[i] = make([]bool, mapLengths.col)
	}
	return result
}

func Day10b() {
	lines := make(chan string)
	go common.ReadLines("input/day10.txt", lines)

	var theMap [][]byte
	var startCoords []coord = make([]coord, 0)

	row := 0
	for line := range lines {
		theMap = append(theMap, []byte(line))
		for col, curByte := range theMap[row] {
			if curByte == startByte {
				startCoords = append(startCoords, coord{row, col})
			}
		}
		row++
	}

	mapLengths := coord{row: len(theMap), col: len(theMap[0])}

	sumRatings := 0

	// Traverse the map
	for _, start := range startCoords {
		sumRatings += startTraversalRatings(start, mapLengths, theMap)
	}
	fmt.Printf("\nSum of all ratings: %d\n", sumRatings)
}

func startTraversalRatings(start coord, mapLengths coord, theMap [][]byte) int {
	sumRatings := 0
	for _, direction := range directions {
		sumRatings += traverseRecursiveRatings(start, start.add(direction), mapLengths, theMap)
	}
	return sumRatings
}

func traverseRecursiveRatings(prev coord, current coord, mapLengths coord, theMap [][]byte) int {
	if current.gt(mapLengths.sub(coord{1, 1})) || current.lt(coord{0, 0}) {
		// Out of bounds
		return 0
	}
	if current.eq(prev) {
		// Do not visit backwards
		return 0
	}
	curByte := theMap[current.row][current.col]
	prevByte := theMap[prev.row][prev.col]
	if curByte-prevByte != 1 {
		// Not increasing by one
		return 0
	}

	if curByte == endByte {
		// Found the end of trail, add one
		return 1
	}

	sumRatings := 0
	// Visit all directions
	for _, direction := range directions {
		sumRatings += traverseRecursiveRatings(current, current.add(direction), mapLengths, theMap)
	}
	return sumRatings
}
