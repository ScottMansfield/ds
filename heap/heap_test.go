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

	if peeked.Rank != 3 {
		t.Errorf("Expected rank 3 but got %v", peeked.Rank)
	}

	m.printInOrder()
}

func TestMaxHeapManyItems(t *testing.T) {
	m := NewMaxHeap()

	numItems := 1000000

	for i := 0; i < numItems; i++ {
		m.Add(fmt.Sprint(i), i, "data")
	}

	for i := 0; i < numItems; i++ {
		val := m.ExtractMax()
		if val.Rank != numItems-1-i {
			t.Errorf("Expected rank %d but got %d", numItems-1-i, val.Rank)
			return
		}
	}
}

func TestMaxHeapSameRanks(t *testing.T) {
	m := NewMaxHeap()

	numItems := 10

	for i := 0; i < numItems; i++ {
		m.Add(fmt.Sprint(i), 5, fmt.Sprintf("data%d", i))
	}

	for i := 0; i < numItems; i++ {
		max := m.ExtractMax()
		if max == nil || max.Rank != 5 {
			t.Error("Max was nil or not of rank 5")
		}
	}
}

func TestMaxHeapChangeRank(t *testing.T) {
	m := NewMaxHeap()

	numItems := 10

	for i := 1; i <= numItems; i++ {
		m.Add(fmt.Sprint(i), i, "data")
	}

	max := m.PeekMax()

	if max.Rank != 10 {
		t.Errorf("Expected rank %d, got rank %d", numItems, max.Rank)
	}

	m.ChangeRank("1", 99)

	max = m.PeekMax()

	if max.Key != "1" {
		t.Errorf("After changing rank, got wrong key as max: %s", max.Key)
	}
	if max.Rank != 100 {
		t.Errorf("After changing rank, got rank %d instead of 100", max.Rank)
	}

	m.ChangeRank("1", -94)

	max = m.PeekMax()

	if max.Key != "10" {
		t.Errorf("After changing rank second time, got wrong key as max: %d", max.Key)
	}
	if max.Rank != 10 {
		t.Errorf("After changing rank, got rank %d instead of 10", max.Rank)
	}
}

func TestMaxHeapClone(t *testing.T) {
	m := NewMaxHeap()

	numItems := 10

	for i := 1; i <= numItems; i++ {
		m.Add(fmt.Sprint(i), i, "data")
	}

	clone := m.Clone()

	m.PrintLevels()

	m.ExtractMax()

	for i := numItems; i > 0; i-- {
		max := clone.ExtractMax()
		if max.Rank != i {
			t.Errorf("Expected rank %d but got %d", i, max.Rank)
		}
	}
}
