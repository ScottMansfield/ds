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

func (m *MaxHeap) printInOrder() {
	if len(m.data) == 0 {
		return
	}
	printInOrderRec(m.data[0], 0)
	fmt.Println()
}

func printInOrderRec(node *heapNode, level int) {
	if node == nil {
		return
	}

	printInOrderRec(node.left, level+1)
	fmt.Printf("%s(%d,%d) ", node.key, level, node.rank)
	printInOrderRec(node.right, level+1)
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

	parentIdx := (idx - 1) / 2
	node.parent = m.data[parentIdx]
	if idx&0x1 == 0 {
		node.parent.right = node
	} else {
		node.parent.left = node
	}

	m.heapifyUp(node)

	return nil
}

func (m *MaxHeap) ChangeRank(key string, amount int) error {
	if amount == 0 {
		return nil
	}

	node, ok := m.keyMap[key]
	if !ok {
		return ErrKeyDoesNotExist
	}

	node.rank += amount

	if amount > 0 {
		m.heapifyUp(node)
	} else {
		m.heapifyDown(node)
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
		key:  root.key,
		rank: root.rank,
		data: root.data,
	}

	// Remove key from the key map so tests for existence work properly
	delete(m.keyMap, root.key)

	// shortcut root removal
	if len(m.data) == 1 {
		m.data = make([]*heapNode, 0)
		return retval
	}

	// Swap "last" node into root
	idx := len(m.data) - 1
	lastNode := m.data[idx]
	m.swapNodes(root, lastNode)

	// shorten slice to exclude last node
	m.data = m.data[:len(m.data)-1]

	// unhook the parent
	if idx&0x1 == 0 {
		lastNode.parent.right = nil
	} else {
		lastNode.parent.left = nil
	}

	m.heapifyDown(root)

	return retval
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

func (m *MaxHeap) heapifyUp(node *heapNode) {
	// stop at the root
	if node.parent == nil {
		return
	}

	// stop when the parent node is no longer bigger
	if node.parent.rank < node.rank {
		m.swapNodes(node.parent, node)
		m.heapifyUp(node.parent)
	}
}

func (m *MaxHeap) heapifyDown(node *heapNode) {
	if node.left == nil && node.right == nil {
		return
	}

	if node.right == nil {
		if node.left.rank > node.rank {
			m.swapNodes(node.left, node)
		}
		return
	}

	if node.left.rank > node.right.rank {
		if node.left.rank > node.rank {
			m.swapNodes(node.left, node)
			m.heapifyDown(node.left)
		}
	} else {
		if node.right.rank > node.rank {
			m.swapNodes(node.right, node)
			m.heapifyDown(node.right)
		}
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
