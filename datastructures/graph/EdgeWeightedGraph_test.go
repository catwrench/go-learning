package graph

import (
	"go_learning/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeWeightedGraphAdj(t *testing.T) {
	g := CreateEdgeWeightGraph()

	// 顶点1的边集
	str := getEdgeWeightedGraphStr("1", g.Adj("1"))
	assert.Equal(t, "952", str)

	str = getEdgeWeightedGraphStr("3", g.Adj("3"))
	assert.Equal(t, "42", str)

	str = getEdgeWeightedGraphStr("7", g.Adj("7"))
	assert.Equal(t, "86", str)

	str = getEdgeWeightedGraphStr("9", g.Adj("9"))
	assert.Equal(t, "158", str)
}

func getEdgeWeightedGraphStr(vertex string, it common.Iterable[Edge[string]]) string {
	str := ""
	for it.HasNext() {
		edge := it.Next()
		vertex1 := edge.Either()
		if vertex1 == vertex {
			vertex2, _ := edge.Other(vertex1)
			str += vertex2
		} else {
			str += vertex1
		}
	}
	return str
}
