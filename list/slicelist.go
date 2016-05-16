package list

import "github.com/ScottMansfield/ds"

type SliceList struct {
	l []ds.Comparable
}

func NewSliceList() List {
	return &SliceList{}
}

func NewSliceListWithCap(cap int) List {
	return &SliceList{
		l: make([]ds.Comparable, 0, cap),
	}
}

func (s *SliceList) Add(item ds.Comparable) {
	s.l = append(s.l, item)
}

func (s *SliceList) AddAll(items ...ds.Comparable) {
	s.l = append(s.l, items...)
}

func (s *SliceList) Contains(item ds.Comparable) bool {
	for _, it := range s.l {
		if it.Compare(item) == ds.Eq {
			return true
		}
	}
	return false
}

func (s *SliceList) ContainsAll(items ...ds.Comparable) bool {
	if len(items) > len(s.l) {
		return false
	}

	// TODO: can this be optimized further with a map?
	// It might be optimizable with a map if we weren't using a comparison func
	// in the type. It might be possible to do a first pass on value comparison
	// but then it's still slow if one of the item isn't found on the first pass.
	// Any item not found becomes a full list scan with the cmp func. This is
	// really not any better, so I'm leaving the naive approach for now.
	for _, item := range items {
		for _, it := range s.l {
			if item.Compare(it) != ds.Eq {
				return false
			}
		}
	}
	return true
}

func (s *SliceList) Remove(item ds.Comparable) {
	idx := -1
	for i, it := range s.l {
		if it.Compare(item) == ds.Eq {
			idx = i
		}
	}

	if idx < 0 {
		return
	}

	if idx == len(s.l)-1 {
		s.l = s.l[:idx]
	} else {
		s.l = append(s.l[:idx], s.l[idx+1:]...)
	}
}

func (s *SliceList) RemoveAll(items ...ds.Comparable) {
	var new []ds.Comparable

	for _, it := range s.l {
		keep := true
		for _, rem := range items {
			if it.Compare(rem) == ds.Eq {
				keep = false
				break
			}
		}
		if keep {
			new = append(new, it)
		}
	}

	s.l = new
}

func (s *SliceList) Clear() {
	s.l = make([]ds.Comparable, 0)
}

func (s *SliceList) Size() int {
	return len(s.l)
}

func (s *SliceList) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceList) ToSlice() []ds.Comparable {
	return append([]ds.Comparable(nil), s.l...)
}
