package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
)

// DepthFirstPaths 深度搜索优先，路径查找
type DepthFirstPaths[T comparable] struct {
	marked *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count  int                                 // 访问过的顶点数
	start  T                                   // 起始顶点
	edgeTo *st.SequentialSearchST[T, T]        // 从顶点A到顶点Z之间，已知路径上的最后一个顶点（如已知ABC，那么这里就是C）
}

func NewDepthFirstPaths[T comparable](G graph.Graph[T]) *DepthFirstPaths[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	return &DepthFirstPaths[T]{
		marked: marked,
		edgeTo: st.NewSequentialSearchST[T, T](),
	}
}

// Search 从顶点开始搜索
func (d *DepthFirstPaths[T]) Search(G graph.Graph[T], startVertex T) {
	d.start = startVertex
	d.dfs(G, startVertex)
}

// dfs 从顶点开始深度优先搜索
func (d *DepthFirstPaths[T]) dfs(G graph.Graph[T], startVertex T) {
	d.Mark(startVertex)
	it := G.Adj(startVertex)
	for it.HasNext() {
		nextVertex := it.Next()
		// 未被标记就继续搜索下一个点，否则回退
		if !d.marked.Get(nextVertex) {
			d.edgeTo.Put(nextVertex, startVertex)
			d.dfs(G, nextVertex)
		}
	}
}

// Mark 标记顶点为已访问
func (d *DepthFirstPaths[T]) Mark(vertex T) {
	if d.marked.Get(vertex) {
		return
	}
	d.marked.Put(vertex, true)
	d.count++
}

// Count 访问过的顶点数
func (d *DepthFirstPaths[T]) Count() int {
	return d.count
}

// HasPathTo 是否存在到该顶点的路径
func (d *DepthFirstPaths[T]) HasPathTo(vertex T) bool {
	return d.marked.Get(vertex)
}

// PathTo 到目标起点的路径
func (d *DepthFirstPaths[T]) PathTo(vertex T) common.Iterable[T] {
	// 从目标顶点开始，逆向搜索边，找到起始点
	res := stack.NewLIFOStack[T]()
	for i := vertex; i != d.start; i = d.edgeTo.Get(i) {
		res.Push(i)
	}
	// 补充一个起始点
	res.Push(d.start)
	res.NewIterator()
	return res
}
