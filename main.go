package main

import (
	"fmt"
	"time"
	"zsoki/aoc/day12"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n\nExecution time: %v\n", time.Since(start))
	}
}

func main() {
	defer timer()()
	day12.Day12a()
}
