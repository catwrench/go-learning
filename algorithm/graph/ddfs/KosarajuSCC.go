package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// KosarajuSCC 有向图 强连通分量查找
type KosarajuSCC[T comparable] struct {
	marked    *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count     int                                 // 连通分量数
	vertexIdx *st.SeparateChainingHashST[T, int]  // 顶点索引,key：顶点，value: 顶点所属连通分量标识 0 到 count-1
}

func NewKosarajuSCC[T comparable](G graph.DiGraph[T]) *KosarajuSCC[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	res := &KosarajuSCC[T]{
		marked:    st.NewSeparateChainingHashST[T, bool](len(VertexArr)),
		vertexIdx: st.NewSeparateChainingHashST[T, int](len(VertexArr)),
	}

	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	// 和无向图不一样的地方是，这里是以dfs逆后序顶点排序，作为遍历顺序
	order := NewDepthFirstOrder[T](G)
	it := order.ReversePost()
	for it.HasNext() {
		nextVertex := it.Next()
		if !res.marked.Get(nextVertex) {
			res.dfs(G, nextVertex)
			res.count++
		}
	}
	return res
}

func (c *KosarajuSCC[T]) dfs(G graph.DiGraph[T], vertex T) {
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

// StronglyConnected 判断两个顶点是否强连通
func (c *KosarajuSCC[T]) StronglyConnected(vertex1, vertex2 T) bool {
	// 通过两个顶点对应的顶点索引标志是否一致进行判断，如果一致说明在同一个连通分量里
	return c.vertexIdx.Get(vertex1) == c.vertexIdx.Get(vertex2)
}

func (c *KosarajuSCC[T]) Count() int {
	return c.count
}

func (c *KosarajuSCC[T]) ID(vertex T) int {
	return c.vertexIdx.Get(vertex)
}
