package cmp

type Res int

const (
	Gt Res = iota
	Lt
	Eq
)

type Cmp interface {
	Compare(a, b interface{}) Res
}
