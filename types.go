package ds

type Sizer interface {
	Size()
}

type Res int

const (
	Gt Res = iota
	Lt
	Eq
)

type Comparable interface {
	Compare(a, b interface{}) Res
}
