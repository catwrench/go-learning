package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
)

// Topological 拓扑排序
type Topological[T comparable] struct {
	order common.Iterable[T]
}

func NewTopological[T comparable](G graph.DiGraph[T]) *Topological[T] {
	res := &Topological[T]{}

	// 先校验是否存在环，再获取深度优先搜索的逆后序顶点排序
	// 拓扑排序即为有向无环图的逆后序顶点排序
	dc := NewDirectedCycle[T](G)
	if !dc.HasCycle() {
		dfo := NewDepthFirstOrder[T](G)
		res.order = dfo.ReversePost()
	}
	return res
}

// IsDAG 是否存在环
func (t *Topological[T]) IsDAG() bool {
	return t.order != nil
}

// Order 返回拓扑顶点排序
func (t *Topological[T]) Order() common.Iterable[T] {
	return t.order
}
