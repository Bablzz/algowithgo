package algo

import "fmt"

func Bs(arr []int, elem int) int {
	low := 0
	high := len(arr)

	for low != high {

		mid := (low + high) / 2
		if elem == arr[mid] {
			fmt.Printf("Find element %d, its index %d", arr[mid], mid)
			return mid
		} else if elem < arr[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
