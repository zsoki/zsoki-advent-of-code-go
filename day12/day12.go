package day12

import (
	"fmt"
	"log"
	"zsoki/aoc/common"
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

var perpendiculars = map[coord][2]coord{
	north: {east, west},
	east:  {north, south},
	south: {east, west},
	west:  {north, south},
}

var opposites = map[coord]coord{
	north: south,
	east:  west,
	south: north,
	west:  east,
}

type queue struct {
	items []coord
}

func (q *queue) enqueue(data coord) {
	q.items = append(q.items, data)
}

func (q *queue) dequeue() coord {
	if q.isEmpty() {
		log.Panic("Queue is empty")
	}
	returnVal := q.items[0]
	q.items = q.items[1:len(q.items)]
	return returnVal
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func Day12a() {
	lines := make(chan string)
	go common.ReadLines("input/day12.txt", lines)

	theMap := make([][]byte, 0)

	for line := range lines {
		theMap = append(theMap, []byte(line))
	}

	mapNWbounds := coord{0, 0}
	mapSEbounds := coord{len(theMap) - 1, len(theMap[0]) - 1}

	visited := make([][]bool, len(theMap))
	for i := range visited {
		visited[i] = make([]bool, len(theMap[i]))
	}

	fencePrice := 0

	startCoord, unvisitedRemaining := firstUnvisited(visited)
	for unvisitedRemaining {
		fifo := queue{make([]coord, 0)}
		visited[startCoord.row][startCoord.col] = true
		fifo.enqueue(startCoord)

		area := 1
		perimeter := 0

		for !fifo.isEmpty() {
			curCoord := fifo.dequeue()
			curByte := theMap[curCoord.row][curCoord.col]

			// Look around
			for _, direction := range directions {
				nextCoord := curCoord.add(direction)

				if nextCoord.lt(mapNWbounds) || nextCoord.gt(mapSEbounds) {
					// Neighbor is out of bounds, but that means we need fence.
					perimeter++
					continue
				}

				nextByte := theMap[nextCoord.row][nextCoord.col]

				// If neighbor in same region, and unvisited, then queue it
				if curByte == nextByte && !visited[nextCoord.row][nextCoord.col] {
					visited[nextCoord.row][nextCoord.col] = true
					area++
					fifo.enqueue(nextCoord)
				} else if curByte != nextByte {
					// Neighbor is another region. Increase perimeter.
					perimeter++
				}
			}
		}

		// Region is traversed.
		fencePrice += area * perimeter
		startCoord, unvisitedRemaining = firstUnvisited(visited)
	}

	fmt.Printf("\nPrice of fences: %v", fencePrice)
}

type coordDir struct {
	crd, dir coord
}

func Day12b() {
	lines := make(chan string)
	go common.ReadLines("input/day12.txt", lines)

	theMap := make([][]byte, 0)
	for line := range lines {
		theMap = append(theMap, []byte(line))
	}

	mapNWbounds := coord{0, 0}
	mapSEbounds := coord{len(theMap) - 1, len(theMap[0]) - 1}

	visited := make([][]bool, len(theMap))
	for i := range visited {
		visited[i] = make([]bool, len(theMap[i]))
	}

	fencePrice := 0

	startCoord, unvisitedRemaining := firstUnvisited(visited)
	for unvisitedRemaining {
		fifo := queue{make([]coord, 0)}
		visited[startCoord.row][startCoord.col] = true
		fifo.enqueue(startCoord)
		area := 1

		fenceVisited := make(map[coordDir]struct{})
		sides := 0

		for !fifo.isEmpty() {
			curCoord := fifo.dequeue()
			curByte := theMap[curCoord.row][curCoord.col]

			// Look around
			for _, direction := range directions {
				nextCoord := curCoord.add(direction)

				if nextCoord.lt(mapNWbounds) || nextCoord.gt(mapSEbounds) {
					// Neighbor is out of bounds, check if we counted the side already
					_, fncVisited := fenceVisited[coordDir{nextCoord, direction}]
					if fncVisited {
						// Already walked along the perimeter
						continue
					}
					walkAlongPerimeter(&fenceVisited, nextCoord, &theMap, direction, curByte, [2]coord{mapNWbounds, mapSEbounds})
					sides++
					continue
				}

				nextByte := theMap[nextCoord.row][nextCoord.col]

				// If neighbor in same region, and unvisited, then queue it
				if curByte == nextByte && !visited[nextCoord.row][nextCoord.col] {
					visited[nextCoord.row][nextCoord.col] = true
					area++
					fifo.enqueue(nextCoord)
				} else if curByte != nextByte {
					// Different region, that means a side
					_, fncVisited := fenceVisited[coordDir{nextCoord, direction}]
					if fncVisited {
						// Already walked along the perimeter
						continue
					}
					walkAlongPerimeter(&fenceVisited, nextCoord, &theMap, direction, curByte, [2]coord{mapNWbounds, mapSEbounds})
					sides++
				}
			}
		}

		// Region is traversed.
		fencePrice += area * sides
		startCoord, unvisitedRemaining = firstUnvisited(visited)
	}

	fmt.Printf("\nPrice of fences: %v", fencePrice)

}

func walkAlongPerimeter(fenceVisited *map[coordDir]struct{}, perimeterCoord coord, theMap *[][]byte, direction coord, regionByte byte, bounds [2]coord) {
	// This is an unvisited side. Mark along the perimeter
	(*fenceVisited)[coordDir{perimeterCoord, direction}] = struct{}{}

	// Walk along the side perimeter along both directions
	for _, perpendicularDir := range perpendiculars[direction] {
		nextPerpCoord := perimeterCoord.add(perpendicularDir)
		nextPerpByte := safeByte(theMap, nextPerpCoord, bounds)
		lookBackCoord := nextPerpCoord.add(opposites[direction])
		lookBackByte := safeByte(theMap, lookBackCoord, bounds)
		for nextPerpByte != regionByte && lookBackByte == regionByte {
			(*fenceVisited)[coordDir{nextPerpCoord, direction}] = struct{}{}
			nextPerpCoord = nextPerpCoord.add(perpendicularDir)
			nextPerpByte = safeByte(theMap, nextPerpCoord, bounds)
			lookBackCoord = nextPerpCoord.add(opposites[direction])
			lookBackByte = safeByte(theMap, lookBackCoord, bounds)
		}
	}
}

// Returns 0 for out of bounds
func safeByte(theMap *[][]byte, nextCoord coord, bounds [2]coord) byte {
	if nextCoord.lt(bounds[0]) || nextCoord.gt(bounds[1]) {
		return 0
	}
	return (*theMap)[nextCoord.row][nextCoord.col]
}

func firstUnvisited(visitedMap [][]bool) (unvisited coord, ok bool) {
	for rowIdx, row := range visitedMap {
		for colIdx, visited := range row {
			if !visited {
				return coord{rowIdx, colIdx}, true
			}
		}
	}
	return coord{}, false
}
