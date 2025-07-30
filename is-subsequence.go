func isSubsequence(s string, t string) bool {
	ti := 0 // total index
	sr := []rune(s)
	tr := []rune(t)
	totalSr := len(sr)
	for _, c := range tr {
		if ti < totalSr && sr[ti] == c {
			ti += 1
		}
	}
	visited := len(s)
	return ti == visited
}