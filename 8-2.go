package main

import "fmt"

type node struct {
	children []*node
	metadata []int
	value    int
}

func read() *node {
	var numChildren, numMetadata int
	fmt.Scanf("%d %d", &numChildren, &numMetadata)
	n := &node{}
	for i := 0; i < numChildren; i++ {
		n.children = append(n.children, read())
	}
	for i := 0; i < numMetadata; i++ {
		var m int
		fmt.Scanf("%d", &m)
		n.metadata = append(n.metadata, m)
	}

	for _, m := range n.metadata {
		if numChildren == 0 {
			n.value += m
		} else if m <= numChildren {
			n.value += n.children[m-1].value
		}
	}
	return n
}

func main() {
	root := read()
	fmt.Println(root.value)
}
