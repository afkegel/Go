package tutorial

// Pic returns png encoded data with an IMAGE tag that get parsed by the Go
// playground. https://play.golang.org/p/blLJbY6bFq
// It remains untested here.
func Pic(dx, dy int) [][]uint8 {

	a := make([][]uint8, dy)

	// create slices within a
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// fill the slices with content
	for i := range a {
		for j := range a[i] {
			a[i][j] = uint8(i * j)
		}
	}

	return a
}
