package tutorial

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNewtonSqrtIteration(t *testing.T) {
	// Tests Newton's sqrt by assigning an iterator to the interface and
	// running it's sqrt method.
	rand.Seed(time.Now().UTC().UnixNano())

	test := make([]float64, 10)
	result := make([]float64, 10)
	expected := make([]float64, 10)
	var N Newton = Iterator(10)

	for i := 0; i < 10; i++ {
		test[i] = float64(rand.Float64())
		result[i], _ = N.sqrt(test[i])
		expected[i] = math.Sqrt(test[i])
		if result[i]-expected[i] > .001 {
			t.Errorf("ComputeNewtonSqrt failed!")
		}
	}
}

func TestNewtonSqrtOptimisation(t *testing.T) {
	// Tests Newton's sqrt by assigning an optimiser to the interface and
	// running it's sqrt method.
	rand.Seed(time.Now().UTC().UnixNano())
	test := make([]float64, 10)
	result := make([]float64, 10)
	expected := make([]float64, 10)
	var N Newton = Optimiser(.001)

	for i := 0; i < 10; i++ {
		test[i] = float64(rand.Float64())
		result[i], _ = N.sqrt(test[i])
		expected[i] = math.Sqrt(test[i])
		if result[i]-expected[i] > .001 {
			t.Errorf("ComputeNewtonSqrt2 failed!")
		}
	}
}
