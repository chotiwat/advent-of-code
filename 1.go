package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("1.in.txt")
	if err != nil {
		panic(err)
	}
	// input := "+1\r\n-2\n+3\n+1\n"
	tokens := strings.Split(string(input), "\n")
	sum := 0
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
		}
	}
	fmt.Println(sum)
}
