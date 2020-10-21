package algo

func FindTwoSum(numbers []int, sum int) (int, int) {

	for i, val := range numbers {
		for j, v := range numbers {
			if i == j {
				continue
			}
			if val+v == sum {
				return i, j
			}
		}
	}
	return -1, -1
}
