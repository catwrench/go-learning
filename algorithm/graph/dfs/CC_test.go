package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCCCount(t *testing.T) {
	G := graph.CreateGraph()

	cc := NewCC[string](*G)
	assert.Equal(t, 2, cc.Count())
}

func TestCCConnected(t *testing.T) {
	G := graph.CreateGraph()

	cc := NewCC[string](*G)
	assert.True(t, cc.Connected("1", "9"))
	assert.True(t, cc.Connected("2", "9"))
	assert.True(t, cc.Connected("3", "7"))
	assert.False(t, cc.Connected("1", "0"))

}
