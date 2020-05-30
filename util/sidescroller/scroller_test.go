package sidescroller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScrollFunctionX(t *testing.T) {
	xMin, xMax := 0.0, 100.0
	tcs := []struct {
		x        float64
		expected float64
	}{
		{-10, -1},
		{0, -1},
		{5, -0.5},
		{10, 0},
		{90, 0},
		{95, 0.5},
		{100, 1},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, scrollFunction(tc.x, xMin, xMax, 10), "Wrong value at %v", tc.x)
	}
}

func TestScrollFunctionXOffset(t *testing.T) {
	xMin, xMax := 100.0, 200.0
	tcs := []struct {
		x        float64
		expected float64
	}{
		{90, -1},
		{100, -1},
		{105, -0.5},
		{110, 0},
		{190, 0},
		{195, 0.5},
		{200, 1},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, scrollFunction(tc.x, xMin, xMax, 10), "Wrong value at %v", tc.x)
	}
}
