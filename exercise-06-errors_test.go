package tutorial

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestError(t *testing.T) {
	err := &errNegSqrt{f: -3}
	if fmt.Sprint(err) != "Negative number: -3" {
		t.Error("Error implementation of errNegSqrt failed.")
	}
}

func TestError_Wrapper(t *testing.T) {
	// Tests the error wrapper by computing square roots on random numbers on
	// the interval [-0.5, 0.5] and saving the error in an error slice.  The
	// switch then extracts the saved float64 from the error and checks whether
	// it's greater equal to zero, which should not happen.
	var N Newton = Optimiser(10)
	fn := NewSqrtFuncDecorator(N.sqrt)

	rand.Seed(time.Now().UTC().UnixNano())
	test := make([]float64, 10)
	result := make([]float64, 10)
	expected := make([]float64, 10)
	num := make([]int, 10)
	err := make([]error, 10)

	for i := 0; i < 10; i++ {
		test[i] = math.Log(float64(rand.Float64())) + 0.5
		result[i], num[i], err[i] = fn(test[i])
		expected[i] = math.Sqrt(test[i])
	}

	for _, r := range err {
		switch {
		case r == nil:
			return
		case r.(*errNegSqrt).f > 0: // error interface value, does not have this field if <nil>
			t.Errorf("ErrorDecorator failed! Err: %v", r)
		default:
			return
		}
	}
}
