package set

type SortedSet struct {
    items []interface{}
    cmp   func(a, b interface{})
}

func NewSortedSet(cmp func(a, b interface{})) *SortedSet {
    items := make([]interface{}, 0)

    return &SortedSet{
        items: items,
        cmp:   cmp,
    }
}

func (s *SortedSet) Add(item interface{}) error {
    // lol
    return nil



foo := []int {1,2,3,4}

fmt.Printf("%v", foo)

end := append([]int{val}, foo[1:]...)

fmt.Printf("%v", foo)

foo = append(foo[:1], end...)

fmt.Printf("%v", foo)
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

