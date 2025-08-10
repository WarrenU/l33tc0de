import "slices"

func reverseVowels(s string) string {
	l := 0
	r := len(s) - 1
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	chars := []byte(s)
	for l < r {
		if !slices.Contains(vowels, chars[l]) {
			l++
		} else if !slices.Contains(vowels, chars[r]) {
			r--
		} else {
			chars[l], chars[r] = chars[r], chars[l]
			l++
			r--
		}
	}
	return string(chars)
}