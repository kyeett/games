package grid

import (
	"github.com/peterhellberg/gfx"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func Test_Grid(t *testing.T) {
	g := New(10, 10, 10, 10, 0, 0)

	tcs := []struct {
		expected image.Point
		input    gfx.Vec
	}{
		{
			gfx.Pt(0, 0),
			gfx.V(0, 0),
		},
		{
			gfx.Pt(1, 0),
			gfx.V(10, 0),
		},
		{
			gfx.Pt(1, 1),
			gfx.V(10, 10),
		},
		{
			gfx.Pt(9, 9),
			gfx.V(99, 99),
		},
	}
	for _, tc := range tcs {
		point, found := g.ToPoint(tc.input)
		assert.True(t, found)
		assert.Equal(t, tc.expected, point)
	}
}

func Test_GridOutside(t *testing.T) {
	g := New(10, 10, 10, 10, 0, 0)

	tcs := []struct {
		input    gfx.Vec
	}{
		{
			gfx.V(-1, -1),
		},
		{
			gfx.V(-1, 0),
		},
		{
			gfx.V(0, -1),
		},
		{
			gfx.V(100, 100),
		},
	}
	for _, tc := range tcs {
		_, found := g.ToPoint(tc.input)
		assert.False(t, found)
	}
}

func Test_GridPadding(t *testing.T) {
	g := New(10, 10, 10, 10, 1, 1)

	tcs := []struct {
		expected image.Point
		input    gfx.Vec
	}{
		{
			gfx.Pt(0, 0),
			gfx.V(0, 0),
		},
		{
			gfx.Pt(0, 0),
			gfx.V(10, 0),
		},
		{
			gfx.Pt(1, 0),
			gfx.V(11, 0),
		},
		{
			gfx.Pt(9, 9),
			gfx.V(109, 109),
		},
	}
	for _, tc := range tcs {
		point, found := g.ToPoint(tc.input)
		assert.True(t, found)
		assert.Equal(t, tc.expected, point)
	}
}
