package pathfinding

import (
	"github.com/beefsack/go-astar"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	w := New(2, 1)

	require.NotNil(t, w.Get(0, 0))
	require.NotNil(t, w.Get(1, 0))

	require.Nil(t, w.Get(-1, -1))
	require.Nil(t, w.Get(2, 0))
	require.Nil(t, w.Get(0, 1))
}

func TestStraight(t *testing.T) {
	w := New(4, 4)
	t1 := w.Get(0, 0)
	t2 := w.Get(3, 0)

	// Act
	path, distance, found := astar.Path(t1, t2)

	// Assert
	require.Len(t, path, 4)
	require.Equal(t, distance, 3.0)
	require.True(t, found)
}

func TestAcross(t *testing.T) {
	w := New(4, 4)
	t1 := w.Get(0, 0)
	t2 := w.Get(3, 3)

	// Act
	path, distance, found := astar.Path(t1, t2)

	// Assert
	require.Len(t, path, 7)
	require.Equal(t, distance, 6.0)
	require.True(t, found)
}

func TestDown(t *testing.T) {
	w := New(4, 4)
	t1 := w.Get(0, 0)
	t2 := w.Get(0, 3)

	// Act
	path, distance, found := astar.Path(t1, t2)

	// Assert
	require.Len(t, path, 4)
	require.Equal(t, distance, 3.0)
	require.True(t, found)
}

func TestPathBlocked(t *testing.T) {
	// Layout:
	//
	//   AXB
	//   ---

	w := New(3, 2)
	t1 := w.Get(0, 0)
	t2 := w.Get(2, 0)

	w.Get(1, 0).Kind = KindBlocker

	// Act
	path, distance, found := astar.Path(t1, t2)

	// Assert
	require.Len(t, path, 5)
	require.Equal(t, distance, 4.0)
	require.True(t, found)
}

func TestNoPath(t *testing.T) {
	// Layout:
	//
	//   AXB
	//   -X-

	w := New(3, 2)
	t1 := w.Get(0, 0)
	t2 := w.Get(2, 0)

	w.Get(1, 0).Kind = KindBlocker
	w.Get(1, 1).Kind = KindBlocker

	// Act
	_, _, found := astar.Path(t1, t2)

	// Assert
	require.False(t, found)
}