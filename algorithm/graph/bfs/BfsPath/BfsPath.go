package BfsPath

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	queue2 "go_learning/datastructures/queue"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
)

// BreadthFirstPaths 广度搜索优先路径
type BreadthFirstPaths[T comparable] struct {
	marked *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count  int                                 // 访问过的顶点数
	start  T                                   // 起始顶点
	edgeTo *st.SequentialSearchST[T, T]
}

func NewBreadthFirstPaths[T comparable](G graph.Graph[T]) *BreadthFirstPaths[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	return &BreadthFirstPaths[T]{
		marked: marked,
		edgeTo: st.NewSequentialSearchST[T, T](),
	}
}

// Search 从顶点开始搜索
func (b *BreadthFirstPaths[T]) Search(G graph.Graph[T], startVertex T) *BreadthFirstPaths[T] {
	b.start = startVertex
	b.bfs(G, startVertex)
	return b
}

// bfs 从顶点开始广度优先搜索
func (b *BreadthFirstPaths[T]) bfs(G graph.Graph[T], startVertex T) {
	// 起始节点标记已访问后，入队
	b.Marked(startVertex)
	queue := queue2.NewQueue[T]()
	queue.EnQueue(startVertex)

	// 队列不为空就出队一个元素
	for !queue.IsEmpty() {
		// 遍历出队顶点的相邻顶点，如果未访问则标记访问，记录访问路径，然后入队
		popVertex := queue.DeQueue()

		it := G.Adj(popVertex)
		for it.HasNext() {
			nextVertex := it.Next()
			if !b.marked.Get(nextVertex) {
				b.Marked(nextVertex)
				b.edgeTo.Put(nextVertex, popVertex)
				queue.EnQueue(nextVertex)
			}
		}
	}
}

// Marked 标记顶点为已访问
func (b *BreadthFirstPaths[T]) Marked(vertex T) {
	if b.marked.Get(vertex) {
		return
	}
	b.marked.Put(vertex, true)
	b.count++
}

// Count 访问过的顶点数
func (b *BreadthFirstPaths[T]) Count() int {
	return b.count
}

// HasPathTo 是否存在到该顶点的路径
func (b *BreadthFirstPaths[T]) HasPathTo(vertex T) bool {
	return b.marked.Get(vertex)
}

// PathTo 到目标起点的路径
func (b *BreadthFirstPaths[T]) PathTo(vertex T) common.Iterable[T] {
	// 从目标顶点开始，逆向搜索边，找到起始点
	res := stack.NewLIFOStack[T]()
	for i := vertex; i != b.start; i = b.edgeTo.Get(i) {
		res.Push(i)
	}
	// 补充一个起始点
	res.Push(b.start)
	res.NewIterator()
	return res
}
