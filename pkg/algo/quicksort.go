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

func Quicksort2(arr []int, less int, greater int) {
	if less < greater {
		part := sorting(arr, less, greater)
		Quicksort2(arr, less, greater-1)
		Quicksort2(arr, part+1, greater)
	}
}

func sorting(arr []int, less int, greater int) int {
	mid := arr[greater]
	i := less

	for j := less; j < greater; j++ {
		if arr[j] <= mid {
			swap(&arr[i], &arr[j])
			i += 1
		}
	}
	swap(&arr[i], &arr[greater])
	return i
}

func swap(elem1 *int, elem2 *int) {
	val := *elem1
	*elem1 = *elem2
	*elem2 = val
}
