func maxVowels(s string, k int) int {
	isVowel := [256]bool{}
	for _, v := range "aeiou" {
		isVowel[v] = true
	}

	var total, maxV int
	for i := 0; i < k; i++ {
		if isVowel[s[i]] {
			total++
		}
	}

	maxV = total
	for i := k; i < len(s); i++ {
		if isVowel[s[i-k]] {
			total--
		}
		if isVowel[s[i]] {
			total++
		}
		if total > maxV {
			maxV = total
		}
	}
	return maxV
}