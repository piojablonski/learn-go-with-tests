package main

import "math"

type Rectangle struct {
	X float64
	Y float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	A float64
	H float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.X + r.Y)
}
func (r Rectangle) Area() float64 {
	return (r.X * r.Y)
}
func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}
func (t Triangle) Area() float64 {
	return (t.A / 2 * t.H)
}
