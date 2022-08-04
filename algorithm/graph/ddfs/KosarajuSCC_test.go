package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCCCount(t *testing.T) {
	// 0 为自环
	G := graph.CreateDiGraph()
	cc := NewKosarajuSCC[string](*G)
	assert.Equal(t, 2, cc.Count())

	G = graph.InitCycleDiGraph()
	cc = NewKosarajuSCC[string](*G)
	assert.Equal(t, 2, cc.Count())
}

func TestCCConnected(t *testing.T) {
	G := graph.InitCycleDiGraph()

	cc := NewKosarajuSCC[string](*G)
	assert.True(t, cc.StronglyConnected("1", "9"))
	assert.True(t, cc.StronglyConnected("2", "9"))
	assert.True(t, cc.StronglyConnected("3", "7"))
	assert.False(t, cc.StronglyConnected("1", "0"))

}
