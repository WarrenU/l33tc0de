func largestAltitude(gain []int) int {
	maxAlt := 0
	currAlt := 0
	for _, g := range gain {
		currAlt += g
		maxAlt = max(currAlt, maxAlt)
	}
	return maxAlt
}