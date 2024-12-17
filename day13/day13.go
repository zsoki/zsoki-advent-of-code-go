package day13

import (
	"fmt"
	"strings"
	"zsoki/aoc/common"
)

type puzzleInput struct {
	aButt  common.Coord
	bButt  common.Coord
	target common.Coord
}

type recursiveParams struct {
	depth, whichButton int
	puzzle             puzzleInput
}

type Stack struct {
	items []recursiveParams
}

var minimum = &common.Coord{}

func Day13a() {
	lines := make(chan string)
	go common.ReadLines("input/day13.txt", lines)

	var puzzleParams []puzzleInput

	lineNum := 0
	var buttonA, buttonB, target common.Coord
	for line := range lines {
		if lineNum == 3 {
			lineNum = 0
			puzzleParams = append(puzzleParams, puzzleInput{buttonA, buttonB, target})
			continue
		}

		fields := strings.Fields(line)
		row := common.ToInt(fields[0])
		col := common.ToInt(fields[1])

		switch lineNum {
		case 0:
			buttonA = common.Coord{Row: row, Col: col}
		case 1:
			buttonB = common.Coord{Row: row, Col: col}
		case 2:
			target = common.Coord{Row: row, Col: col}
		}
		lineNum++
	}

	resultSum := doTheThing(puzzleParams)

	fmt.Println(resultSum)
}

func doTheThing(puzzleParams []puzzleInput) int {
	resultSum := 0

	for _, param := range puzzleParams {
		bButton := param.bButt
		aButton := param.aButt

		maxBPush := 0
		acc := common.Coord{}
		for acc.Lt(param.target) {
			acc = bButton.Times(maxBPush)
			maxBPush++
		}
		if acc.Eq(param.target) {
			resultSum += maxBPush
		} else {
			maxBPush = maxBPush - 2

			result := param.target
			possible := false
			aPush := 0

			for ; maxBPush > -1; maxBPush-- {
				aCurr := 0
				aSum := *minimum
				bSum := bButton.Times(maxBPush)
				remaining := result.Sub(bSum)
				for aSum.Lt(remaining) {
					aCurr++
					aSum = aButton.Times(aCurr)
				}
				if aSum.Eq(remaining) {
					aPush = aCurr
					possible = true
					break
				}
			}
			//fmt.Printf("\naPush=%v, maxBPush=%v", aPush, maxBPush)
			if possible && aPush <= 100 && maxBPush <= 100 {
				resultSum += aPush*3 + maxBPush
			}
		}
	}
	return resultSum
}

func Day13b() {
	lines := make(chan string)
	go common.ReadLines("input/day13test2.txt", lines)

	var puzzleParams []puzzleInput

	lineNum := 0
	var buttonA, buttonB, target common.Coord
	for line := range lines {
		if lineNum == 3 {
			lineNum = 0
			target = target.Add(common.Coord{Row: 10000000000000, Col: 10000000000000})
			puzzleParams = append(puzzleParams, puzzleInput{buttonA, buttonB, target})
			continue
		}

		fields := strings.Fields(line)
		row := common.ToInt(fields[0])
		col := common.ToInt(fields[1])

		switch lineNum {
		case 0:
			buttonA = common.Coord{Row: row, Col: col}
		case 1:
			buttonB = common.Coord{Row: row, Col: col}
		case 2:
			target = common.Coord{Row: row, Col: col}
		}
		lineNum++
	}

	resultSum := doTheThing2(puzzleParams)

	fmt.Println(resultSum)
}

func doTheThing2(puzzleParams []puzzleInput) int {
	var resultSum = 0

	for _, param := range puzzleParams {
		bButton := param.bButt
		aButton := param.aButt

		var maxBPush = 0
		bDiv1 := param.target.Row / bButton.Row
		bDiv2 := param.target.Col / bButton.Col
		if bDiv1 < bDiv2 {
			maxBPush = bDiv1
		} else {
			maxBPush = bDiv2
		}

		acc := bButton.Times(maxBPush)

		if acc.Eq(param.target) {
			resultSum += maxBPush
		} else {

			possible := false

			aPush := 0
			bPush := 0

			bSum := acc

			canSkip := false
			lastFound := 0
			skip := 1

			for bPush = maxBPush; ; {
				if bPush < skip {
					// Not found
					fmt.Printf("\nNot found! bPush=%d lastFound=%d maxBPush=%d skip=%d", bPush, lastFound, maxBPush, skip)
					break
				}
				bPush -= skip
				bSum = bButton.Times(bPush)
				remaining := param.target.Sub(bSum)

				aMod1 := remaining.Row % aButton.Row
				aMod2 := remaining.Col % aButton.Col

				fmt.Printf("\naMod1=%d aMod2=%d", aMod1, aMod2)
				if aMod1 == 0 && aMod2 == 0 {
					// Problem is that the modulos could be synced and there never be an offset that aligns 0 with 0
					fmt.Printf("\nbPush=%v, since last found=%v, aMod1=%v, aMod2=%v", bPush, skip, aMod1, aMod2)
					if !canSkip {
						lastFound = bPush
						canSkip = true
					} else if skip == 1 {
						skip = lastFound - bPush
					}
					aDiv1 := remaining.Row / aButton.Row
					aDiv2 := remaining.Col / aButton.Col
					if aDiv1 == aDiv2 {
						aPush = aDiv1
						possible = true
						break
					}
				}
			}

			fmt.Printf("\naPush=%v, bPush=%v\n\n", aPush, bPush)
			if possible {
				resultSum += aPush*3 + bPush
			}
		}
	}

	return resultSum
}
