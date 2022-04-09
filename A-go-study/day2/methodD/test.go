package main

import "fmt"

func main() {
	var f = 2.0
	p := Point{
		X: 1,
		Y: 3,
	}
	p.ScaleBy(f)
	fmt.Println(p)

}

type Point struct {
	X float64
	Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X += factor
	p.Y *= factor
}
