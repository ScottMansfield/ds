package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	m := NewMaxHeap()

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	t.Logf("%+v", m)

	m.Add("foo", 1, "data")
	m.printInOrder()
	m.Add("bar", 2, "data")
	m.printInOrder()
	m.Add("bip", 4, "data")
	m.printInOrder()
	m.Add("baz", 3, "data")
	m.printInOrder()

	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	max := m.ExtractMax()
	t.Logf("Extracted max: %+v", max)
	t.Logf("Size: %v", m.Size())
	t.Logf("Max element: %+v", m.PeekMax())

	peeked := m.PeekMax()

	if peeked.rank != 3 {
		t.Errorf("Expected rank 3 but got %v", peeked.rank)
	}

	m.printInOrder()
}

func TestMaxHeapManyItems(t *testing.T) {
	m := NewMaxHeap()

	numItems := 1000000

	for i := 0; i < numItems; i++ {
		m.Add(fmt.Sprint(i), i, "data")
	}

	//m.printInOrder()

	for i := 0; i < numItems; i++ {
		m.ExtractMax()
	}

	//m.printInOrder()
}

func TestMaxHeapSameRanks(t *testing.T) {
	m := NewMaxHeap()

	numItems := 10

	for i := 0; i < numItems; i++ {
		m.Add(fmt.Sprint(i), 5, fmt.Sprintf("data%d", i))
	}

	for m.Size() > 0 {
		fmt.Printf("%+v\n", m.ExtractMax())
	}
}
