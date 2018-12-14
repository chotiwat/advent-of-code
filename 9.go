package main

import "fmt"

type node struct {
	value   int
	cw, ccw *node
}

func (n *node) insertCW(new *node) *node {
	new.cw = n.cw
	new.ccw = n
	n.cw.ccw = new
	n.cw = new
	return new
}

func (n *node) removeSelf() *node {
	n.ccw.cw, n.cw.ccw = n.cw, n.ccw
	return n.cw
}

func (n *node) print() {
	cur := n
	for {
		fmt.Printf("%d ", cur.value)
		cur = cur.cw
		if cur == n {
			break
		}
	}
	fmt.Println()
}

func main() {
	// numPlayers, lastMarble := 10, 1618
	// numPlayers, lastMarble := 13, 7999
	// numPlayers, lastMarble := 424, 71482 // part 1
	numPlayers, lastMarble := 424, 7148200 // part 2

	var scores []int
	cur := &node{value: 0}
	cur.cw = cur
	cur.ccw = cur
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			// remove 7 ccw
			removing := cur
			for j := 0; j < 7; j++ {
				removing = removing.ccw
			}
			// fmt.Println(i, removing.value)
			scores = append(scores, i+removing.value)
			cur = removing.removeSelf()
		} else {
			// insert 1 cw
			cur = cur.cw
			cur = cur.insertCW(&node{value: i})
		}
	}

	// fmt.Println(numPlayers, scores)

	sumScores := make([]int, numPlayers)
	for i, score := range scores {
		index := (i + 1) * 23
		sumScores[index%numPlayers] += score
	}
	max := -1
	for _, sum := range sumScores {
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
