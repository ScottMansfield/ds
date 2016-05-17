package tree

type Tree interface {
	Add(interface{}) error
	AddAll(...interface{}) error
	Contains(interface{}) (bool, error)
	ContainsAll(...interface{}) (bool, error)
	Remove(interface{}) error
	RemoveAll(...interface{}) error
	Clear() error
	Size() (int, error)
	IsEmpty() (bool, error)
	ToSlice() ([]interface{}, error)
}
