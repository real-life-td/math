package raycast

import (
	"github.com/real-life-td/math/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

const delta = 0.000000001

func TestClosestPointTo(t *testing.T) {
	test := func(lp1x, lp1y, lp2x, lp2y, px, py, ecx, ecy int, ed float64) {
		linePoint1 := primitives.NewPoint(lp1x, lp1y)
		linePoint2 := primitives.NewPoint(lp2x, lp2y)
		p := primitives.NewPoint(px, py)

		closest, distance := ClosestPointTo(linePoint1, linePoint2, p)
		require.Equal(t, primitives.NewPoint(ecx, ecy), closest)
		require.InDelta(t, ed, distance, delta)
	}

	test(0, 0, 100, 0, 50, 50, 50, 0, 50.0)
	test(0, 0, 100, 0, 0, 50, 0, 0, 50.0)
	test(0, 0, 100, 0, -50, 50, 0, 0, 70.71067811865476)
	test(0, 0, 100, 0, -50, 0, 0, 0, 50.0)
	test(0, 0, 100, 0, -50, -50, 0, 0, 70.71067811865476)
	test(0, 0, 100, 0, 0, -50, 0, 0, 50.0)
	test(0, 0, 100, 0, 50, -50, 50, 0, 50.0)
	test(0, 0, 100, 0, 100, -50, 100, 0, 50.0)
	test(0, 0, 100, 0, 150, -50, 100, 0, 70.71067811865476)
	test(0, 0, 100, 0, 150, 0, 100, 0, 50.0)
	test(0, 0, 100, 0, 150, 50, 100, 0, 70.71067811865476)
	test(0, 0, 100, 0, 100, 50, 100, 0, 50.0)
}