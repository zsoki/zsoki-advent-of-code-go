package common

import (
	"strconv"
	"testing"
)

func TestGrid(t *testing.T) {
	lines := make(chan string)
	go ReadLines("grid_test_in.txt", lines)

	grid := Grid[byte]{}

	for line := range lines {
		grid = append(grid, []byte(line))
	}

	center := Coord{2, 2}

	for dirIdx, dir := range AllDirections {
		expected := strconv.Itoa(dirIdx)
		actual := string(grid.get(center.Add(dir)))
		if actual != expected {
			t.Errorf("Direction %d: expected %s, actual %s", dirIdx, expected, actual)
		}
	}

}
