package main

import (
	"fmt"
	"time"
	"zsoki/aoc/day17"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n\nExecution time: %v\n", time.Since(start))
	}
}

func main() {
	defer timer()()
	day17.Day17a()
}
