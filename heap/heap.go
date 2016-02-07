package heap

import (
	"errors"
	"fmt"
)

var (
	ErrKeyExists       = errors.New("Key already exists")
	ErrKeyDoesNotExist = errors.New("Key does not exist")
)

type heapNode struct {
	rank int
	key  string
	data interface{}
}

func (h *heapNode) clone() *heapNode {
	return &heapNode{
		rank: h.rank,
		key:  h.key,
		data: h.data,
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return idx*2 + 1
}

func right(idx int) int {
	return idx*2 + 2
}

type Val struct {
	Rank int
	Key  string
	Data interface{}
}

type MaxHeap struct {
	data   []*heapNode
	keyMap map[string]int
	size   int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:   make([]*heapNode, 0),
		keyMap: make(map[string]int),
		size:   0,
	}
}

func (m *MaxHeap) Clone() *MaxHeap {
	data := make([]*heapNode, len(m.data))
	keyMap := make(map[string]int)

	for i := 0; i < len(m.data); i++ {
		temp := m.data[i].clone()
		data[i] = temp

		keyMap[temp.key] = i
	}

	return &MaxHeap{
		data:   data,
		keyMap: keyMap,
		size:   len(data),
	}
}

func (m *MaxHeap) printInOrder() {
	if len(m.data) == 0 {
		return
	}
	printInOrderRec(0, 0, m.data)
	fmt.Println()
}

func printInOrderRec(idx, level int, data []*heapNode) {
	if idx >= len(data) {
		return
	}

	node := data[idx]
	left := left(idx)
	right := right(idx)

	printInOrderRec(left, level+1, data)
	fmt.Printf("%s(l:%d,r:%d) ", node.key, level, node.rank)
	printInOrderRec(right, level+1, data)
}

func (m *MaxHeap) PrintLevels() {
	if m.size == 0 {
		return
	}

	q := make(chan int, m.size/2+1)
	levelSize := 1
	numInLevel := 0

	q <- 0

	for {
		select {
		case idx := <-q:
			numInLevel++

			node := m.data[idx]
			left := left(idx)
			right := right(idx)

			fmt.Printf("%s(r:%d) ", node.key, node.rank)

			if left < len(m.data) {
				q <- left
			}
			if right < len(m.data) {
				q <- right
			}

			if numInLevel >= levelSize {
				fmt.Println()
				levelSize <<= 1
				numInLevel = 0
			}
		default:
			fmt.Println()
			return
		}
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
	m.keyMap[key] = idx
	m.data = append(m.data, node)

	// If this is the first node, return
	if idx == 0 {
		return nil
	}

	m.heapifyUp(idx)

	return nil
}

func (m *MaxHeap) ChangeRank(key string, amount int) error {
	if amount == 0 {
		return nil
	}

	idx, ok := m.keyMap[key]
	if !ok {
		return ErrKeyDoesNotExist
	}

	node := m.data[idx]
	node.rank += amount

	if amount > 0 {
		m.heapifyUp(idx)
	} else {
		m.heapifyDown(idx)
	}

	return nil
}

func (m *MaxHeap) ExtractMax() *Val {
	if len(m.data) == 0 {
		return nil
	}

	// take care of this upfront to avoid bugs on different return paths
	m.size--

	// Copy data for return value
	root := m.data[0]
	retval := &Val{
		Key:  root.key,
		Rank: root.rank,
		Data: root.data,
	}

	// Remove key from the key map so tests for existence work properly
	delete(m.keyMap, root.key)

	// shortcut root removal
	if len(m.data) == 1 {
		m.data = make([]*heapNode, 0)
		return retval
	}

	// Swap "last" node into root
	m.swapNodes(0, len(m.data)-1)

	// shorten slice to exclude last node
	m.data = m.data[:len(m.data)-1]

	m.heapifyDown(0)

	return retval
}

func (m *MaxHeap) Size() int {
	return m.size
}

func (m *MaxHeap) Empty() bool {
	return m.size == 0
}

func (m *MaxHeap) PeekMax() *Val {
	if len(m.data) == 0 {
		return nil
	}
	root := m.data[0]
	return &Val{
		Key:  root.key,
		Rank: root.rank,
		Data: root.data,
	}
}

func (m *MaxHeap) heapifyUp(idx int) {
	// stop at the root
	if idx == 0 {
		return
	}

	pidx := parent(idx)
	parent := m.data[pidx]
	node := m.data[idx]

	// stop when the parent node is no longer smaller
	if parent.rank < node.rank {
		m.swapNodes(idx, pidx)
		m.heapifyUp(pidx)
	}
}

func (m *MaxHeap) heapifyDown(idx int) {
	// If there's no children, break
	lidx := left(idx)
	if lidx >= len(m.data) {
		return
	}

	node := m.data[idx]
	ridx := right(idx)
	left := m.data[lidx]

	// simpler case of one child. The left one.
	if ridx >= len(m.data) {
		if left.rank > node.rank {
			m.swapNodes(lidx, idx)
		}
		return
	}

	right := m.data[ridx]

	if left.rank > right.rank {
		if left.rank > node.rank {
			m.swapNodes(lidx, idx)
			m.heapifyDown(lidx)
		}
	} else {
		if right.rank > node.rank {
			m.swapNodes(ridx, idx)
			m.heapifyDown(ridx)
		}
	}
}

// Swap two nodes while ensuring the map stays up to date
func (m *MaxHeap) swapNodes(a, b int) {
	temp := m.data[a]
	m.data[a] = m.data[b]
	m.data[b] = temp

	m.keyMap[m.data[a].key] = a
	m.keyMap[m.data[b].key] = b
}
