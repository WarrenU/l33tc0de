func moveZeroes(nums []int) {
	ip := 0 // insert index/position (ip)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ip] = nums[i]
			ip += 1
		}
	}
	for ; ip < len(nums); ip++ {
		nums[ip] = 0
	}
}