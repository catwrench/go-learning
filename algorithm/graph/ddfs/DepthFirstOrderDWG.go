package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/queue"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
)

// DepthFirstOrderDWG 带权有向图基于深度优先搜索的顶点排序
type DepthFirstOrderDWG[T comparable] struct {
	marked      *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	pre         *queue.Queue[T]                     // 顶点前序排列
	post        *queue.Queue[T]                     // 顶点后序排序
	reversePost *stack.LIFOStack[T]                 // 顶点逆后序排列
}

func NewDepthFirstOrderDWG[T comparable](G graph.EdgeWeightedDiGraph[T]) *DepthFirstOrderDWG[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	res := &DepthFirstOrderDWG[T]{
		marked:      marked,
		pre:         queue.NewQueue[T](),
		post:        queue.NewQueue[T](),
		reversePost: stack.NewLIFOStack[T](),
	}
	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	for i := range VertexArr {
		if !res.marked.Get(VertexArr[i]) {
			res.dfs(G, VertexArr[i])
		}
	}
	return res
}

func (d *DepthFirstOrderDWG[T]) dfs(G graph.EdgeWeightedDiGraph[T], vertex T) {
	d.pre.EnQueue(vertex)

	d.marked.Put(vertex, true)
	G.Adj(vertex)
	for G.HasNext() {
		nextVertex := G.Next()
		if !d.marked.Get(nextVertex.To()) {
			d.dfs(G, nextVertex.To())
		}
	}

	d.post.EnQueue(vertex)
	d.reversePost.Push(vertex)
}

// Pre 返回顶点的前序遍历
func (d *DepthFirstOrderDWG[T]) Pre() common.Iterable[T] {
	return d.pre.NewIterator()
}

// Post 返回顶点的后序遍历
func (d *DepthFirstOrderDWG[T]) Post() common.Iterable[T] {
	return d.post.NewIterator()
}

// ReversePost 返回顶点的逆后续遍历
func (d *DepthFirstOrderDWG[T]) ReversePost() common.Iterable[T] {
	return d.reversePost.NewIterator()
}
