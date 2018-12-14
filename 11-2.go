package main

import (
	"fmt"
	"math"
)

func power(serial, x, y int) int {
	rackID := x + 10
	return ((rackID*y+serial)*rackID/100)%10 - 5
}

func main() {
	const gridSize = 300
	const serial = 9221
	// const serial = 42
	var powers [gridSize + 1][gridSize + 1]int
	// fmt.Println(power(42, 21, 61))

	for y := 1; y <= gridSize; y++ {
		for x := 1; x <= gridSize; x++ {
			powers[y][x] = power(serial, x, y)
		}
	}

	var rowSums [gridSize + 1][gridSize + 2][gridSize + 2]int
	var colSums [gridSize + 1][gridSize + 2][gridSize + 2]int
	for size := 1; size <= gridSize; size++ {
		for y := gridSize; y >= 1; y-- {
			for x := gridSize; x >= 1; x-- {
				rowSums[size][y][x] = rowSums[size-1][y][x+1] + powers[y][x]
				colSums[size][y][x] = colSums[size-1][y+1][x] + powers[y][x]
			}
		}
	}

	var maxX, maxY, maxSize int
	maxSum := math.MinInt64
	var sums [gridSize + 1][gridSize + 2][gridSize + 2]int
	for size := 1; size <= gridSize; size++ {
		for y := gridSize; y >= 1; y-- {
			for x := gridSize; x >= 1; x-- {
				sums[size][y][x] = sums[size-1][y+1][x+1] + rowSums[size-1][y][x+1] + colSums[size-1][y+1][x] + powers[y][x]
				if sums[size][y][x] > maxSum {
					maxX, maxY, maxSize = x, y, size
					maxSum = sums[size][y][x]
				}
			}
		}
	}
	fmt.Println(maxSum)
	fmt.Printf("%d,%d,%d\n", maxX, maxY, maxSize)
}
