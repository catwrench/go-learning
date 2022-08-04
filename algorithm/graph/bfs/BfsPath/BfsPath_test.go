package BfsPath

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func input() (V []string, edge []string) {
	V = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	// 没有到0的边
	edge = []string{
		"1,2", "2,3", "3,4", "4,5", "5,6", "6,7", "7,8",
		"1,3", "3,5", "5,7", "7,9",
		"1,4", "1,8", "9,8",
	}
	return
}

func CreateGraph() *graph.Graph[string] {
	g := graph.NewGraph[string]()
	V, edgeArr := input()
	for i := range V {
		g.AddVertex(V[i])
	}
	for _, edgeItem := range edgeArr {
		vArr := strings.Split(edgeItem, ",")
		g.AddEdge(vArr[0], vArr[1])
	}
	return g
}

func TestDfsPath(t *testing.T) {
	g := CreateGraph()

	bfsPath := NewBreadthFirstPaths[string](*g)
	bfsPath.Search(*g, "1")
	assert.True(t, bfsPath.HasPathTo("9"))
	assert.True(t, bfsPath.HasPathTo("1"))
	assert.False(t, bfsPath.HasPathTo("0"))

	str := ""
	it := bfsPath.PathTo("9")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "189", str)

	bfsPath = NewBreadthFirstPaths[string](*g)
	bfsPath.Search(*g, "2")
	assert.True(t, bfsPath.HasPathTo("7"))
	str = ""
	it = bfsPath.PathTo("7")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "2357", str)
}
