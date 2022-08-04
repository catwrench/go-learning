package prim

import (
	"go_learning/common"
	bag2 "go_learning/datastructures/bag"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/others"
	"go_learning/datastructures/st"
	"math"
)

// PrimMST 最小生成树-即时普利姆算法
type PrimMST[T comparable] struct {
	edgeTo *st.SequentialSearchST[T, *graph.Edge[T]] // 顶点 到最小生成树的边
	distTo *st.SeparateChainingHashST[T, float64]    // 顶点 到最小生成树的最小权值, distTo = edgeTo.get(Vertex).Weight()
	marked *st.SeparateChainingHashST[T, bool]       // 顶点 是否在最小生成树中
	minPQ  *others.SymbolIndexMinPQ[T, float64]      // 有效的横切边
}

func NewPrimMST[T comparable](G graph.EdgeWeightedGraph[T]) *PrimMST[T] {
	VertexArr := G.V()
	lenVertex := len(VertexArr)
	marked := st.NewSeparateChainingHashST[T, bool](lenVertex)
	distTo := st.NewSeparateChainingHashST[T, float64](lenVertex)

	for _, vertex := range VertexArr {
		marked.Put(vertex, false)           // 初始化顶点都未访问
		distTo.Put(vertex, math.MaxFloat64) // 初始化每个顶点到最小生成树的权值为 MAX
	}
	res := &PrimMST[T]{
		edgeTo: st.NewSequentialSearchST[T, *graph.Edge[T]](),
		distTo: distTo,
		marked: marked,
		minPQ:  others.NewSymbolIndexMinPQ[T, float64](lenVertex),
	}

	// 	从顶点0开始访问，设定权值为0
	res.distTo.Put(VertexArr[0], 0.0)
	res.minPQ.Insert(VertexArr[0], 0.0)

	for !res.minPQ.IsEmpty() {
		// 将最近的顶点添加到树中
		res.visit(G, res.minPQ.DelMin())
	}

	return res
}

func (p *PrimMST[T]) visit(G graph.EdgeWeightedGraph[T], vertex T) {
	p.marked.Put(vertex, true)

	it := G.Adj(vertex)
	for it.HasNext() {
		edge := it.Next()
		other, _ := edge.Other(vertex)
		if p.marked.Get(other) { // 跳过无效的边
			continue
		}

		// 如果边权重比『相邻顶点』到『最小生成树』的权重更小时，更新相邻顶点到最小生成树的数据，否则跳过
		if edge.Weight() < p.distTo.Get(other) {

			p.edgeTo.Put(other, &edge)
			p.distTo.Put(other, edge.Weight())

			if p.minPQ.Contain(other) {
				p.minPQ.Change(other, p.distTo.Get(other))
			} else {
				p.minPQ.Insert(other, p.distTo.Get(other))
			}
		}
	}
}

// Mst 返回最小生成树
func (p *PrimMST[T]) Mst() common.Iterable[*graph.Edge[T]] {
	mst := bag2.NewBag[*graph.Edge[T]]()
	it := p.edgeTo.NewIterator()
	for it.HasNext() {
		_, edge := it.Next()
		mst.Add(edge)
	}
	return mst
}

// Weight 返回最小生成树的权重之和
func (p *PrimMST[T]) Weight() (weight float64) {
	it := p.edgeTo.NewIterator()
	for it.HasNext() {
		_, edge := it.Next()
		weight += edge.Weight()
	}
	return
}
