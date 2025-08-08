import (
	"reflect"
	"sort"
)

func closeStrings(w1 string, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	cMap1 := make(map[rune]int)
	cMap2 := make(map[rune]int)

	for _, r := range w1 {
		cMap1[r]++
	}
	for _, r := range w2 {
		cMap2[r]++
	}

	// Check if the set of characters is the same
	if len(cMap1) != len(cMap2) {
		return false
	}
	for ch := range cMap1 {
		if _, ok := cMap2[ch]; !ok {
			return false
		}
	}

	// Compare sorted frequency lists
	s1 := make([]int, 0, len(cMap1))
	s2 := make([]int, 0, len(cMap2))

	for _, count := range cMap1 {
		s1 = append(s1, count)
	}
	for _, count := range cMap2 {
		s2 = append(s2, count)
	}

	sort.Ints(s1)
	sort.Ints(s2)

	return reflect.DeepEqual(s1, s2)
}
