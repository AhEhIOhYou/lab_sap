package triangle

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Triangle struct {
	V1 Point
	V2 Point
	V3 Point
}

func (t Triangle) Perimeter() float64 {
	return distance(t.V1, t.V2) + distance(t.V2, t.V3) + distance(t.V3, t.V1)
}

func (t Triangle) Area() float64 {
	a := distance(t.V1, t.V2)
	b := distance(t.V2, t.V3)
	c := distance(t.V3, t.V1)
	s := (a + b + c) / 2.0
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func distance(p1 Point, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func PointInTriangle(pt Point, v1 Point, v2 Point, v3 Point) bool {
	b1 := (pt.X-v2.X)*(v1.Y-v2.Y)-(v1.X-v2.X)*(pt.Y-v2.Y) < 0.0
	b2 := (pt.X-v3.X)*(v2.Y-v3.Y)-(v2.X-v3.X)*(pt.Y-v3.Y) < 0.0
	b3 := (pt.X-v1.X)*(v3.Y-v1.Y)-(v3.X-v1.X)*(pt.Y-v1.Y) < 0.0
	return ((b1 == b2) && (b2 == b3))
}
