package algo

func SumN(n int) int {
	return n * (n + 1) / 2
}

func SumEven(n int) int {
	return n * (n + 2) / 4
}

func SumOdd(n int) int {
	return (n + 1) * (n + 1) / 4
}

func MissNumber(n []int) int {
	sum := SumN(len(n))
	var missSum int
	for _, v := range n {
		missSum += v
	}
	return missSum - sum
}
