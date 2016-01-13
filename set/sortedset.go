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
