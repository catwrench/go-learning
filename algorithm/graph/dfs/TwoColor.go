package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// TwoColor 双色图
type TwoColor[T comparable] struct {
	marked     *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	isTowColor bool                                // 是否双色图
	vertexIdx  *st.SeparateChainingHashST[T, bool] // 顶点索引,key：顶点，value: 顶点颜色
}

func NewTwoColor[T comparable](G graph.Graph[T]) *TwoColor[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	res := &TwoColor[T]{
		marked:     st.NewSeparateChainingHashST[T, bool](len(VertexArr)),
		vertexIdx:  st.NewSeparateChainingHashST[T, bool](len(VertexArr)),
		isTowColor: true,
	}
	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	for i := range VertexArr {
		if !res.marked.Get(VertexArr[i]) {
			res.dfs(G, VertexArr[i])
		}
	}
	return res
}

func (c *TwoColor[T]) dfs(G graph.Graph[T], vertex T) {
	c.marked.Put(vertex, true)
	c.vertexIdx.Put(vertex, c.isTowColor)

	it := G.Adj(vertex)
	for it.HasNext() {
		nextVertex := it.Next()
		if !c.marked.Get(nextVertex) {
			// 如果颜色邻接顶点颜色不一致，就继续dfs
			c.vertexIdx.Put(nextVertex, !c.vertexIdx.Get(vertex))
			c.dfs(G, nextVertex)
		} else if c.vertexIdx.Get(nextVertex) == c.vertexIdx.Get(vertex) {
			// 邻接顶点颜色一致，记录为非双色图
			c.isTowColor = false
		}
	}
}

// IsTwoColor 判断是否双色图
func (c *TwoColor[T]) IsTwoColor() bool {
	return c.isTowColor
}
