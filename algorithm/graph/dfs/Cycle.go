package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// Cycle 查找图里是否存在环
type Cycle[T comparable] struct {
	marked   *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	hasCycle bool
}

func NewCycle[T comparable](G graph.Graph[T]) *Cycle[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	res := &Cycle[T]{
		marked: st.NewSeparateChainingHashST[T, bool](len(VertexArr)),
	}
	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	for _, vertex := range VertexArr {
		if !res.marked.Get(vertex) {
			res.dfs(G, vertex, vertex)
		}
	}
	return res
}

// vertex:当前访问的顶点
// vertexU:搜索的目标顶点，二次访问就构成环
func (c *Cycle[T]) dfs(G graph.Graph[T], vertex, vertexU T) {
	c.marked.Put(vertex, true)

	it := G.Adj(vertex)
	for it.HasNext() {
		nextVertex := it.Next()
		if !c.marked.Get(nextVertex) {
			c.dfs(G, nextVertex, vertex)
		} else if nextVertex != vertexU {
			c.hasCycle = true
		}
	}
}

func (c *Cycle[T]) HasCycle() bool {
	return c.hasCycle
}
