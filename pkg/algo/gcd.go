package algo

func Gcd(num1, num2 int) int {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}
