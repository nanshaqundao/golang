package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	outputTwoDimensionSlice := make([][]uint8, dy)

	fmt.Println(dx)
	fmt.Println(dy)

	for j := range outputTwoDimensionSlice {

		outputTwoDimensionSlice[j] = make([]uint8, dx)

		for i := range outputTwoDimensionSlice[j] {
			outputTwoDimensionSlice[j][i] = uint8(i * j)
		}
	}

	return outputTwoDimensionSlice
}

func main() {
	pic.Show(Pic)
}
