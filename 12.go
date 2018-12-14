package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const scope = 5
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

	fmt.Println(rules)
	var states [][]rune
	states = append(states, []rune(initialState))
	fmt.Println(len(rules), string(states[0]))
	for gen := 1; gen <= 200; gen++ { // TODO: use alternating states
		padded := append([]rune("....."), states[gen-1]...)
		padded = append(padded, []rune(".....")...)
		states = append(states, make([]rune, len(padded)))
		// offset := gen * scope
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
						firstIndex = i
					}
					lastIndex = i
				}
			} else {
				states[gen][i+2] = '.'
			}
		}
		lastIndex = lastIndex
		// fmt.Println(gen, string(padded[offset-5:]), firstIndex-offset, lastIndex-offset)
		// fmt.Println(gen, string(states[gen][offset-5:]))
		// if string(states[gen][0:offsetLeft]) != string(empty) || string(states[gen][len(states[gen])-offsetRight:]) != string(empty) {
		// 	fmt.Println(gen, string(states[gen]))
		// }
	}
	sum := 0
	for i, r := range states[200] {
		if r == '#' {
			fmt.Printf("%d ", i-200*scope)
			sum += (i - 200*scope)
		}
	}
	fmt.Println(sum)
}
