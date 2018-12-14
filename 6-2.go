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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

	threshold := 10000
	count := 0
	visited := map[Point]bool{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if !visited[cur.p] {
			visited[cur.p] = true
			// sum distance
			sum := 0
			for _, point := range points {
				sum += abs(point.row-cur.p.row) + abs(point.col-cur.p.col)
			}
			if sum < threshold || (cur.p.row >= minRow && cur.p.col >= minCol && cur.p.row <= maxRow && cur.p.col <= maxCol) {
				if sum < threshold {
					count++
				}
				q = append(q, QueueItem{
					id: cur.id,
					p:  Point{row: cur.p.row + 1, col: cur.p.col},
				},
					QueueItem{
						id: cur.id,
						p:  Point{row: cur.p.row, col: cur.p.col + 1},
					},
					QueueItem{
						id: cur.id,
						p:  Point{row: cur.p.row, col: cur.p.col - 1},
					},
					QueueItem{
						id: cur.id,
						p:  Point{row: cur.p.row - 1, col: cur.p.col},
					})
			}
		}
	}

	fmt.Println("safe area:", count)
}
