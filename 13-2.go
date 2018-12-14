package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type direction int
type intPair [2]int

const (
	up direction = iota
	down
	left
	right
)

type cart struct {
	row, col, counter int
	direction         direction
}

func (c *cart) String() string {
	return fmt.Sprintf("(%d %d %d %d)", c.row, c.col, c.direction, c.counter)
}

func (c *cart) straight() {
	// fmt.Println("straight", c)
	switch c.direction {
	case up:
		c.row--
	case down:
		c.row++
	case left:
		c.col--
	case right:
		c.col++
	}
}

func (c *cart) turnLeft() {
	// fmt.Println("left", c)
	switch c.direction {
	case up:
		c.direction = left
	case left:
		c.direction = down
	case down:
		c.direction = right
	case right:
		c.direction = up
	}
}

func (c *cart) turnRight() {
	// fmt.Println("right", c)
	switch c.direction {
	case up:
		c.direction = right
	case right:
		c.direction = down
	case down:
		c.direction = left
	case left:
		c.direction = up
	}
}

func (c *cart) move(track rune) {
	// fmt.Println("move", string(track), c)
	switch track {
	case '|':
		if c.direction != up && c.direction != down {
			panic("wrong! |")
		}
	case '-':
		if c.direction != left && c.direction != right {
			panic("wrong! -")
		}
	case '/':
		// right if updown, left if leftright
		if c.direction == up || c.direction == down {
			c.turnRight()
		} else {
			c.turnLeft()
		}
	case '\\':
		// left if updown, right if leftright
		if c.direction == up || c.direction == down {
			c.turnLeft()
		} else {
			c.turnRight()
		}
	case '+':
		if c.counter%3 == 0 {
			c.turnLeft()
		} else if c.counter%3 == 2 {
			c.turnRight()
		}
		c.counter++
	}
	c.straight()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	occupied := make(map[intPair]*cart)
	var m [][]rune
	var carts []*cart

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		row := []rune(scanner.Text())

		for col, cell := range row {
			var d direction
			var replacement rune
			switch cell {
			case '<':
				d = left
				replacement = '-'
			case '>':
				d = right
				replacement = '-'
			case '^':
				d = up
				replacement = '|'
			case 'v':
				d = down
				replacement = '|'
			default:
				continue
			}
			c := &cart{row: len(m), col: col, direction: d}
			occupied[intPair{len(m), col}] = c
			carts = append(carts, c)
			row[col] = replacement
		}
		m = append(m, row)
	}

	for len(carts) > 1 {
		for _, c := range carts {
			if c.counter == -1 {
				continue
			}
			delete(occupied, intPair{c.row, c.col})
			// fmt.Println(c.row, c.col)
			c.move(m[c.row][c.col])
			if occupied[intPair{c.row, c.col}] != nil {
				another := occupied[intPair{c.row, c.col}]
				// crash
				fmt.Printf("removing %d,%d %d,%d\n", c.col, c.row, another.col, another.row)
				// mark both as crashed
				c.counter = -1
				another.counter = -1
				// remove occupied
				delete(occupied, intPair{c.row, c.col})
				continue
			}
			occupied[intPair{c.row, c.col}] = c
		}
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].counter == -1 {
				return false
			}
			if carts[j].counter == -1 {
				return true
			}
			if carts[i].row == carts[j].row {
				return carts[i].col < carts[j].col
			}
			return carts[i].row < carts[j].row
		})
		lastIndex := len(carts) - 1
		for ; lastIndex >= 0 && carts[lastIndex].counter == -1; lastIndex-- {
			// pass
		}
		carts = carts[0 : lastIndex+1]
		if len(carts)%2 == 0 {
			panic("!!!")
		}
	}
	fmt.Printf("%d,%d\n", carts[0].col, carts[0].row)
}
