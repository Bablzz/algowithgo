package algo

func LCM(a, b int) int {
	return a / Gcd(a, b) * b
}
