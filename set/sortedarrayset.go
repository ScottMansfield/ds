package set

type SortedArraySet struct {
	items []interface{}
	cmp   func(a, b interface{}) int
}

func NewSortedArraySet(cmp func(a, b interface{}) int) *SortedArraySet {
	items := make([]interface{}, 0)

	return &SortedArraySet{
		items: items,
		cmp:   cmp,
	}
}

func (s *SortedArraySet) Add(item interface{}) error {
	// lol
	return nil
}

func (s *SortedArraySet) AddAll(...interface{}) error {
	// lol
	return nil
}

func (s *SortedArraySet) Contains(interface{}) (bool, error) {
	// lol
	return false, nil
}

func (s *SortedArraySet) ContainsAll(...interface{}) (bool, error) {
	// lol
	return false, nil
}

func (s *SortedArraySet) Remove(interface{}) error {
	// lol
	return nil
}

func (s *SortedArraySet) RemoveAll(...interface{}) error {
	// lol
	return nil
}

func (s *SortedArraySet) Clear() error {
	// lol
	return nil
}

func (s *SortedArraySet) Size() (int, error) {
	// lol
	return -1, nil
}

func (s *SortedArraySet) IsEmpty() (bool, error) {
	// lol
	return false, nil
}

func (s *SortedArraySet) ToSlice() ([]interface{}, error) {
	// lol
	return nil, nil
}
