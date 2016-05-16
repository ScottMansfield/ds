package heap

type binomialNode struct {
	next     *binomialNode
	children *binomialNode
	order    int
	rank     int
	key      string
	data     interface{}
}

type BinomialMaxHeap struct {
	size   int
	keyMap map[string]binomialNode
	root   *binomialNode
}

func NewBinomialMaxHeap() *BinomialMaxHeap {
	return &BinomialMaxHeap{}
}

func (b *BinomialMaxHeap) Add(key string, rank int, data interface{}) {
	newNode := &binomialNode{
		key:  key,
		rank: rank,
		data: data,
	}

	if b.root == nil {
		b.root = newNode
		return
	}

	// Adding a node is the same as making a new one-item heap and merging

}

func (b *BinomialMaxHeap) ChangeRank(key string, amount int) {

}

func (b *BinomialMaxHeap) Delete(key string) {

}

func (b *BinomialMaxHeap) ExtractMax() *Val {
	return nil
}

func (b *BinomialMaxHeap) PeekMax() *Val {
	return nil
}

// Merges BinomialHeap h into this heap,
func (b *BinomialMaxHeap) Merge(h BinomialMaxHeap) {
	// do nothing
}

func (bmh *BinomialMaxHeap) merge(b *binomialNode) {
	a := bmh.root

	// make a fake list head to make coding simpler
	newList := &binomialNode{}
	var extra *binomialNode

	for a != nil && b != nil {
		if extra != nil {
			if extra.order > a.order && extra.order > b.order {
				newList.next = extra
				extra = nil
			} else if extra.order == a.order {
				if extra.order > b.order {

				}
			}
		}
		if a.rank < b.rank {
			newList.next = a
			a = a.next
		} else if b.rank < a.rank {
			newList.next = b
			b = b.next
		} else {
			// Merge
			// make smaller one the end of the child linked list
		}
	}

	// chop off the fake head
	newList = newList.next
}

//func mergeTrees
