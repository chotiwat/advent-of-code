package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

type point struct {
	x, y   int
	vx, vy int
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

func add(a point, b point) point {
	p := a
	p.x += b.x
	p.y += b.y
	return p
}

func scaleInt(x int, s float64) int {
	return int(math.Round(float64(x) / s))
}

func main() {
	var points []*point
	minPoint := point{x: math.MaxInt64, y: math.MaxInt64}
	maxPoint := point{x: math.MinInt64, y: math.MinInt64}
	for {
		p := &point{}
		_, err := fmt.Scanf("position=<%d,  %d> velocity=<%d, %d>\n", &p.x, &p.y, &p.vx, &p.vy)
		if err != nil {
			break
		}
		points = append(points, p)
	}

	scale := 1.0
	iterations := 200
	start := 10920

	for t := start; t < start+iterations; t++ {
		var transformed []*point
		for _, p := range points {
			p = &point{
				x: scaleInt(p.x, scale) + p.vx*t,
				y: scaleInt(p.y, scale) + p.vy*t,
			}
			transformed = append(transformed, p)

			minPoint.x, minPoint.y = min(minPoint.x, p.x), min(minPoint.y, p.y)
			maxPoint.x, maxPoint.y = max(maxPoint.x, p.x), max(maxPoint.y, p.y)
		}
		gray := image.NewGray(image.Rect(0, 0, maxPoint.x-minPoint.x+1, maxPoint.y-minPoint.y+1))
		for _, p := range transformed {
			gray.SetGray(p.x-minPoint.x, p.y-minPoint.y, color.Gray{Y: 255})
		}
		f, err := os.Create(fmt.Sprintf("10-iter-%d.gif", t))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := gif.Encode(f, gray, nil); err != nil {
			panic(err)
		}
	}
}
