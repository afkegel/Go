package tutorial

import "fmt"

type errNegSqrt struct {
	f float64
}

// Error is an implementation of the error interface on the errNegSqrt type.
func (e *errNegSqrt) Error() string {
	return fmt.Sprintf("Negative number: %v", float64(e.f))
}

// NewSqrtFuncDecorator is a constructor for a decorated square root function
// func(float64) (floa64, int).
func NewSqrtFuncDecorator(f func(float64) (float64, int)) SqrtFuncDecorated {
	return sqrtFuncer(f).decorate()
}

// SqrtFuncDecorated is the exported type for an error decorated square root
// function defined as func(float64)(float64, int).
type SqrtFuncDecorated func(float64) (float64, int, error)

type sqrtFuncer func(float64) (float64, int)

func (fn sqrtFuncer) decorate() SqrtFuncDecorated {
	return func(f float64) (float64, int, error) {
		if f < 0 {
			return f, 0, &errNegSqrt{f}
		}
		result, i := fn(f)
		return result, i, nil
	}
}
