package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
)

// EdgeWeightedDirectedCycle 寻找有向环
type EdgeWeightedDirectedCycle[T comparable] struct {
	marked  *st.SeparateChainingHashST[T, bool]                   // 用于记录访问过的点（hash符号表-拉链法）
	edgeTo  *st.SeparateChainingHashST[T, *graph.DirectedEdge[T]] // 从顶点A到顶点Z之间，已知路径上的最后一个顶点（如已知ABC，那么这里就是C）
	inCycle *st.SequentialSearchST[T, bool]                       // 顶点是否存在环中
	cycle   *stack.LIFOStack[*graph.DirectedEdge[T]]              // 环路径
}

func NewEdgeWeightedDirectedCycle[T comparable](G graph.EdgeWeightedDiGraph[T]) *EdgeWeightedDirectedCycle[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	edgeTo := st.NewSeparateChainingHashST[T, *graph.DirectedEdge[T]](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	res := &EdgeWeightedDirectedCycle[T]{
		marked:  marked,
		edgeTo:  edgeTo,
		inCycle: st.NewSequentialSearchST[T, bool](),
	}
	// 遍历图的顶点集，未访问过的都执行一次 dfs 预计算
	for i := range VertexArr {
		if !res.marked.Get(VertexArr[i]) {
			res.dfs(G, VertexArr[i])
		}
	}
	return res
}

// dfs 从顶点开始搜索
func (d *EdgeWeightedDirectedCycle[T]) dfs(G graph.EdgeWeightedDiGraph[T], startVertex T) {
	// 一开始先将顶点加入环中，如果最后回退了，再从环中移除（标记为false）
	d.inCycle.Put(startVertex, true)
	d.Mark(startVertex)

	it := G.Adj(startVertex)
	for it.HasNext() {
		edge := it.Next()
		to := edge.To()
		// 存在环就返回
		// 未被标记就继续搜索下一个点
		// 如果在环中，那么遍历edge写入到环
		if d.HasCycle() {
			return
		} else if !d.marked.Get(to) {
			d.edgeTo.Put(to, &edge)
			d.dfs(G, to)
		} else if d.inCycle.Get(to) {
			d.cycle = stack.NewLIFOStack[*graph.DirectedEdge[T]]()
			i := &edge
			for ; i.From() != to; i = d.edgeTo.Get(i.From()) {
				d.cycle.Push(i)
			}
			// 追加起始点,退出点
			d.cycle.Push(i)
			return
		}
	}
	d.inCycle.Put(startVertex, false)
}

// Mark 标记顶点为已访问
func (d *EdgeWeightedDirectedCycle[T]) Mark(vertex T) {
	if d.marked.Get(vertex) {
		return
	}
	d.marked.Put(vertex, true)
}

// Marked 顶点是否访问过
func (d *EdgeWeightedDirectedCycle[T]) Marked(vertex T) bool {
	return d.marked.Get(vertex)
}

// HasCycle 是否存在环
func (d *EdgeWeightedDirectedCycle[T]) HasCycle() bool {
	return d.cycle != nil
}

// Cycle 遍历环所有的顶点
func (d *EdgeWeightedDirectedCycle[T]) Cycle() common.Iterable[*graph.DirectedEdge[T]] {
	return d.cycle.NewIterator()
}
