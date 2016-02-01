package graph

//import "github.com/ScottMansfield/ds/heap"
import "../heap"
import "github.com/ScottMansfield/ds/set"

type DiGraph struct {
    in  map[string]set.Set
    out map[string]set.Set
}

func NewDiGraph() *DiGraph {
    return &DiGraph{
        in:  make(map[string]set.Set),
        out: make(map[string]set.Set),
    }
}

func (d *DiGraph) AddEdge(from, to string) {
    if _, ok := d.out[from]; !ok {
        d.out[from] = set.New()
    }
    d.out[from].Add(to)

    if _, ok := d.in[to]; !ok {
        d.in[to] = set.New()
    }
    d.in[to].Add(from)
}

type RanksResult struct {
    key string
    rank int
}

func (d *DiGraph) TopNInRanks(num int) []RanksResult {
    if num < 0 {
        num = 1
    }

    // use heap to sort
}
func (d *DiGraph) TopNOutRanks(num int) []RanksResult {
    if num < 0 {
        num = 1
    }

    // use heap to sort
}
