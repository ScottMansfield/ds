package list

import "github.com/ScottMansfield/ds"

type List interface {
	Add(ds.Comparable)
	AddAll(...ds.Comparable)
	Contains(ds.Comparable) bool
	ContainsAll(...ds.Comparable) bool
	Remove(ds.Comparable)
	RemoveAll(...ds.Comparable)
	Clear()
	Size() int
	IsEmpty() bool
	ToSlice() []ds.Comparable
}
