package algo

import "fmt"

// arr - our arrays, n = len(array), k = len of subsequence
// Complexity O(n)
func MaxSum(arr []int, k int) (int, error) {
	if len(arr) < k {
		return -1, fmt.Errorf("subsequence must be less thna array")
	}

	var (
		n         = len(arr)
		maxSum    = 0
		windowSum = 0
	)

	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	windowSum = maxSum
	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		} else {
			maxSum = maxSum
		}
	}

	return maxSum, nil
}
