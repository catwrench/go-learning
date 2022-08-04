package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
)

// TopologicalDWG 拓扑排序 带权有向图
type TopologicalDWG[T comparable] struct {
	order common.Iterable[T]
}

func NewTopologicalDWG[T comparable](G graph.EdgeWeightedDiGraph[T]) *TopologicalDWG[T] {
	res := &TopologicalDWG[T]{}

	// 略，先校验是否存在环，再获取深度优先搜索的逆后序顶点排序
	// 拓扑排序即为有向无环图的逆后序顶点排序
	dfo := NewDepthFirstOrderDWG[T](G)
	res.order = dfo.ReversePost()
	return res
}

// IsDAG 是否存在环
func (t *TopologicalDWG[T]) IsDAG() bool {
	return t.order != nil
}

// Order 返回拓扑顶点排序
func (t *TopologicalDWG[T]) Order() common.Iterable[T] {
	return t.order
}
