package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ids []string
	for scanner.Scan() {
		id := scanner.Text()
		if len(id) == 0 {
			continue
		}
		ids = append(ids, id)
	}

	for i := 0; i+1 < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			mismatch := 0
			mismatchIndex := -1
			for k, c := range ids[i] {
				if c != rune(ids[j][k]) {
					mismatch++
					mismatchIndex = k
				}
				if mismatch > 1 {
					break
				}
			}
			if mismatch == 1 {
				fmt.Printf("%s%s\n", ids[i][0:mismatchIndex], ids[i][mismatchIndex+1:])
				return
			}
		}
	}
}
