package main

import "fmt"

func main() {
	recipes := []int{3, 7}
	cur1, cur2 := 0, 1
	n := 440231

	for len(recipes) < n+10 {
		sum := recipes[cur1] + recipes[cur2]
		if sum >= 10 {
			recipes = append(recipes, 1)
			sum %= 10
		}
		recipes = append(recipes, sum)
		cur1 = (cur1 + 1 + recipes[cur1]) % len(recipes)
		cur2 = (cur2 + 1 + recipes[cur2]) % len(recipes)
	}
	for i := n; i < n+10; i++ {
		fmt.Printf("%d", recipes[i])
	}
	fmt.Println()
}
