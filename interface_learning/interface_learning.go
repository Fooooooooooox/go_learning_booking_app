package interface_learning

import (
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func GetArea(s Shape) float64 {
	return s.area()
}

func GetAreas() []float64 {
	c1 := Circle{4.5}
	r1 := Rect{1, 2}
	Shapes := []Shape{c1, r1}
	var areas []float64
	for _, shape := range Shapes {
		areas = append(areas, GetArea(shape))
	}
	return areas
}
