package algo

import "math/rand"

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		fresh := make([]int, 0, len(arr))
		pivot := arr[rand.Intn(len(arr))]

		less := findLess(arr, pivot)
		greater := findMore(arr, pivot)
		fresh = append(fresh, Quicksort(less)...)
		fresh = append(fresh, pivot)
		fresh = append(fresh, Quicksort(greater)...)
		return fresh
	}
}

func findLess(arr []int, p int) (new []int) {
	for _, v := range arr {
		if v < p {
			new = append(new, v)
		}
	}
	return
}

func findMore(arr []int, p int) (new []int) {
	for _, v := range arr {
		if v > p {
			new = append(new, v)
		}
	}
	return
}
