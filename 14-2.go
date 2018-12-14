package main

import "fmt"

func check(recipes, target []int) bool {
	if len(recipes) >= len(target) {
		match := true
		n, m := len(recipes), len(target)
		for i := 1; i <= len(target); i++ {
			if recipes[n-i] != target[m-i] {
				match = false
				break
			}
		}
		return match
	}
	return false
}

func main() {
	recipes := []int{3, 7}
	cur1, cur2 := 0, 1
	target := []int{4, 4, 0, 2, 3, 1}

	for {
		sum := recipes[cur1] + recipes[cur2]
		if sum >= 10 {
			recipes = append(recipes, 1)
			sum %= 10
			if check(recipes, target) {
				fmt.Println(len(recipes) - len(target))
				return
			}
		}
		recipes = append(recipes, sum)
		cur1 = (cur1 + 1 + recipes[cur1]) % len(recipes)
		cur2 = (cur2 + 1 + recipes[cur2]) % len(recipes)

		if check(recipes, target) {
			fmt.Println(len(recipes) - len(target))
			return
		}
	}
}
