package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Capture struct {
	id, distance int
}

type Point struct {
	row, col int
}

type QueueItem struct {
	id, distance int
	p            Point
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// input, err := ioutil.ReadFile("6.in.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// // input := "+1\r\n-2\n+3\n+1\n"
	var points []Point
	var q []QueueItem
	scanner := bufio.NewScanner(os.Stdin)
	minRow, minCol := math.MaxInt64, math.MaxInt64
	maxRow, maxCol := math.MinInt64, math.MinInt64
	nextId := 0
	for scanner.Scan() {
		line := scanner.Text()
		p := Point{}
		fmt.Sscanf(line, "%d, %d", &p.row, &p.col)
		fmt.Println(p)
		points = append(points, p)
		q = append(q, QueueItem{
			id: nextId,
			p:  p,
		})
		nextId++
		minRow, minCol = min(minRow, p.row), min(minCol, p.col)
		maxRow, maxCol = max(maxRow, p.row), max(maxCol, p.col)
	}

	infinites := map[int]bool{}
	counts := map[int]int{}
	m := map[Point]*Capture{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if capture, ok := m[cur.p]; ok {
			if capture.distance == cur.distance && capture.id != cur.id {
				// tie
				counts[capture.id]--
				capture.id = -1
			}
		} else {
			// empty
			// find border, infinite area, strike out
			if cur.p.row <= minRow || cur.p.row >= maxRow || cur.p.col <= minCol || cur.p.col >= maxCol {
				infinites[cur.id] = true
				continue
			}
			// capture it, increment count
			m[cur.p] = &Capture{
				id:       cur.id,
				distance: cur.distance,
			}
			counts[cur.id] = counts[cur.id] + 1
			q = append(q, QueueItem{
				id:       cur.id,
				distance: cur.distance + 1,
				p:        Point{row: cur.p.row + 1, col: cur.p.col},
			},
				QueueItem{
					id:       cur.id,
					distance: cur.distance + 1,
					p:        Point{row: cur.p.row, col: cur.p.col + 1},
				},
				QueueItem{
					id:       cur.id,
					distance: cur.distance + 1,
					p:        Point{row: cur.p.row, col: cur.p.col - 1},
				},
				QueueItem{
					id:       cur.id,
					distance: cur.distance + 1,
					p:        Point{row: cur.p.row - 1, col: cur.p.col},
				})
		}
	}

	maxCount := 0
	for _, count := range counts {
		if count >= 0 && count > maxCount {
			maxCount = count
		}
	}
	fmt.Println("max area:", maxCount)
}
