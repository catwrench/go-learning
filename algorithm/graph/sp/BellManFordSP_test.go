package sp

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBellManFordSP(t *testing.T) {
	G := graph.CreateEdgeWeightDiGraphForBell()
	sp := NewBellManFordSP[int](*G, 0)

	assert.True(t, sp.HasPathTo(7))

	it, ok := sp.PathTo(7)
	assert.True(t, ok)
	path := make([]int, 0)
	for it.HasNext() {
		edge := it.Next()
		path = append(path, edge.To())
	}
	assert.Equal(t, []int{2, 7}, path)
}

func TestBellManFordSPCycle(t *testing.T) {
	G := graph.CreateEdgeWeightDiGraphForBell()
	sp := NewBellManFordSP[int](*G, 0)

	assert.True(t, sp.HasPathTo(5))
	it, _ := sp.PathTo(5)
	path := make([]int, 0)
	for it.HasNext() {
		edge := it.Next()
		path = append(path, edge.To())
	}
	assert.Equal(t, []int{2, 7, 3, 6, 4, 5}, path)

	assert.False(t, sp.HasNegativeCycle())
}
