package main

import (
	"bufio"
	"fmt"
	"os"
)

type cache struct {
	offset, gen int
}

func main() {
	const scope = 5
	const numGenerations = 50000000000
	// empty := []rune(".....")
	// offsetLeft, offsetRight := scope, scope
	var initialState string
	fmt.Scanf("initial state: %s\n", &initialState)
	fmt.Scanln()
	scanner := bufio.NewScanner(os.Stdin)
	rules := map[string]rune{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		var state string
		var output rune
		fmt.Sscanf(scanner.Text(), "%s => %c", &state, &output)
		if output == '#' {
			rules[state] = output
		}
	}

	var states [][]rune
	states = append(states, []rune(initialState))

	offset := 0

	visited := map[string]cache{}
	for gen := 1; gen <= numGenerations; gen++ { // TODO: use alternating states
		padded := append([]rune("....."), states[gen-1]...)
		padded = append(padded, []rune(".....")...)
		states = append(states, make([]rune, len(padded)))
		offset += scope
		for i := 0; i < scope; i++ {
			states[gen][i] = '.'
		}
		for i := 1; i <= scope; i++ {
			states[gen][len(padded)-i] = '.'
		}
		firstIndex, lastIndex := -1, -1
		for i := 0; i < len(padded)-scope; i++ {
			chunk := padded[i : i+scope]
			if output, ok := rules[string(chunk)]; ok {
				states[gen][i+2] = output
				if output == '#' {
					if firstIndex == -1 {
						firstIndex = i + 2
					}
					lastIndex = i + 2
				}
			} else {
				states[gen][i+2] = '.'
			}
		}

		offset -= firstIndex

		states[gen] = states[gen][firstIndex : lastIndex+1]

		k := string(states[gen])
		if v, ok := visited[k]; ok {
			// fmt.Println(v, gen, numGenerations%v.gen, offset, v.offset)

			offset += (numGenerations - gen) * (offset - v.offset)

			// sum := big.NewInt(0)
			sum := 0
			for i, r := range k {
				if r == '#' {
					// fmt.Printf("%d ", bigI.Sub(bigI, bigOffset))
					// fmt.Printf("%d ", i-offset)
					sum += i - offset
				}
			}
			fmt.Println(sum)
			return
		}

		visited[k] = cache{
			offset: offset,
			gen:    gen,
		}
	}
}
