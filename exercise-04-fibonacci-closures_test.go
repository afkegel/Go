package tutorial

import (
	"math/rand"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// Tests Fibonacci by creating the Fibonacci object, calling it a number of
	// times n in the interval [0, 10] and comparing the result against the
	// correct number fn.
	test := Fibonacci()
	result := []int{}
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	max := rand.Intn(10)
	for i := 0; i < max; i++ {
		result = append(result, test())
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Fibonacci failed!")
		}
	}
}
