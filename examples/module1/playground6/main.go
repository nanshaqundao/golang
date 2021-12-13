package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func Pic(dx, dy int) [][]uint8 {

	outputTwoDimensionSlice := make([][]uint8, dy)

	for j := range outputTwoDimensionSlice {

		outputTwoDimensionSlice[j] = make([]uint8, dx)

		for i := range outputTwoDimensionSlice[j] {
			outputTwoDimensionSlice[j][i] = uint8(i ^ j - i*j + i*i + j*j)
		}
	}

	return outputTwoDimensionSlice
}

func main() {
	pic.Show(Pic)
	m := &Image{32, 32}
	pic.ShowImage(m)

}
