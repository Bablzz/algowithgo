package algo

func InterpolationSearch(elements []int, element int) (bool, int) {
	var (
		mid  int
		low  int
		high int
	)

	low = 0
	high = len(elements) - 1

	if element == elements[low] {
		return true, low
	}
	if element == elements[high] {
		return true, high
	}
	for elements[low] < element && elements[high] > element {
		mid = low + ((element-elements[low])*(high-low))/(elements[high]-elements[low])

		if element > elements[mid] {
			low = mid + 1
		}
		if element < elements[mid] {
			high = mid - 1
		} else {
			return true, mid
		}
	}

	return false, -1
}
