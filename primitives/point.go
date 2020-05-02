package primitives

import "math"

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) InRectangle(r *Rectangle) bool {
	return r.ContainsPoint(p)
}

func (p *Point) Dot(p2 *Point) int {
	return p.x*p2.x + p.y*p2.y
}

func (p *Point) Cross(p2 *Point) int {
	return p.x*p2.y + p.y*p2.x
}

func (p *Point) Add(p2 *Point) *Point {
	return NewPoint(p.x+p2.x, p.y+p2.y)
}

func (p *Point) Subtract(p2 *Point) *Point {
	return NewPoint(p.x-p2.x, p.y-p2.y)
}

func (p *Point) Multiply(scalar int) *Point {
	return NewPoint(p.x*scalar, p.y*scalar)
}

func (p *Point) MultiplyF(scalar float64) *Point {
	fx := float64(p.x) * scalar
	fy := float64(p.y) * scalar
	return NewPoint(int(math.Round(fx)), int(math.Round(fy)))
}
