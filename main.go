package main

import (
	"fmt"
	"time"
	"zsoki/aoc/day11"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n\nExecution time: %v\n", time.Since(start))
	}
}

func main() {
	defer timer()()
	day11.Day11b()
	//day11.Test11()
}
