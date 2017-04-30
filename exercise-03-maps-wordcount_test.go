package tutorial

import "testing"

func TestWordCount(t *testing.T) {
	// Tests word count by comparing the computed answer to the correct answer
	// known beforehand.
	result := WordCount("hello hello you you")
	expected := map[string]int{"hello": 2, "you": 2}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("WordCount failed!")
		}
	}
}
