package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	c()
}

func c() {
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	cp.Color = color.RGBA{
		R: 1,
		G: 23,
		B: 45,
		A: 78,
	}

	fmt.Printf("%+v\n", cp)
}

func add(i, j int) (a int) {
	a = i + j
	return
}
