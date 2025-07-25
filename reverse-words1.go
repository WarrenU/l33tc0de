import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}