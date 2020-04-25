package primitives

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	r := NewRectangle(0, 1, 2, 3)
	require.Equal(t, 0, r.X1())
	require.Equal(t, 1, r.Y1())
	require.Equal(t, 2, r.X2())
	require.Equal(t, 3, r.Y2())

	r = NewRectangle(2, 3, 0, 1)
	require.Equal(t, 0, r.X1(), "should auto-sort least a greatest coordinates")
	require.Equal(t, 1, r.Y1(), "should auto-sort least a greatest coordinates")
	require.Equal(t, 2, r.X2(), "should auto-sort least a greatest coordinates")
	require.Equal(t, 3, r.Y2(), "should auto-sort least a greatest coordinates")
}

func TestRectangle_Width(t *testing.T) {
	r := NewRectangle(0, 0, 3, 4)
	require.Equal(t, 3, r.Width())
}

func TestRectangle_Height(t *testing.T) {
	r := NewRectangle(0, 0, 3, 4)
	require.Equal(t, 4, r.Height())
}

func TestRectangle_ContainsPoint(t *testing.T) {
	r := NewRectangle(-1, -1, 1, 1)
	require.True(t, r.ContainsPoint(NewPoint(0, 0)))
	require.True(t, r.ContainsPoint(NewPoint(-1, 0)), "on edge")
	require.True(t, r.ContainsPoint(NewPoint(1, 0)), "on edge")
	require.True(t, r.ContainsPoint(NewPoint(0, -1)), "on edge")
	require.True(t, r.ContainsPoint(NewPoint(0, 1)), "on edge")
	require.True(t, r.ContainsPoint(NewPoint(-1, -1)), "on corner")
	require.True(t, r.ContainsPoint(NewPoint(1, -1)), "on corner")
	require.True(t, r.ContainsPoint(NewPoint(1, 1)), "on corner")
	require.True(t, r.ContainsPoint(NewPoint(-1, 1)), "on corner")
	require.False(t, r.ContainsPoint(NewPoint(0, -2)), "Too Low")
	require.False(t, r.ContainsPoint(NewPoint(0, 2)), "Too High")
	require.False(t, r.ContainsPoint(NewPoint(-2, 0)), "Too Leftwards")
	require.False(t, r.ContainsPoint(NewPoint(2, 0)), "Too Rightwards")
}
