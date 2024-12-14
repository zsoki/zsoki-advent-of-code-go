package main

import (
	"fmt"
	"time"
	"zsoki/aoc/day13"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n\nExecution time: %v\n", time.Since(start))
	}
}

func main() {
	defer timer()()
	day13.Day13b()
}
