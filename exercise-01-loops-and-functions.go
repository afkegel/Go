package tutorial

import "math"

// Newton is an exported interface that has a sqrter
type Newton interface {
	sqrter
}

// Iterator sets the number of iterations for the Newton square root calculation.
type Iterator int

// Optimiser sets the optimisation criteria for which the iterative Newton square root calculation stops.
type Optimiser float64

// sqrt computes the Newton sqrt.
func (it Iterator) sqrt(f float64) (float64, int) {
	fn := f
	i := 0
	for i <= int(it) {
		fn = fn - (((fn * fn) - f) / (2 * fn))
		i++
	}
	return fn, i
}

func (op Optimiser) sqrt(f float64) (float64, int) {
	fn := f
	fo := f + 1 // so that it passes entry to the loop
	delta := float64(op)
	i := 0
	for math.Abs(fo-fn) > delta {
		fn, fo = fn-(((fn*fn)-f)/(2*fn)), fn
		i++
	}
	return fn, i
}

type sqrter interface {
	sqrt(float64) (float64, int)
}
