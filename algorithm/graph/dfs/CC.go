package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// CC 连通分量查找
type CC[T comparable] struct {
	marked    *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count     int                                 // 连通分量数
	vertexIdx *st.SeparateChainingHashST[T, int]  // 顶点索引,key：顶点，value: 顶点所属连通分量标识 0 到 count-1
}

func NewCC[T comparable](G graph.Graph[T]) *CC[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	res := &CC[T]{
		marked:    st.NewSeparateChainingHashST[T, bool](len(VertexArr)),
		vertexIdx: st.NewSeparateChainingHashST[T, int](len(VertexArr)),
	}
	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	for i := range VertexArr {
		if !res.marked.Get(VertexArr[i]) {
			res.dfs(G, VertexArr[i])
			res.count++
		}
	}
	return res
}

func (c *CC[T]) dfs(G graph.Graph[T], vertex T) {
	c.marked.Put(vertex, true)
	c.vertexIdx.Put(vertex, c.count)

	it := G.Adj(vertex)
	for it.HasNext() {
		nextVertex := it.Next()
		if !c.marked.Get(nextVertex) {
			c.dfs(G, nextVertex)
		}
	}
}

// Connected 判断两个顶点是否连通
func (c *CC[T]) Connected(vertex1, vertex2 T) bool {
	// 通过两个顶点对应的顶点索引标志是否一致进行判断，如果一致说明在同一个连通分量里
	return c.vertexIdx.Get(vertex1) == c.vertexIdx.Get(vertex2)
}

func (c *CC[T]) Count() int {
	return c.count
}

func (c *CC[T]) ID(vertex T) int {
	return c.vertexIdx.Get(vertex)
}
