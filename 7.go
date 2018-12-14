package main

import (
	"fmt"
	"strconv"
)

type node struct {
	exists     bool
	in         int
	dependents int64
}

func print(nodes []node) {
	for i, node := range nodes {
		if !node.exists {
			continue
		}
		fmt.Printf("%d %d %26s\n", i, node.in, strconv.FormatInt(node.dependents, 2))
	}
	fmt.Println("------")
	return
}

func main() {
	var src, dst rune
	var err error
	nodes := make([]node, 26)
	for {
		_, err = fmt.Scanf("Step %c must be finished before step %c can begin.\n", &src, &dst)
		if err != nil {
			break
		}
		srcIndex, dstIndex := src-'A', dst-'A'
		nodes[srcIndex].exists = true
		nodes[dstIndex].exists = true
		nodes[dstIndex].in++
		nodes[srcIndex].dependents |= 1 << uint(dstIndex)
	}

	// print(nodes)

	var order []rune
	for k := 0; k < 26; k++ {
		for i, node := range nodes {
			if !node.exists {
				continue
			}
			if node.in == 0 {
				order = append(order, rune('A'+i))
				nodes[i].exists = false
				for j := 0; j < 26; j++ {
					var mask int64 = 1 << uint(j)
					if mask&node.dependents == mask {
						nodes[j].in--
					}
				}
				break
			}
		}
		// print(nodes)
	}
	fmt.Println(string(order))
}
