package list

type List interface {
	Add(interface{})
	AddAll(...interface{})
	Contains(interface{}) bool
	ContainsAll(...interface{}) bool
	Remove(interface{})
	RemoveAll(...interface{})
	Clear()
	Size() int
	IsEmpty() bool
	ToSlice() []interface{}
}
