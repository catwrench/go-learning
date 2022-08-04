package prim

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/others"
	"go_learning/datastructures/queue"
	"go_learning/datastructures/st"
)

// LazyPrimMST 最小生成树-延迟普利姆算法
// 从顶点开始，依次选取关联的最小权值的边
type LazyPrimMST[T comparable] struct {
	marked *st.SeparateChainingHashST[T, bool] // 最小生成树的顶点
	mst    *queue.Queue[*graph.Edge[T]]        // 最小生成树的的边

	// 因为 Go 1.18 目前还不支持实现 Operator（General notes on type sets）， 因此比如想针对某一自定义类型实现 Less，CompareTo 的自比较就会遇到问题
	// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#general-notes-on-type-sets
	// 单独实现了一个边的的最小优先队列 EdgeMinPQ，进行权值比较
	minPQ *others.EdgeMinPQ[T] // 最小生成树的横切边
}

func NewLazyPrimMST[T comparable](G graph.EdgeWeightedGraph[T]) *LazyPrimMST[T] {
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}
	res := &LazyPrimMST[T]{
		marked: marked,
		mst:    queue.NewQueue[*graph.Edge[T]](),
		minPQ:  others.NewEdgeMinPQ[T](),
	}

	// 图的预处理：从任意一个顶点开始遍历，假设图是联通的
	res.Visit(G, VertexArr[0])
	for !res.minPQ.IsEmpty() {
		// 最小权重的顶点出列
		edge := res.minPQ.DelMin()

		vertex1 := edge.Either()
		vertex2, _ := edge.Other(vertex1)
		// 跳过被访问过的顶点
		if res.Marked(vertex1) && res.Marked(vertex2) {
			continue
		}
		// 边 入队，继续访问未被标记的相邻顶点
		res.mst.EnQueue(edge)
		if !res.Marked(vertex1) {
			res.Visit(G, vertex1)
		}
		if !res.Marked(vertex2) {
			res.Visit(G, vertex2)

		}
	}
	return res
}

func (p *LazyPrimMST[T]) Visit(G graph.EdgeWeightedGraph[T], vertex T) {
	p.Mark(vertex)

	// 所有和顶点 vertex 相连的且未被访问过的顶点 入列
	it := G.Adj(vertex)
	for it.HasNext() {
		edge := it.Next()
		other, _ := edge.Other(vertex)
		if !p.Marked(other) {
			p.minPQ.Insert(&edge)
		}
	}
}

func (p *LazyPrimMST[T]) Mark(vertex T) {
	if p.marked.Get(vertex) {
		return
	}
	p.marked.Put(vertex, true)
}

func (p *LazyPrimMST[T]) Marked(vertex T) bool {
	return p.marked.Get(vertex)
}

// Mst 返回最小生成树
func (p *LazyPrimMST[T]) Mst() common.Iterable[*graph.Edge[T]] {
	return p.mst.NewIterator()
}

// Weight 返回最小生成树的权重之和
func (p *LazyPrimMST[T]) Weight() (weight float64) {
	it := p.Mst()
	for it.HasNext() {
		edge := it.Next()
		weight += edge.Weight()
	}
	return
}
