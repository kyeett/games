package move

import (
	"fmt"
	"github.com/peterhellberg/gfx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Move(t *testing.T) {
	from := gfx.V(0, 0)
	target := gfx.V(10, 0)

	tcs := []struct {
		distance float64

		expectedAfter         gfx.Vec
		expectedRemainder     float64
		expectedTargetReached bool

	}{
		{
			5.0,

			gfx.V(5.0, 0),
			0.0,
			false,
		},
		{
			10.0,

			gfx.V(10.0, 0),
			0.0,
			true,
		},
		{
			20.0,

			gfx.V(10.0, 0),
			10.0,
			true,
		},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("distance:%v", tc.distance), func(t *testing.T) {
			posAfter, movementAfter, targetReached := Towards(from, target, tc.distance)
			assert.Equal(t, tc.expectedAfter, posAfter)
			assert.Equal(t, tc.expectedRemainder, movementAfter)
			assert.Equal(t, tc.expectedTargetReached, targetReached)
		})
	}
}
