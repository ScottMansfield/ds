package heap

import "errors"

var ErrKeyExists = errors.New("Key already exists")

type heapNode struct {
	parent *heapNode
	left   *heapNode
	right  *heapNode
	rank   int
	key    string
	data   interface{}
}

type Val struct {
	rank int
	key  string
	data interface{}
}

type MaxHeap struct {
	data   []*heapNode
	keyMap map[string]*heapNode
	size   int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:   make([]*heapNode, 0),
		keyMap: make(map[string]*heapNode),
		size:   0,
	}
}

// Adds a new node into the heap. If the key already exists, returns ErrKeyExists
func (m *MaxHeap) Add(key string, rank int, data interface{}) error {
	if _, ok := m.keyMap[key]; ok {
		return ErrKeyExists
	}

	m.size++

	node := &heapNode{
		key:  key,
		rank: rank,
		data: data,
	}

	idx := len(m.data)
	m.data = append(m.data, node)

	// If this is the first node, return
	if idx == 0 {
		return nil
	}

	node.parent = m.data[idx/2]
	if idx&0x1 == 1 {
		node.parent.right = node
	} else {
		node.parent.left = node
	}

	m.heapifyUp(node)

	return nil
}

func (m *MaxHeap) ChangeRank(key string, amount int) {
	// lol
}

func (m *MaxHeap) ExtractMax() heapNode {
	// lol
	return heapNode{}
}

func (m *MaxHeap) Size() int {
	return m.size
}

func (m *MaxHeap) PeekMax() *Val {
    if len(m.data) == 0 {
        return nil
    }
	root := m.data[0]
	return &Val{
		key:  root.key,
		rank: root.rank,
		data: root.data,
	}
}

func (m *MaxHeap) TopN(n int) []heapNode {
	// Make a copy to not disturb the heap
	temp := make([]*heapNode, len(m.data), len(m.data))
	for i, data := range m.data {
		temp[i] = data
	}
	tempHeap := &MaxHeap{
		data: temp,
		size: len(temp),
	}

	if tempHeap.Size() > n {
		n = tempHeap.Size()
	}

	ret := make([]heapNode, n, n)
	for i := 0; i < n; i++ {
		ret[i] = tempHeap.ExtractMax()
	}

	return nil
}

func (m *MaxHeap) heapifyUp(node *heapNode) {
	if node.parent == nil {
		return
	}

	if node.parent.rank < node.rank {
		m.swapNodes(node, node.parent)
		m.heapifyUp(node.parent)
	}
}

// Swap the data inside to avoid having to reach out and touch many other nodes
// to update left / right / parent pointers
func (m *MaxHeap) swapNodes(a, b *heapNode) {
	tempKey := a.key
	tempRank := a.rank
	tempData := a.data

	a.key = b.key
	a.rank = b.rank
	a.data = b.data

	b.key = tempKey
	b.rank = tempRank
	b.data = tempData

	m.keyMap[a.key] = a
	m.keyMap[b.key] = b
}
