package day16

import (
	"bytes"
	"fmt"
	"zsoki/aoc/common"
)

var byteGrid [][]byte
var pointGrid [][]int

func Day16a() {
	lines := make(chan string)
	go common.ReadLines("input/day16.txt", lines)

	byteGrid = make([][]byte, 0)
	var startCoord common.Coord
	var endCoord common.Coord

	curRow := 0
	for line := range lines {
		byteLine := []byte(line)
		startIdx := bytes.IndexByte(byteLine, 'S')
		if startIdx != -1 {
			startCoord = common.Coord{Row: curRow, Col: startIdx}
			byteLine = bytes.Replace(byteLine, []byte{'S'}, []byte{'.'}, 1)
		}
		endIdx := bytes.IndexByte(byteLine, 'E')
		if endIdx != -1 {
			endCoord = common.Coord{Row: curRow, Col: endIdx}
			byteLine = bytes.Replace(byteLine, []byte{'E'}, []byte{'.'}, 1)
		}
		byteGrid = append(byteGrid, byteLine)
		curRow++
	}

	pointGrid = make([][]int, len(byteGrid))
	for idx := range pointGrid {
		pointGrid[idx] = make([]int, len(byteGrid[0]))
	}

	fmt.Println("Traversing the graph.")
	bfsQueue := common.Queue{Items: make([]any, 0)}
	type bfsParams struct {
		next    common.Coord
		current common.Coord
		points  int
	}

	for _, neighbor := range traversableNeighbors(startCoord) {
		// Determine whether we needed to rotate
		headingEast := neighbor.Col > startCoord.Col && neighbor.Row == startCoord.Row
		startPoints := 1
		if !headingEast {
			startPoints += 1000
		}
		pointGrid[startCoord.Row][startCoord.Col] += startPoints
		bfsQueue.Enqueue(bfsParams{next: neighbor, current: startCoord, points: startPoints})
	}

	for !bfsQueue.IsEmpty() {
		bfsParam := bfsQueue.Dequeue().(bfsParams)
		last := bfsParam.current
		current := bfsParam.next
		points := bfsParam.points

		if pointGrid[current.Row][current.Col] != 0 && pointGrid[current.Row][current.Col] < points {
			continue
		}

		pointGrid[current.Row][current.Col] = points

		if current == endCoord {
			continue
		}

		for _, neighbor := range traversableNeighbors(current) {
			newPoints := points + 1
			if !common.SameLine(last, neighbor) {
				newPoints += 1000 // Turning.
			}
			bfsQueue.Enqueue(bfsParams{neighbor, current, newPoints})
		}
	}

	fmt.Printf("\nThis is end: %v", pointGrid[endCoord.Row][endCoord.Col])
}

func traversableNeighbors(startCoord common.Coord) []common.Coord {
	return startCoord.CardinalNeighborsWhere(func(neighbor common.Coord) bool {
		return byteGrid[neighbor.Row][neighbor.Col] == '.'
	})
}
