package codewars

import "sort"

// kata https://www.codewars.com/kata/550498447451fbbd7600041c

func Comp(arr1 []int, arr2 []int) bool {
	if arr1 == nil || arr2 == nil || (len(arr1) != len(arr2)) {
		return false
	}
	isRight := true

	for i, v := range arr1 {
		if v < 0 {
			arr1[i] = -1 * v
		}
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i, v := range arr1 {
		if v*v != arr2[i] {
			isRight = false
			break
		}
	}
	return isRight
}
