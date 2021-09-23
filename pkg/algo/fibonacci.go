package algo

import "math"

func FibRecursive(max int) int {
	if max == 0 {
		return 0
	}
	if max == 1 || max == 2 {
		return 1
	}
	return FibRecursive(max-1) + FibRecursive(max-2)
}

func FibNonRecursive(max int) int {
	a, b := 0, 1

	for max != 0 {
		a, b = b, a+b
		max -= 1
	}

	return a
}

// Closed-form expression
func FibByBinet(max int) int64 {
	if max == 0 {
		return 0
	}
	if max == 1 || max == 2 {
		return 1
	}
	phi := (1 + math.Sqrt(5)) / 2 // golden ratio 1,61803
	divide := math.Pow(phi, float64(max))
	divider := math.Sqrt(5)
	return int64(divide / divider)
}
