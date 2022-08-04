package graph

import (
	"go_learning/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphAdj(t *testing.T) {
	g := CreateGraph()

	// 顶点1的边集
	str := getGraphStr(g.Adj("1"))
	assert.Equal(t, "8432", str)

	str = getGraphStr(g.Adj("3"))
	assert.Equal(t, "5142", str)

	str = getGraphStr(g.Adj("0"))
	assert.Equal(t, "", str)

	str = getGraphStr(g.Adj("9"))
	assert.Equal(t, "87", str)
}

func getGraphStr(it common.Iterable[string]) string {
	str := ""
	for it.HasNext() {
		str += it.Next()
	}
	return str
}
