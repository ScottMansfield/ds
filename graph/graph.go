package graph

import (
	"fmt"
	"io"

	"github.com/ScottMansfield/ds/heap"
	"github.com/ScottMansfield/ds/set"
)

type DiGraph struct {
	in      map[string]set.Set
	out     map[string]set.Set
	inHeap  *heap.MaxHeap
	outHeap *heap.MaxHeap
}

func NewDiGraph() *DiGraph {
	return &DiGraph{
		in:      make(map[string]set.Set),
		out:     make(map[string]set.Set),
		inHeap:  heap.NewMaxHeap(),
		outHeap: heap.NewMaxHeap(),
	}
}

func (d *DiGraph) AddEdge(from, to string) {
	if _, ok := d.out[from]; !ok {
		d.out[from] = set.New()
		if err := d.outHeap.Add(from, 0, nil); err != nil {
			panic(err)
		}
	}
	d.out[from].Add(to)
	if err := d.outHeap.ChangeRank(from, 1); err != nil {
		panic(err)
	}

	if _, ok := d.in[to]; !ok {
		d.in[to] = set.New()
		if err := d.inHeap.Add(to, 0, nil); err != nil {
			panic(err)
		}
	}
	d.in[to].Add(from)
	if err := d.inHeap.ChangeRank(to, 1); err != nil {
		panic(err)
	}
}

func (d *DiGraph) WriteDOT(w io.Writer) {
	fmt.Fprintln(w, "digraph goimports {")
	for src, dsts := range d.out {
		dstSlice, _ := dsts.ToSlice()
		for _, dst := range dstSlice {
			fmt.Fprintf(w, "\t\"%s\" -> \"%s\";\n", src, dst)
		}
	}

	fmt.Fprintln(w, "}")
}

type RanksResult struct {
	Key  string
	Rank int
}

// Panics if num < 0
func (d *DiGraph) TopNInRanks(num int) []RanksResult {
	if num < 0 {
		panic("Number of requested nodes must be 0 or greater")
	}

	heap := d.inHeap.Clone()
	return topN(heap, num)
}

// Panics if num < 0
func (d *DiGraph) TopNOutRanks(num int) []RanksResult {
	if num < 0 {
		panic("Number of requested nodes must be 0 or greater")
	}

	heap := d.outHeap.Clone()
	return topN(heap, num)
}

func topN(heap *heap.MaxHeap, num int) []RanksResult {
	limit := heap.Size()
	if num < limit {
		limit = num
	}

	retval := make([]RanksResult, limit)
	for i := 0; i < limit; i++ {
		max := heap.ExtractMax()
		retval[i] = RanksResult{
			Key:  max.Key,
			Rank: max.Rank,
		}
	}

	return retval
}
