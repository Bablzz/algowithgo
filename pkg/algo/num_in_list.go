package algo

func NumInList(list []int, num int) bool {
	if len(list) == 0 {
		return false
	}
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
