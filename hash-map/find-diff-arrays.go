func findDifference(nums1 []int, nums2 []int) [][]int {
	n1m := make(map[int]struct{})
	n2m := make(map[int]struct{})

	// Fill sets for nums1 and nums2
	for _, n := range nums1 {
		n1m[n] = struct{}{}
	}
	for _, n := range nums2 {
		n2m[n] = struct{}{}
	}

	var a []int
	var b []int

	// Values in nums1 but not in nums2
	for n := range n1m {
		if _, ok := n2m[n]; !ok {
			a = append(a, n)
		}
	}

	// Values in nums2 but not in nums1
	for n := range n2m {
		if _, ok := n1m[n]; !ok {
			b = append(b, n)
		}
	}

	return [][]int{a, b}
}