package tutorial

// Fibonacci returns a closure that computes the next Fibonacci number
// with every subsequent call.
func Fibonacci() func() int {
	a := []int{0, 1}

	return func() int {
		x := a[len(a)-1] + a[len(a)-2]
		a = append(a, x)
		return a[len(a)-3]
	}
}
