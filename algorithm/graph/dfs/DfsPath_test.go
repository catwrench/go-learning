package dfs

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsPath(t *testing.T) {
	g := graph.CreateGraph()

	dfsPath := NewDepthFirstPaths[string](*g)
	dfsPath.Search(*g, "1")
	assert.True(t, dfsPath.HasPathTo("9"))
	assert.True(t, dfsPath.HasPathTo("1"))
	assert.False(t, dfsPath.HasPathTo("0"))

	str := ""
	it := dfsPath.PathTo("9")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "189", str)

	dfsPath = NewDepthFirstPaths[string](*g)
	dfsPath.Search(*g, "2")
	assert.True(t, dfsPath.HasPathTo("7"))
	str = ""
	it = dfsPath.PathTo("7")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "2357", str)
}
