package algo

// The Answer to the question is => {2,3,5}, {-3,5,8}

func SumOfThree(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == target {
					var row []int
					row = append(row, nums[i], nums[j], nums[k])
					result = append(result, row)
				}
			}
		}
	}
	return result
}
