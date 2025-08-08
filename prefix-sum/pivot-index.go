

func pivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	leftSum := 0
	for i, n := range nums {
		if leftSum == total-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}