package primitives

type Rectangle struct {
	x1, y1, x2, y2 int
}

func NewRectangle(x1, y1, x2, y2 int) *Rectangle {
	r := new(Rectangle)

	if x1 > x2 {
		r.x1 = x2
		r.x2 = x1
	} else {
		r.x1 = x1
		r.x2 = x2
	}

	if y1 > y2 {
		r.y1 = y2
		r.y2 = y1
	} else {
		r.y1 = y1
		r.y2 = y2
	}

	return r
}

func RectangleFromPoints(p1, p2 *Point) *Rectangle {
	return NewRectangle(p1.x, p1.y, p2.x, p2.y)
}

// The least x-value of the rectangle
func (r *Rectangle) X1() int {
	return r.x1
}

// The least y-value of the rectangle
func (r *Rectangle) Y1() int {
	return r.y1
}

// The greatest x-value of the rectangle
func (r *Rectangle) X2() int {
	return r.x2
}

// The greatest y-value of the rectangle
func (r *Rectangle) Y2() int {
	return r.y2
}

func (r *Rectangle) Width() int {
	return r.x2 - r.x1
}

func (r *Rectangle) Height() int {
	return r.y2 - r.y1
}

func (r *Rectangle) ContainsPoint(p *Point) bool {
	return r.x1 <= p.x && p.x <= r.x2 && r.y1 <= p.y && p.y <= r.y2
}
