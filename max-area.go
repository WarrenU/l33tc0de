func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		length := right - left
		currArea := min(height[left], height[right]) * length
		maxArea = max(currArea, maxArea)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}

	}
	return maxArea
}