func mergeAlternately(word1 string, word2 string) string {
	res := ""
	w1L := len(word1)
	w2L := len(word2)
	maxLen := w1L
	if w2L > w1L {
		maxLen = w2L
	}
	for i := range maxLen {
		if i < w1L {
			res += string(word1[i])
		}
		if i < w2L {
			res += string(word2[i])
		}
	}
	return res
}