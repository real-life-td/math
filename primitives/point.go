package primitives

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
