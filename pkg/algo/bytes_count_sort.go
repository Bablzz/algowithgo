package algo

func SortBytes(arr []byte) {
	if len(arr) <= 1 {
		return
	}
	var bucketArr [256]int
	for _, n := range arr {
		bucketArr[n]++
	}
	arr = arr[:0]

	for i, n := range bucketArr {
		for ; n > 0; n-- {
			arr = append(arr, byte(i))
		}
	}
}
