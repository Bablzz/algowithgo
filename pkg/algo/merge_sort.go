package algo

func sort(leftArr, rightArr []int) []int {
	resultArr := make([]int, 0, len(leftArr)+len(rightArr))
	for len(leftArr) > 0 || len(rightArr) > 0 {
		if len(leftArr) == 0 {
			return append(resultArr, rightArr...)
		}
		if len(rightArr) == 0 {
			return append(resultArr, leftArr...)
		}
		if leftArr[0] <= rightArr[0] {
			resultArr = append(resultArr, leftArr[0])
			leftArr = leftArr[1:]
		} else {
			resultArr = append(resultArr, rightArr[0])
			rightArr = rightArr[1:]
		}
	}
	return resultArr
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return sort(left, right)
}
