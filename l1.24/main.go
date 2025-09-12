package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(o *Point) float64 {
	return math.Sqrt(math.Pow(p.x-o.x, 2) + math.Pow(p.y-o.y, 2))
}

func main() {
	point := NewPoint(-3, -4)
	other := NewPoint(1.5, 2)

	fmt.Printf("point 1: x = %.2f, y = %.2f\n", point.x, point.y)
	fmt.Printf("point 2: x = %.2f, y = %.2f\n", other.x, other.y)
	fmt.Printf("distance: %.2f", point.Distance(other))
}
