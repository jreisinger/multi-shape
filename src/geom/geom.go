package geom

import (
	"math"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (r *Rectangle) Area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	Area() float64
}

type MultiShape struct {
	Shapes []Shape
}

func (m *MultiShape) Area() float64 {
	var area float64
	for _, s := range m.Shapes {
		area += s.Area()
	}
	return area
}
