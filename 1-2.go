package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type empty struct{}

var e empty

func main() {
	input, err := ioutil.ReadFile("1.in.txt")
	if err != nil {
		panic(err)
	}
	// input := "+1\r\n-2\n+3\n+1\n"
	tokens := strings.Split(string(input), "\n")
	sum := 0
	visited := map[int]empty{
		0: e,
	}
	for {
		for _, token := range tokens {
			token = strings.TrimSpace(token)
			if len(token) > 0 {
				sign, numString := token[0], token[1:]
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}
				if sign == '+' {
					sum += num
				} else {
					sum -= num
				}
				if _, ok := visited[sum]; ok {
					fmt.Println(sum)
					return
				}
				visited[sum] = e
			}
		}
	}
}
