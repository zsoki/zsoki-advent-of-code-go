package day17

import (
	"fmt"
	"strings"
	"zsoki/aoc/common"
)

var A int
var B int
var C int

type instruction byte

const (
	adv instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

var (
	lit0 = 0
	lit1 = 1
	lit2 = 2
	lit3 = 3
)

var comboOperand = map[byte]*int{
	0: &lit0,
	1: &lit1,
	2: &lit2,
	3: &lit3,
	4: &A,
	5: &B,
	6: &C,
}

func Day17a() {
	lines := make(chan string)
	go common.ReadLines("input/day17test.txt", lines)

	var prg program

	lineNum := 0
	for line := range lines {
		fields := strings.Fields(line)
		switch lineNum {
		case 0:
			A = common.ToInt(fields[2])
		case 1:
			B = common.ToInt(fields[2])
		case 2:
			C = common.ToInt(fields[2])
		case 4:
			prg = parseProgram(fields[1])
		}
		lineNum++
	}

	fmt.Printf("A: %d, B: %d, C: %d\n", A, B, C)
	output := prg.run()
	fmt.Println(output)
}

func parseProgram(s string) program {
	prg := program{make([]byte, 0)}
	for _, numString := range strings.Split(s, ",") {
		prg.code = append(prg.code, byte(common.ToInt(numString)))
	}
	return prg
}

type program struct {
	code []byte
}

func (p *program) run() string {
	output := ""
	for iPtr := 0; iPtr < len(p.code); iPtr += 2 {
		instr := instruction(p.code[iPtr])
		literalOp := p.code[iPtr+1]
		switch instr {
		case adv:
			opValue := *comboOperand[literalOp]
			numerator := A
			denominator := denom(opValue)
			A = numerator / denominator // Truncation?
		case bxl:
			B = B ^ int(literalOp)
		case bst:
			opValue := *comboOperand[literalOp]
			B = opValue % 8
		case jnz:
			if A == 0 {
				continue
			}
			iPtr = int(literalOp) - 2
		case bxc:
			B = B ^ C
		case out:
			opValue := *comboOperand[literalOp]
			output += fmt.Sprintf("%d,", opValue%8)
		case bdv:
			opValue := *comboOperand[literalOp]
			numerator := A
			denominator := denom(opValue)
			B = numerator / denominator
		case cdv:
			opValue := *comboOperand[literalOp]
			numerator := A
			denominator := denom(opValue)
			C = numerator / denominator
		}
	}
	if len(output) > 0 {
		return output[:len(output)-1] // Trim last ','
	}
	return output
}

func denom(value int) int {
	switch value {
	case 0:
		return 1
	case 1:
		return 2
	default:
		return 1 << value
	}
}
