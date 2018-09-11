package main

import (
	"fmt"
)

type area interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	var t triangle
	var s square
	t.height = 6
	t.base = 5
	s.sideLength = 5
	t.getArea()
	s.getArea()
	printArea(t)
	printArea(s)

}

func (t triangle) getArea() float64 {
	var height = t.height
	var base = t.base
	return height * base
}

func (s square) getArea() float64 {

	var sidelength = s.sideLength
	return sidelength * sidelength
}

func printArea(a area) {
	fmt.Println(a.getArea())
}
