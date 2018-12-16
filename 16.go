package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type opcode struct {
	// id int
	name                   string
	aImmediate, bImmediate bool
	op                     func(valA, valB int) int
}

func (o opcode) doOp(registers [4]int, a, b, c int) [4]int {
	var valA, valB int
	if o.aImmediate {
		valA = a
	} else {
		if a >= len(registers) {
			panic("out of range")
		}
		valA = registers[a]
	}
	if o.bImmediate {
		valB = b
	} else {
		if b >= len(registers) {
			panic("out of range")
		}
		valB = registers[b]
	}
	registers[c] = o.op(valA, valB)
	return registers
}

func main() {
	add := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }
	ban := func(a, b int) int { return a & b }
	bor := func(a, b int) int { return a | b }
	set := func(a, b int) int { return a }
	gt := func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	}
	eq := func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	}
	opcodes := []opcode{
		{name: "addr", op: add},
		{name: "addi", bImmediate: true, op: add},
		{name: "mulr", op: mul},
		{name: "muli", bImmediate: true, op: mul},
		{name: "banr", op: ban},
		{name: "bani", bImmediate: true, op: ban},
		{name: "borr", op: bor},
		{name: "bori", bImmediate: true, op: bor},
		{name: "setr", op: set},
		{name: "seti", aImmediate: true, op: set},
		{name: "gtir", aImmediate: true, op: gt},
		{name: "gtri", bImmediate: true, op: gt},
		{name: "gtrr", op: gt},
		{name: "eqir", aImmediate: true, op: eq},
		{name: "eqri", bImmediate: true, op: eq},
		{name: "eqrr", op: eq},
	}

	numAmbiguousSamples := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && len(scanner.Text()) > 0 {
		beforeLine := strings.NewReader(strings.TrimPrefix(scanner.Text(), "Before: ["))
		if !scanner.Scan() {
			panic("can't scan")
		}
		instructionLine := strings.NewReader(scanner.Text())
		if !scanner.Scan() {
			panic("can't scan")
		}
		afterLine := strings.NewReader(strings.TrimPrefix(scanner.Text(), "After: ["))
		// scan out empty line
		if !scanner.Scan() {
			panic("can't scan")
		}

		var rBefore, rAfter, instruction [4]int
		fmt.Fscanf(beforeLine, "Before: [")
		fmt.Fscanf(afterLine, "After: [")
		for i, n := 0, len(rBefore); i < n; i++ {
			fmt.Fscanf(beforeLine, "%d,", &rBefore[i])
			fmt.Fscanf(afterLine, "%d,", &rAfter[i])
			fmt.Fscanf(instructionLine, "%d", &instruction[i])
		}

		numCandidates := 0
		for _, o := range opcodes {
			result := o.doOp(rBefore, instruction[1], instruction[2], instruction[3])
			equal := true
			for i, n := 0, len(result); i < n; i++ {
				if result[i] != rAfter[i] {
					equal = false
					break
				}
			}
			if equal {
				numCandidates++
			}
		}

		if numCandidates >= 3 {
			numAmbiguousSamples++
		} else if numCandidates == 0 {
			fmt.Println(rBefore, rAfter, instruction, numCandidates)
			panic("no candidates found")
		}
	}
	fmt.Println(numAmbiguousSamples)
}
