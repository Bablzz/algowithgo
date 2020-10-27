package algo

func Gcd(num1, num2 int) int {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}

func GCDRecur(num1, num2 int) int {
	if num2 == 0 {
		return num1
	}
	return GCDRecur(num2, num1%num2)
}
