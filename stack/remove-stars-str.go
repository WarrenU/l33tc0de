func removeStars(s string) string {
	edit := make([]rune, 0, len(s)) // preallocate with s length capacity

	for _, r := range s { // iterates over runes directly
		if r == '*' {
			if len(edit) > 0 {
				edit = edit[:len(edit)-1] // pop last rune
			}
		} else {
			edit = append(edit, r)
		}
	}

	return string(edit)
}
