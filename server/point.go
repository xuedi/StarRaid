package server

import (
	"math"
)

type point struct {
	x int64
	y int64
}

func (p point) Distance(p2 point) int64 {
	first := math.Pow(float64(p2.x-p.x), 2)
	second := math.Pow(float64(p2.y-p.y), 2)
	return int64(math.Sqrt(first + second))
}
