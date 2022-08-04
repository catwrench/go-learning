package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutDirectedCycle(t *testing.T) {
	g := graph.CreateDiGraph()
	dc := NewDirectedCycle(*g)

	assert.False(t, dc.HasCycle())

	path := make([]string, 0)
	it := dc.Cycle()
	for it.HasNext() {
		path = append(path, it.Next())
	}
	assert.Equal(t, "", strings.Join(path, ""))

}

func TestDirectedCycle(t *testing.T) {
	g := graph.InitCycleDiGraph()
	dc := NewDirectedCycle(*g)

	assert.True(t, dc.HasCycle())

	path := make([]string, 0)
	it := dc.Cycle()
	for it.HasNext() {
		path = append(path, it.Next())
	}

	// 顶点遍历时为逆序，9开始,dfs->12345678
	// 所以写入到edge的数据为87654321，入inCycle栈后为12345678，加入头尾8912345678
	assert.Equal(t, "8912345678", strings.Join(path, ""))

}
