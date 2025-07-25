func canPlaceFlowers(fb []int, n int) bool {
	mfb := len(fb)
	if n == 0 {
		return true
	}
	for i := 0; i < mfb; i++ {
		if (i == 0 || fb[i-1] == 0) && (i == mfb-1 || fb[i+1] == 0) && fb[i] == 0 {
			fb[i] = 1
			n--
			if n == 0 {
				return true
			}
			i += 1
		}
	}
	return n == 0
}