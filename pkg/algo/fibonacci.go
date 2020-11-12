package algo

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
