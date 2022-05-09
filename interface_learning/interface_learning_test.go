package interface_learning_test

import (
	"math"
	"testing"
)

// func TestInterface(t *testing.T) {
// 	t.Log("hhhhh")
// 	TestAreas := interface_learning.GetAreas()
// 	t.Log(TestAreas)
// }

// let's try this interface thing in a simpler way:
// 1. what if i run without shape interface?
// i think it runs fine
// type Shape interface {
// 	area() float64
// }

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

// func GetArea(s Shape) float64 {
// 	return s.area()
// }

func TestSimple(t *testing.T) {
	c1 := Circle{4.5}
	r1 := Rect{1, 2}
	trythisc1 := c1.area()
	trythisr1 := r1.area()
	t.Log(trythisc1)
	t.Log(trythisr1)
}
