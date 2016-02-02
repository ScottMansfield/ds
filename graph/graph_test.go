package graph

import (
	"fmt"
	"testing"
)

func TestGraphTopRanks(t *testing.T) {
	g := NewDiGraph()

	numNodes := 100

	for i := 0; i < numNodes; i++ {
		for j := i; j < numNodes; j++ {
			g.AddEdge(fmt.Sprintf("%d", i), fmt.Sprintf("%d", j))
		}
	}

	t.Log("Top 10 in ranks:")
	for i, res := range g.TopNInRanks(10) {
		if res.Rank != numNodes-i {
			t.Errorf("Expected rank %d but got %d", numNodes-i, res.Rank)
		}
		t.Logf("%d: key %s, rank %d\n", i, res.Key, res.Rank)
	}

	t.Log("Top 10 out ranks:")
	for i, res := range g.TopNOutRanks(10) {
		if res.Rank != numNodes-i {
			t.Errorf("Expected rank %d but got %d", numNodes-i, res.Rank)
		}
		t.Logf("%d: key %s, rank %d\n", i, res.Key, res.Rank)
	}
}
