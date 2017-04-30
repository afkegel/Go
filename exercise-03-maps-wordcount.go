package tutorial

import "strings"

// WordCount counts the words in a string.
func WordCount(s string) map[string]int {
	wordArray := strings.Fields(s)
	m := map[string]int{}

	for _, elem := range wordArray {
		switch count, ok := m[elem]; ok {
		case false:
			m[elem] = 1
		case true:
			m[elem] = count + 1
		}
	}
	return m
}
