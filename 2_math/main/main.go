package main

import (
	"fmt"
	"math"
)

type rightTriangle struct {
	base       float64
	height     float64
	hypotenuse float64
}

func (r *rightTriangle) calcHypo() {
	if r.hypotenuse == float64(0) {
		r.hypotenuse = math.Hypot(r.base, r.height)
	} else {
		return
	}
}

func (r *rightTriangle) calcArea() float64 {
	return r.base * r.height
}

func main() {
	pi := math.Pi
	rad := 7
	area := pi * math.Pow(float64(rad), 2)
	fmt.Println(math.Round(area) == 154)

	// Assignment need to be done to pointer but assignment address is also valid in Go
	tri := &rightTriangle{}
	// To show two ways of assigning, written for base, height in different ways
	(*tri).base = 4.0
	tri.height = 3.0
	(*tri).calcHypo()
	fmt.Println(tri.calcArea())

	x := math.Inf(1)
	fmt.Println(x)

	fmt.Println(math.Trunc(-10.5) == -10, math.Trunc(10.5) == 10)
}
