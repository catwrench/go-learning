package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutWeightedDirectedCycle(t *testing.T) {
	g := graph.CreateEdgeWeightDiGraph()
	dc := NewEdgeWeightedDirectedCycle[string](*g)

	assert.False(t, dc.HasCycle())
}

func TestWeightedDirectedCycle(t *testing.T) {
	g := graph.CreateEdgeWeightDiGraph2()
	dc := NewEdgeWeightedDirectedCycle(*g)

	assert.True(t, dc.HasCycle())

	path := make([]string, 0)
	it := dc.Cycle()
	for it.HasNext() {
		path = append(path, it.Next().To())
	}
	assert.Equal(t, "159", strings.Join(path, ""))
}
