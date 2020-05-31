package move

import "github.com/peterhellberg/gfx"

// Towards moves from `from` towards `target`
func Towards(from, target gfx.Vec, distance float64) (posAfter gfx.Vec, remainder float64, targetReached bool) {
	targetDistance := from.To(target).Len()

	switch  {
	case targetDistance <= distance:
		targetReached = true
		posAfter = target
		remainder = distance - targetDistance
	default:
		targetReached = false
		move := from.To(target).Unit().Scaled(distance)
		posAfter = from.Add(move)
		remainder = 0
	}

	return posAfter, remainder, targetReached
}