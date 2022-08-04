package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

type DeepFirstSearch[T comparable] struct {
	marked *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count  int                                 // 访问过的顶点数
}

func NewDeepFirstSearch[T comparable](G graph.Graph[T]) *DeepFirstSearch[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	return &DeepFirstSearch[T]{
		marked: marked,
	}
}

// Search 从顶点开始搜索
func (d *DeepFirstSearch[T]) Search(G graph.Graph[T], startVertex T) {
	d.Mark(startVertex)
	it := G.Adj(startVertex)
	for it.HasNext() {
		nextVertex := it.Next()
		// 未被标记就继续搜索下一个点，否则回退
		if !d.marked.Get(nextVertex) {
			d.Search(G, nextVertex)
		}
	}
}

// Mark 标记顶点为已访问
func (d *DeepFirstSearch[T]) Mark(vertex T) {
	if d.marked.Get(vertex) {
		return
	}
	d.marked.Put(vertex, true)
	d.count++
}

// Count 访问过的顶点数
func (d *DeepFirstSearch[T]) Count() int {
	return d.count
}
