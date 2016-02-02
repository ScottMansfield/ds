package set

type color bool

var (
	red   = color(true)
	black = color(false)
)

type node struct {
	color  color
	parent *node
	left   *node
	right  *node
	data   interface{}
}

var leaf = &node{
	color:  black,
	parent: nil,
	left:   nil,
	right:  nil,
	data:   struct{}{},
}

type SortedTreeSet struct {
	root *node
	cmp  func(a, b interface{}) int
}

func NewSortedTreeSet(cmp func(a, b interface{}) int) *SortedTreeSet {
	return &SortedTreeSet{
		root: nil,
		cmp:  cmp,
	}
}

func (s *SortedTreeSet) Add(item interface{}) error {
	if s.root == nil {
		s.root = &node{
			color:  black,
			parent: nil,
			left:   leaf,
			right:  leaf,
			data:   item,
		}
	} else {
		insert(s.root, item, s.cmp)
	}

	return nil
}

func insert(root *node, item interface{}, cmp func(a, b interface{}) int) error {
	//foo
	return nil
}

func (s *SortedTreeSet) AddAll(...interface{}) error {
	// lol
	return nil
}

func (s *SortedTreeSet) Contains(interface{}) (bool, error) {
	// lol
	return false, nil
}

func (s *SortedTreeSet) ContainsAll(...interface{}) (bool, error) {
	// lol
	return false, nil
}

func (s *SortedTreeSet) Remove(interface{}) error {
	// lol
	return nil
}

func (s *SortedTreeSet) RemoveAll(...interface{}) error {
	// lol
	return nil
}

func (s *SortedTreeSet) Clear() error {
	// lol
	return nil
}

func (s *SortedTreeSet) Size() (int, error) {
	// lol
	return -1, nil
}

func (s *SortedTreeSet) IsEmpty() (bool, error) {
	// lol
	return false, nil
}

func (s *SortedTreeSet) ToSlice() ([]interface{}, error) {
	// lol
	return nil, nil
}
