package primitives

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(0, 1)
	require.Equal(t, 0, p.X())
	require.Equal(t, 1, p.Y())
}

func TestPoint_InRectangle(t *testing.T) {
	p := NewPoint(0, 0)
	require.True(t, p.InRectangle(NewRectangle(-1, -1, 1, 1)))
	require.True(t, p.InRectangle(NewRectangle(-1, 0, 1, 0)), "zero height")
	require.True(t, p.InRectangle(NewRectangle(0, -1, 0, 1)), "zero width")
	require.False(t, p.InRectangle(NewRectangle(-1, -2, 1, -1)), "too high")
	require.False(t, p.InRectangle(NewRectangle(-1, 1, 1, 2)), "too low")
	require.False(t, p.InRectangle(NewRectangle(-2, -1, -1, 1)), "too leftwards")
	require.False(t, p.InRectangle(NewRectangle(1, -1, 2, -1)), "too rightwards")
}

func TestPoint_Dot(t *testing.T) {
	p1 := NewPoint(10, 5)
	p2 := NewPoint(2, 3)

	require.Equal(t, 125, p1.Dot(p1))
	require.Equal(t, 13, p2.Dot(p2))
	require.Equal(t, 35, p1.Dot(p2))
	require.Equal(t, 35, p2.Dot(p1))
}

func TestPoint_Cross(t *testing.T) {
	p1 := NewPoint(10, 5)
	p2 := NewPoint(2, 3)

	require.Equal(t, 100, p1.Cross(p1))
	require.Equal(t, 12, p2.Cross(p2))
	require.Equal(t, 40, p1.Cross(p2))
	require.Equal(t, 40, p2.Cross(p1))
}

func TestPoint_Add(t *testing.T) {
	p1 := NewPoint(10, 5)
	p2 := NewPoint(2, 3)

	require.Equal(t, NewPoint(20, 10), p1.Add(p1))
	require.Equal(t, NewPoint(4, 6), p2.Add(p2))
	require.Equal(t, NewPoint(12, 8), p1.Add(p2))
	require.Equal(t, NewPoint(12, 8), p2.Add(p1))
}

func TestPoint_Subtract(t *testing.T) {
	p1 := NewPoint(10, 5)
	p2 := NewPoint(2, 3)

	require.Equal(t, NewPoint(0, 0), p1.Subtract(p1))
	require.Equal(t, NewPoint(0, 0), p2.Subtract(p2))
	require.Equal(t, NewPoint(8, 2), p1.Subtract(p2))
	require.Equal(t, NewPoint(-8, -2), p2.Subtract(p1))
}

func TestPoint_Multiply(t *testing.T) {
	p := NewPoint(2, 3)

	require.Equal(t, NewPoint(-2, -3), p.Multiply(-1))
	require.Equal(t, NewPoint(0, 0), p.Multiply(0))
	require.Equal(t, NewPoint(2, 3), p.Multiply(1))
	require.Equal(t, NewPoint(4, 6), p.Multiply(2))
}

func TestPoint_MultiplyF(t *testing.T) {
	p := NewPoint(2, 3)

	require.Equal(t, NewPoint(-2, -3), p.MultiplyF(-1.0))
	require.Equal(t, NewPoint(-1, -2), p.MultiplyF(-(2.0 / 3.0)))
	require.Equal(t, NewPoint(-1, -1), p.MultiplyF(-(1.0 / 3.0)))
	require.Equal(t, NewPoint(0, 0), p.MultiplyF(0.0))
	require.Equal(t, NewPoint(1, 1), p.MultiplyF(1.0/3.0))
	require.Equal(t, NewPoint(1, 2), p.MultiplyF(2.0/3.0))
	require.Equal(t, NewPoint(2, 3), p.MultiplyF(1.0))
	require.Equal(t, NewPoint(4, 6), p.MultiplyF(2.0))
}
