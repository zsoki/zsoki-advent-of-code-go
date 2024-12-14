package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"zsoki/aoc/common"
)

type puzzleInput struct {
	buttons *[2]common.Coord
	reward  common.Coord
}

const (
	aIdx = iota
	bIdx
)

type recursiveParams struct {
	depth, whichButton int
	puzzle             puzzleInput
}

type Stack struct {
	items []recursiveParams
}

var minimum *common.Coord = &common.Coord{}

func (q *Stack) Push(data recursiveParams) {
	q.items = append(q.items, data)
}

func (q *Stack) Pop() recursiveParams {
	if q.isEmpty() {
		log.Panic("Queue is empty")
	}
	returnVal := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return returnVal
}

func (q *Stack) isEmpty() bool {
	return len(q.items) == 0
}

func Day13a() {
	lines := make(chan string)
	go common.ReadLines("input/day13.txt", lines)

	var puzzleParams []puzzleInput

	lineNum := 0
	var buttonA, buttonB, reward common.Coord
	for line := range lines {
		if lineNum == 3 {
			lineNum = 0
			puzzleParams = append(puzzleParams, puzzleInput{&[2]common.Coord{buttonA, buttonB}, reward})
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
			reward = common.Coord{Row: row, Col: col}
		}
		lineNum++
	}

	resultSum := doTheThing(puzzleParams)

	fmt.Println(resultSum)
}

func doTheThing(puzzleParams []puzzleInput) int {
	resultSum := 0

	for _, param := range puzzleParams {
		bButton := param.buttons[bIdx]
		aButton := param.buttons[aIdx]

		maxBPush := 0
		acc := common.Coord{}
		for acc.Lt(param.reward) {
			acc = bButton.Times(maxBPush)
			maxBPush++
		}
		if acc.Eq(param.reward) {
			resultSum += maxBPush
		} else {
			maxBPush = maxBPush - 2

			result := param.reward
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

type bigCoord struct {
	Row, Col uint64
}

func (left bigCoord) add(right bigCoord) bigCoord {
	return bigCoord{left.Row + right.Row, left.Col + right.Col}
}

func (left bigCoord) sub(right bigCoord) bigCoord {
	return bigCoord{left.Row - right.Row, left.Col - right.Col}
}

func (left bigCoord) gt(right bigCoord) bool {
	return left.Row > right.Row && left.Col > right.Col
}

func (left bigCoord) lt(right bigCoord) bool {
	return left.Row < right.Row && left.Col < right.Col
}

func (left bigCoord) eq(right bigCoord) bool {
	return left.Row == right.Row && left.Col == right.Col
}

func (left bigCoord) times(mult uint64) bigCoord { return bigCoord{left.Row * mult, left.Col * mult} }

type bigPuzzle struct {
	buttons *[2]bigCoord
	target  bigCoord
}

func Day13b() {
	lines := make(chan string)
	go common.ReadLines("input/day13test2.txt", lines)

	var puzzleParams []bigPuzzle

	lineNum := 0
	var buttonA, buttonB, reward bigCoord
	for line := range lines {
		if lineNum == 3 {
			lineNum = 0
			//puzzleParams = append(puzzleParams, bigPuzzle{&[2]bigCoord{buttonA, buttonB}, reward.add(bigCoord{10000000000000, 10000000000000})})
			puzzleParams = append(puzzleParams, bigPuzzle{&[2]bigCoord{buttonA, buttonB}, reward})
			continue
		}

		fields := strings.Fields(line)
		row, _ := strconv.ParseUint(fields[0], 10, 64)
		col, _ := strconv.ParseUint(fields[1], 10, 64)

		switch lineNum {
		case 0:
			buttonA = bigCoord{Row: row, Col: col}
		case 1:
			buttonB = bigCoord{Row: row, Col: col}
		case 2:
			reward = bigCoord{Row: row, Col: col}
		}
		lineNum++
	}

	resultSum := doTheThing2(puzzleParams)

	fmt.Println(resultSum)
}

func doTheThing2(puzzleParams []bigPuzzle) uint64 {
	var resultSum uint64 = 0

	for _, param := range puzzleParams {
		bButton := param.buttons[bIdx]
		aButton := param.buttons[aIdx]

		var maxBPush uint64 = 0
		bDiv1 := param.target.Row / bButton.Row
		bDiv2 := param.target.Col / bButton.Col
		if bDiv1 < bDiv2 {
			maxBPush = bDiv1
		} else {
			maxBPush = bDiv2
		}

		acc := bButton.times(maxBPush)

		if acc.eq(param.target) {
			resultSum += maxBPush
		} else {

			possible := false

			var aPush uint64 = 0
			var bPush uint64 = 0

			bSum := acc

			canSkip := false
			lastFound := uint64(0)
			skip := uint64(1)

			for bPush = maxBPush; bPush > 0; bPush -= skip {
				//fmt.Printf("\nmaxBPush=%v", maxBPush)
				bSum = bButton.times(bPush)
				remaining := param.target.sub(bSum)

				//fmt.Printf("\nBelementem, bPush=%v, skip=%v", bPush, skip)

				aMod1 := remaining.Row % aButton.Row
				aMod2 := remaining.Col % aButton.Col
				if aMod1 == 0 && aMod2 == 0 {
					//fmt.Printf("\nbPush=%v, since last found=%v, aMod1=%v, aMod2=%v", bPush, skip, aMod1, aMod2)
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

				//if aMod1 == 0 && aMod2 == 0 {
				//	aMod2 := remaining.Col % aButton.Col
				//	if aMod2 == 0 {
				//		aDiv1 := remaining.Row / aButton.Row
				//		aDiv2 := remaining.Col / aButton.Col
				//		if aDiv1 == aDiv2 {
				//			aPush = aDiv1
				//			possible = true
				//			break
				//		}
				//	}
				//}
				//bSum = bSum.sub(bButton)
			}

			fmt.Printf("\naPush=%v, bPush=%v\n\n", aPush, bPush)
			if possible {
				resultSum += aPush*3 + bPush
			}
		}
	}

	return resultSum
}
