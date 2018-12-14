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
	const squareSize = 3
	const serial = 9221
	var powers [gridSize + 1][gridSize + 1]int
	// fmt.Println(power(42, 21, 61))

	for y := 1; y <= gridSize; y++ {
		for x := 1; x <= gridSize; x++ {
			powers[y][x] = power(serial, x, y)
		}
	}

	var rowSums [gridSize + 2][gridSize + 2]int
	for y := 1; y <= gridSize; y++ {
		for x := gridSize; x >= 1; x-- {
			rowSums[y][x] = rowSums[y][x+1] + powers[y][x]
			if x+squareSize <= gridSize {
				rowSums[y][x] -= powers[y][x+squareSize]
			}
		}
	}

	var maxX, maxY int
	maxSum := math.MinInt64
	var sums [gridSize + 2][gridSize + 2]int
	for y := gridSize; y >= 1; y-- {
		for x := 1; x <= gridSize; x++ {
			sums[y][x] = sums[y+1][x] + rowSums[y][x]
			if y+squareSize <= gridSize {
				sums[y][x] -= rowSums[y+squareSize][x]
			}
			if sums[y][x] > maxSum {
				maxX, maxY = x, y
				maxSum = sums[y][x]
			}
		}
	}
	fmt.Printf("%d,%d\n", maxX, maxY)
}
