func uniqueOccurrences(arr []int) bool {
	visited := make(map[int]int, len(arr))
	for _, n := range arr {
		visited[n]++
	}

	seenOccurrence := make(map[int]struct{}, len(visited))
	for _, count := range visited {
		if _, exists := seenOccurrence[count]; exists {
			return false
		}
		seenOccurrence[count] = struct{}{}
	}

	return true
}