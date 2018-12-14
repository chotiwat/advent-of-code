package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count2, count3 := 0, 0
	for scanner.Scan() {
		id := scanner.Text()
		if len(id) == 0 {
			continue
		}
		freq := [26]int{}
		for _, ch := range id {
			freq[ch-'a']++
		}
		var has2, has3 bool
		for i := 0; i < 26; i++ {
			if freq[i] == 2 {
				has2 = true
			} else if freq[i] == 3 {
				has3 = true
			}
		}
		if has2 {
			count2++
		}
		if has3 {
			count3++
		}
	}
	fmt.Println(count2 * count3)
}
