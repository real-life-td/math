package raycast

import (
	"github.com/real-life-td/math/primitives"
	"math"
)

// Algorithm from https://math.stackexchange.com/a/2193733
func ClosestPointTo(linePoint1, linePoint2, p *primitives.Point) (closest *primitives.Point, distance float64) {
	v := linePoint2.Subtract(linePoint1)
	u := linePoint1.Subtract(p)

	vDotV := v.Dot(v)
	vDotU := v.Dot(u)

	t := - (float64(vDotU) / float64(vDotV))

	if 0.0 <= t && t <= 1.0 {
		// Closest point lies somewhere along the line
		closest = linePoint1.MultiplyF(1 - t).Add(linePoint2.MultiplyF(t))
		delta := closest.Subtract(p)
		return closest, math.Sqrt(float64(delta.Dot(delta)))
	} else {
		// Closest point lies on one of the endpoints
		g0 := u.Dot(u)
		g1 := vDotV + 2 * vDotU + g0

		if g0 < g1 {
			return linePoint1, math.Sqrt(float64(g0))
		} else {
			return linePoint2, math.Sqrt(float64(g1))
		}
	}
}