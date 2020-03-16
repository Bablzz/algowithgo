package algo


func SumList(list []int) (sum int) {
	for _, val := range list {
		sum = sum + val
	}
	return
}

func ReverseSum (list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + ReverseSum(list[1:])
}