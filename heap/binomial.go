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

func (b *BinomialMaxHeap) Merge(h BinomialMaxHeap) *BinomialMaxHeap {
	return nil
}
