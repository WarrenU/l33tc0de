func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	// Initial sum of the first window
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	// Slide the window
	for i := k; i < n; i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}