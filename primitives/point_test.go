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
