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

type sample struct {
	opID          int
	numCandidates int
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

	unidentified := make([][]*sample, len(opcodes))

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

		var candidates []int
		for index, o := range opcodes {
			result := o.doOp(rBefore, instruction[1], instruction[2], instruction[3])
			equal := true
			for i, n := 0, len(result); i < n; i++ {
				if result[i] != rAfter[i] {
					equal = false
					break
				}
			}
			if equal {
				candidates = append(candidates, index)
			}
		}

		s := &sample{
			opID:          instruction[0],
			numCandidates: len(candidates),
		}
		for _, c := range candidates {
			unidentified[c] = append(unidentified[c], s)
		}
	}

	// for each iteration:
	// find samples with one unidentified candidate
	// map opcode with the candidate
	// remove all samples with identified opcodes
	identified := make(map[int]int)
	hasSamples := true
	for hasSamples {
		hasSamples = false
		for opIndex, samples := range unidentified {
			if len(samples) > 0 {
				hasSamples = true
				var filtered []*sample
				for _, s := range samples {
					if _, ok := identified[s.opID]; !ok {
						if s.numCandidates == 1 {
							// identified
							identified[s.opID] = opIndex
							for _, ss := range samples {
								if ss.opID != s.opID {
									ss.numCandidates--
								}
							}
							filtered = nil
							break
						}
						filtered = append(filtered, s)
					}
				}
				unidentified[opIndex] = filtered
			}
		}
	}
	fmt.Println(identified)

	var registers, instruction [4]int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		line := strings.NewReader(scanner.Text())
		for i, n := 0, len(instruction); i < n; i++ {
			fmt.Fscanf(line, "%d", &instruction[i])
		}
		opID, a, b, c := instruction[0], instruction[1], instruction[2], instruction[3]
		registers = opcodes[identified[opID]].doOp(registers, a, b, c)
		// fmt.Println(registers)
	}
	fmt.Println(registers[0])
}
