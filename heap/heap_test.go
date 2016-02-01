package heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	m := NewMaxHeap()

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	m.Add("foo", 1, "data")

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	m.Add("bar", 2, "data")

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	m.Add("baz", 3, "data")

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())
}
