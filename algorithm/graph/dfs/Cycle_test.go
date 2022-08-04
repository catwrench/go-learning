package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {

	G := graph.CreateGraph()
	cycle := NewCycle(*G)

	// 1-2-3 , 1-3 构成环
	assert.True(t, cycle.HasCycle())
}
