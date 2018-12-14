package main

import "fmt"

type node struct {
	children []*node
	metadata []int
}

func read() (*node, int) {
	var numChildren, numMetadata int
	fmt.Scanf("%d %d", &numChildren, &numMetadata)
	sumMetadata := 0
	n := &node{}
	for i := 0; i < numChildren; i++ {
		child, sum := read()
		sumMetadata += sum
		n.children = append(n.children, child)
	}
	for i := 0; i < numMetadata; i++ {
		var m int
		fmt.Scanf("%d", &m)
		sumMetadata += m
		n.metadata = append(n.metadata, m)
	}
	return n, sumMetadata
}

func main() {
	_, sum := read()
	fmt.Println(sum)
}
