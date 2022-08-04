package graph

import (
	"go_learning/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeWeightedDiGraphAdj(t *testing.T) {
	g := CreateEdgeWeightDiGraph()

	// 顶点1的边集
	str := getEdgeWeightedDiGraphStr("1", g.Adj("1"))
	assert.Equal(t, "52", str)

	str = getEdgeWeightedDiGraphStr("3", g.Adj("3"))
	assert.Equal(t, "4", str)

	str = getEdgeWeightedDiGraphStr("7", g.Adj("7"))
	assert.Equal(t, "8", str)

	str = getEdgeWeightedDiGraphStr("9", g.Adj("9"))
	assert.Equal(t, "", str)
}

func getEdgeWeightedDiGraphStr(vertex string, it common.Iterable[DirectedEdge[string]]) string {
	str := ""
	for it.HasNext() {
		edge := it.Next()
		vertex1 := edge.From()
		if vertex1 == vertex {
			vertex2 := edge.To()
			str += vertex2
		} else {
			str += vertex1
		}
	}
	return str
}
