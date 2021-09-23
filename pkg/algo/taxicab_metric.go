package algo

import "math"

type Point struct {
	X, Y float64
}

func Distance(p1 Point, p2 Point) int {
	return int(math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y))
}
