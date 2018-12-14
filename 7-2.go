package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type node struct {
	exists     bool
	in         int
	dependents int64
}

type completion struct {
	id, time, numWorkers int
}

type completionHeap []*completion

func (h completionHeap) Len() int           { return len(h) }
func (h completionHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h completionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *completionHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*completion))
}

func (h *completionHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

	// event types
	// worker starts task
	// worker finishes task -> starts task

	h := &completionHeap{}
	heap.Init(h)
	numWorkers := 5
	baseDuration := 61 // 60 + 1 for zero-indexed id

	// for i := 0; i < numWorkers; i++ {
	h.Push(&completion{id: -1, time: 0})
	// }

	time := 0
	for h.Len() > 0 {
		c := heap.Pop(h).(*completion)
		fmt.Println(c.id, c.time)
		if c.id >= 0 {
			for j := 0; j < 26; j++ {
				var mask int64 = 1 << uint(j)
				if mask&nodes[c.id].dependents == mask {
					nodes[j].in--
				}
			}
			numWorkers++
		}
		time = c.time

		working := 0
		for w := 0; w < numWorkers; w++ {
			for i, node := range nodes {
				if !node.exists {
					continue
				}
				if node.in == 0 {
					// order = append(order, rune('A'+i))
					heap.Push(h, &completion{id: i, time: time + baseDuration + i})
					nodes[i].exists = false
					working++
					break
				}
			}
		}
		numWorkers -= working
	}
	fmt.Println(time)
}
