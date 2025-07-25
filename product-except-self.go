func productExceptSelf(nums []int) []int {
	total := len(nums)
	left := make([]int, total)
	right := make([]int, total)
	left[0] = 1
	right[total-1] = 1
	for i := 1; i < total; i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	for j := total - 2; j > -1; j-- {
		right[j] = nums[j+1] * right[j+1]
	}
	for i := 0; i < total; i++ {
		nums[i] = left[i] * right[i]
	}
	return nums
}