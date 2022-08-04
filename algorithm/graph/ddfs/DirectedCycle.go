package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
)

// DirectedCycle 寻找有向环
type DirectedCycle[T comparable] struct {
	marked  *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count   int                                 // 访问过的顶点数
	edgeTo  *st.SequentialSearchST[T, T]        // 从顶点A到顶点Z之间，已知路径上的最后一个顶点（如已知ABC，那么这里就是C）
	inCycle *st.SequentialSearchST[T, bool]     // 顶点是否存在环中
	cycle   *stack.LIFOStack[T]                 // 环路径
}

func NewDirectedCycle[T comparable](G graph.DiGraph[T]) *DirectedCycle[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	res := &DirectedCycle[T]{
		marked:  marked,
		edgeTo:  st.NewSequentialSearchST[T, T](),
		inCycle: st.NewSequentialSearchST[T, bool](),
		cycle:   stack.NewLIFOStack[T](),
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

// dfs 从顶点开始搜索
func (d *DirectedCycle[T]) dfs(G graph.DiGraph[T], startVertex T) {
	// 一开始先将顶点加入环中，如果最后回退了，再从环中移除（标记为false）
	d.inCycle.Put(startVertex, true)

	d.Mark(startVertex)
	it := G.Adj(startVertex)
	for it.HasNext() {
		nextVertex := it.Next()
		// 存在环就返回
		// 未被标记就继续搜索下一个点
		// 如果在环中，那么遍历edge写入到环
		if d.HasCycle() {
			return
		} else if !d.marked.Get(nextVertex) {
			d.edgeTo.Put(nextVertex, startVertex)
			d.dfs(G, nextVertex)
		} else if d.inCycle.Get(nextVertex) {
			for i := startVertex; i != nextVertex; i = d.edgeTo.Get(i) {
				d.cycle.Push(i)
			}
			// 追加起始点,退出点
			d.cycle.Push(nextVertex)
			d.cycle.Push(startVertex)
		}
	}
	d.inCycle.Put(startVertex, false)
}

// Mark 标记顶点为已访问
func (d *DirectedCycle[T]) Mark(vertex T) {
	if d.marked.Get(vertex) {
		return
	}
	d.marked.Put(vertex, true)
	d.count++
}

// Marked 顶点是否访问过
func (d *DirectedCycle[T]) Marked(vertex T) bool {
	return d.marked.Get(vertex)
}

// Count 访问过的顶点数
func (d *DirectedCycle[T]) Count() int {
	return d.count
}

// HasCycle 是否存在环
func (d *DirectedCycle[T]) HasCycle() bool {
	return !d.cycle.IsEmpty()
}

// Cycle 遍历环所有的顶点
func (d *DirectedCycle[T]) Cycle() common.Iterable[T] {
	return d.cycle.NewIterator()
}
