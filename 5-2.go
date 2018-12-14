package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func isOpposite(a, b rune) bool {
	// upper is before lower in ascii
	if a > b {
		a, b = b, a
	}
	// after this point, a <= b. this means if a is lower, b can't be upper
	return unicode.IsUpper(a) && unicode.IsLower(b) && a == unicode.ToUpper(b)
}

func main() {
	// polymer := "dabAcCaCBAcCcaDA" // -> dabCBAcaDA
	input, err := ioutil.ReadFile("5-2.in.txt")
	if err != nil {
		panic(err)
	}
	polymer := string(input)

	polymer = strings.TrimSpace(polymer)

	var minLen = len(polymer)

	// find all possibles
	runeMaps := map[rune]bool{}
	for _, unit := range polymer {
		runeMaps[unicode.ToUpper(unit)] = true
	}

	for r := range runeMaps {
		fmt.Println("removing ", string(r))
		var result []rune
		for _, unit := range polymer {
			if unicode.ToUpper(unit) == r { // skipping r
				continue
			}
			n := len(result)
			if n > 0 && isOpposite(result[n-1], unit) {
				result = result[:n-1]
			} else {
				result = append(result, unit)
			}
		}
		fmt.Println(string(r), "len vs min", len(result), minLen)
		if len(result) < minLen {
			minLen = len(result)
		}
	}
	fmt.Println(minLen)
}
