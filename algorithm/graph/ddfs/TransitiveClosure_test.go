package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransitiveClosure(t *testing.T) {
	G := graph.CreateDiGraph()

	tc := NewTransitiveClosure(*G)
	assert.True(t, tc.Reachable("1", "2"))
	assert.True(t, tc.Reachable("1", "8"))
	assert.True(t, tc.Reachable("3", "7"))
	assert.False(t, tc.Reachable("1", "0"))
}
