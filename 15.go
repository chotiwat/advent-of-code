package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

const (
	kindWall   = '#'
	kindElf    = 'E'
	kindGoblin = 'G'
	kindEmpty  = '.'
)

type pos struct {
	row, col int
}

type unit struct {
	kind rune
	hp   int
	pos
}

var wall = &unit{kind: kindWall}

// An unitHeap is a min-heap of ints.
type unitHeap []*unit

func (h unitHeap) Len() int { return len(h) }
func (h unitHeap) Less(i, j int) bool {
	return h[i].row == h[j].row && h[i].col < h[j].col || h[i].row < h[j].row
}
func (h unitHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *unitHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*unit))
}

func (h *unitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type path struct {
	distance int
	unit     *unit
}

type cell struct {
	occupant *unit
	// reachables []path
}

type battlefield [][]cell // because map is taken

type queueItem struct {
	pos
	distance int
}

func oob(b battlefield, p pos) bool {
	return p.row < 0 || p.row >= len(b) || p.col < 0 && p.col >= len(b[0]) // assume at least one row
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hamilton(a, b pos) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func (u *unit) fill(b battlefield, dst pos, maxDistance int, visited map[pos]*path) int {
	var q []queueItem
	q = append(q, queueItem{pos: u.pos, distance: 0})

	for len(q) > 0 {
		cur := q[0]
		if cur.pos == dst {
			return cur.distance
		}
		q = q[1:]
		c := &(b[cur.row][cur.col])
		if oob(b, cur.pos) || c.occupant != nil || visited[cur.pos] != nil && visited[cur.pos].distance <= cur.distance {
			continue
		}
		// c.reachables = append(c.reachables, path{distance: cur.distance, unit: u})
		if cur.distance == maxDistance {
			break
		}
		visited[cur.pos] = &path{distance: cur.distance, unit: u}
		// heuristic
		if cur.distance+hamilton(cur.pos, dst) >= maxDistance {
			continue
		}
		q = append(q,
			queueItem{
				pos: pos{
					row: cur.row + 1,
					col: cur.col,
				},
				distance: cur.distance + 1,
			},
			queueItem{
				pos: pos{
					row: cur.row,
					col: cur.col + 1,
				},
				distance: cur.distance + 1,
			},
			queueItem{
				pos: pos{
					row: cur.row - 1,
					col: cur.col,
				},
				distance: cur.distance + 1,
			},
			queueItem{
				pos: pos{
					row: cur.row,
					col: cur.col - 1,
				},
				distance: cur.distance + 1,
			})
	}
	return maxDistance
}

func adjacents(p pos) []pos {
	return []pos{
		{row: p.row - 1, col: p.col},
		{row: p.row, col: p.col - 1},
		{row: p.row, col: p.col + 1},
		{row: p.row + 1, col: p.col},
	}
}

var inifinityPath = &path{distance: math.MaxInt64}

func (u *unit) pick(b battlefield) []*unit {
	adj := adjacents(u.pos)
	var targets []*unit
	for _, p := range adj {
		if oob(b, p) {
			continue
		}
		c := &(b[p.row][p.col])
		if c.occupant != nil && c.occupant != wall && c.occupant.kind != u.kind {
			targets = append(targets, c.occupant)
		}
	}
	return targets
}

func (u *unit) move(b battlefield, units []*unit) {
	if len(u.pick(b)) > 0 {
		return
	}

	targets := &unitHeap{}
	heap.Init(targets)
	visited := map[pos]*path{}
	for _, t := range units {
		// if u.kind != t.kind {
		adj := adjacents(t.pos)
		for _, p := range adj {
			if visited[p] == nil {
				visited[p] = inifinityPath
				heap.Push(targets, &unit{kind: t.kind, pos: p})
			}
		}
		// }
	}

	cutoff := math.MaxInt64
	for targets.Len() > 0 {
		t := heap.Pop(targets).(*unit)
		cutoff = t.fill(b, u.pos, cutoff, visited)
	}

	minDistance, minPos := inifinityPath.distance, pos{}
	adj := adjacents(u.pos)
	for _, p := range adj {
		if visited[p] != nil && visited[p].distance < minDistance {
			minDistance, minPos = visited[p].distance, p
		}
	}
	if minDistance != inifinityPath.distance {
		// move
		// fmt.Printf("move %c from %d,%d to %d,%d\n", u.kind, u.row, u.col, minPos.row, minPos.col)
		b[u.row][u.col].occupant = nil
		u.pos = minPos
		b[u.row][u.col].occupant = u
	}
}

func (u *unit) attack(b battlefield) {
	targets := u.pick(b)
	if len(targets) == 0 {
		return
	}
	// find min hp
	minHP := math.MaxInt64
	var target *unit
	for _, t := range targets {
		if t.hp < minHP {
			minHP, target = t.hp, t
		}
	}
	target.hp -= 3
	if target.hp <= 0 {
		b[target.row][target.col].occupant = nil
	}
}

func (u *unit) enemies(units []*unit) []*unit {
	var targets []*unit
	for _, t := range units {
		if t.kind != u.kind && t.hp > 0 {
			targets = append(targets, t)
		}
	}
	return targets
}

func main() {
	var b battlefield
	scanner := bufio.NewScanner(os.Stdin)

	h := &unitHeap{}
	heap.Init(h)

	var units []*unit

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var row []cell
		for col, ch := range line {
			var u *unit
			if ch == kindWall {
				u = wall
			} else if ch != kindEmpty {
				u = &unit{
					kind: ch,
					hp:   200,
					pos: pos{
						row: len(b),
						col: col,
					},
				}
				heap.Push(h, u)
				units = append(units, u)
			}
			row = append(row, cell{occupant: u})
		}
		b = append(b, row)
	}

	for turn := 1; true; turn++ {
		nextTurn := &unitHeap{}
		heap.Init(nextTurn)
		for h.Len() > 0 {
			u := heap.Pop(h).(*unit)
			if u.hp <= 0 {
				continue
			}
			targets := u.enemies(units)
			if len(targets) == 0 {
				// game ends
				sumHP := 0
				for _, t := range units {
					if t.kind == u.kind && t.hp > 0 {
						sumHP += t.hp
					}
				}
				fmt.Println(sumHP, turn)
				fmt.Println(sumHP * (turn - 1))
				return
			}
			u.move(b, targets)
			u.attack(b)
			heap.Push(nextTurn, u)
		}
		h = nextTurn
	}
}
