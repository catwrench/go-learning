package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoColor(t *testing.T) {

	G := graph.CreateGraph()
	cycle := NewTwoColor(*G)

	assert.False(t, cycle.IsTwoColor())
}
