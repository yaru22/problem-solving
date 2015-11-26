package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x * y),
		uint8(x - y),
		uint8(255 - x + y),
		uint8(x + y),
	}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(&m)
}
