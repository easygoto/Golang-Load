package oo

import "math"

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}
