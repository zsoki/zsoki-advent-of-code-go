package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4a() {
	file, err := os.Open("input/day4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	crosswords := make([]string, 0, 140)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		crosswords = append(crosswords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	rows := len(crosswords)
	cols := len(crosswords[0])

	check_W := func(col int, row int, cw []string) bool {
		if col < 3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row][col-i], i) {
				return false
			}
		}
		return true
	}
	check_NW := func(col int, row int, cw []string) bool {
		if col < 3 || row < 3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row-i][col-i], i) {
				return false
			}
		}
		return true
	}
	check_N := func(col int, row int, cw []string) bool {
		if row < 3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row-i][col], i) {
				return false
			}
		}
		return true
	}
	check_NE := func(col int, row int, cw []string) bool {
		if row < 3 || col > cols-1-3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row-i][col+i], i) {
				return false
			}
		}
		return true
	}
	check_E := func(col int, row int, cw []string) bool {
		if col > cols-1-3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row][col+i], i) {
				return false
			}
		}
		return true
	}
	check_SE := func(col int, row int, cw []string) bool {
		if row > rows-1-3 || col > cols-1-3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row+i][col+i], i) {
				return false
			}
		}
		return true
	}
	check_S := func(col int, row int, cw []string) bool {
		if row > rows-1-3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row+i][col], i) {
				return false
			}
		}
		return true
	}
	check_SW := func(col int, row int, cw []string) bool {
		if row > rows-1-3 || col < 3 {
			return false
		}
		for i := 0; i <= 3; i++ {
			if !correctChar(cw[row+i][col-i], i) {
				return false
			}
		}
		return true
	}

	fmt.Printf("\nRows: %v", rows)
	fmt.Printf("\nCols: %v", cols)

	occurrence := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if check_W(col, row, crosswords) {
				occurrence++
			}
			if check_NW(col, row, crosswords) {
				occurrence++
			}
			if check_N(col, row, crosswords) {
				occurrence++
			}
			if check_NE(col, row, crosswords) {
				occurrence++
			}
			if check_E(col, row, crosswords) {
				occurrence++
			}
			if check_SE(col, row, crosswords) {
				occurrence++
			}
			if check_S(col, row, crosswords) {
				occurrence++
			}
			if check_SW(col, row, crosswords) {
				occurrence++
			}
		}
	}

	fmt.Printf("\nOccurences: %v", occurrence)
}

func correctChar(char byte, pos int) bool {
	switch pos {
	case 0:
		return char == 'X'
	case 1:
		return char == 'M'
	case 2:
		return char == 'A'
	case 3:
		return char == 'S'
	default:
		return false
	}
}

func Day4b() {
	file, err := os.Open("input/day4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	crosswords := make([]string, 0, 140)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		crosswords = append(crosswords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	rows := len(crosswords)
	cols := len(crosswords[0])

	check_NW_SE := func(col int, row int, cw []string) bool {
		var exclude byte
		var correct bool
		for i := -1; i <= 1; i++ {
			correct, exclude = correctChar2(cw[row+i][col+i], i, exclude)
			if !correct {
				return false
			}
		}
		return true
	}
	check_NE_SW := func(col int, row int, cw []string) bool {
		var exclude byte
		var correct bool
		for i := -1; i <= 1; i++ {
			correct, exclude = correctChar2(cw[row+i][col-i], i, exclude)
			if !correct {
				return false
			}
		}
		return true
	}

	fmt.Printf("\nRows: %v", rows)
	fmt.Printf("\nCols: %v", cols)

	occurrence := 0

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if check_NW_SE(col, row, crosswords) && check_NE_SW(col, row, crosswords) {
				occurrence++
			}
		}
	}

	fmt.Printf("\nOccurences: %v", occurrence)
}

func correctChar2(char byte, pos int, exclude byte) (bool, byte) {
	switch pos {
	case -1:
		if char == 'M' && char != exclude {
			return true, 'M'
		} else if char == 'S' && char != exclude {
			return true, 'S'
		} else {
			return false, exclude
		}
	case 0:
		return char == 'A', exclude
	case 1:
		if char == 'M' && char != exclude {
			return true, 'M'
		} else if char == 'S' && char != exclude {
			return true, 'S'
		} else {
			return false, char
		}
	default:
		{
			fmt.Printf("Could not determine!!! %v %v %v", char, pos, exclude)
			return false, exclude
		}
	}
}
