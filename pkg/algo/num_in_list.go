package algo

func NumInList(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
