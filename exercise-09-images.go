package tutorial

import (
	"image"
	"image/color"
)

// Image implements the Image interface. Playground
// https://play.golang.org/p/wumbYwlBFl. Remains untested here.
type Image struct {
}

// Bounds returnes a rectangle.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 64, 48)
}

// ColorModel returns an RBGA color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the same color for all pixels (x, y).
func (i Image) At(x, y int) color.Color {
	return color.RGBA{100, 100, 255, 255}
}
