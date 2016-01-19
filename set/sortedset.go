package set

type SortedSet struct {
    items []interface{}
    cmp   func(a, b interface{})
}

func NewSortedSet(size int, cmp func(a, b interface{})) *SortedSet {
    if size < 0 {
        size = 16
    }

    items := make([]interface{}, size)

    return &SortedSet{
        items: items,
        cmp:   cmp,
    }
}

func (s *SortedSet) Add(item interface{}) error {
    // lol
    return nil
}

func (s *SortedSet) AddAll(...interface{}) error {
    // lol
    return nil
}

func (s *SortedSet) Contains(interface{}) (bool, error) {
    // lol
    return false, nil
}

func (s *SortedSet) ContainsAll(...interface{}) (bool, error) {
    // lol
    return false, nil
}

func (s *SortedSet) Remove(interface{}) error {
    // lol
    return nil
}

func (s *SortedSet) RemoveAll(...interface{}) error {
    // lol
    return nil
}

func (s *SortedSet) Clear() error {
    // lol
    return nil
}

func (s *SortedSet) Size() (int, error) {
    // lol
    return -1, nil
}

func (s *SortedSet) IsEmpty() (bool, error) {
    // lol
    return false, nil
}

func (s *SortedSet) ToSlice() ([]interface{}, error) {
    // lol
    return nil, nil
}

