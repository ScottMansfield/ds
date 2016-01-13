package set

type Set interface {
    Add(interface{}) error
}

func New() Set {
    return NewSyncMapSet()
}
