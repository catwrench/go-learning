package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopological(t *testing.T) {
	CG := graph.InitCycleDiGraph()
	tp := NewTopological(*CG)
	assert.False(t, tp.IsDAG())

	G := graph.CreateDiGraph()
	tp = NewTopological(*G)
	assert.True(t, tp.IsDAG())

	order := make([]string, 0)
	it := tp.Order()
	for it.HasNext() {
		order = append(order, it.Next())
	}
	assert.Equal(t, "1234567980", strings.Join(order, ""))
}
