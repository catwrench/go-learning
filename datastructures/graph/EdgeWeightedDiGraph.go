package graph

import (
	"go_learning/common"
	"go_learning/datastructures/bag"
	"go_learning/datastructures/st"
)

// EdgeWeightedDiGraph 加权有向图
type EdgeWeightedDiGraph[T comparable] struct {
	vertex  int                                                  // 顶点数
	edge    int                                                  // 边数量
	adj     *st.SequentialSearchST[T, *bag.Bag[DirectedEdge[T]]] // 邻接表，key:顶点，value:使用背包存放顶点的边集
	idxIter common.Iterable[DirectedEdge[T]]                     // 遍历背包的迭代器
}

func NewEdgeWeightedDiGraph[T comparable]() *EdgeWeightedDiGraph[T] {
	return &EdgeWeightedDiGraph[T]{
		adj: st.NewSequentialSearchST[T, *bag.Bag[DirectedEdge[T]]](),
	}
}

// AddEdge 添加带权重的边
func (g *EdgeWeightedDiGraph[T]) AddEdge(edge DirectedEdge[T]) {
	vertex1 := edge.From()
	// 第一次添加顶点的边集时，需要先初始化背包
	if !g.adj.Contains(vertex1) {
		g.vertex++
		g.adj.Put(vertex1, bag.NewBag[DirectedEdge[T]]().Add(edge))
	} else {
		g.adj.Get(vertex1).Add(edge)
	}

	// 初始化一下to顶点，避免adj遍历不到
	if !g.adj.Contains(edge.To()) {
		g.adj.Put(edge.To(), bag.NewBag[DirectedEdge[T]]())
	}
	g.edge++
}

// V 返回顶点数组
func (g *EdgeWeightedDiGraph[T]) V() []T {
	res := make([]T, 0, g.vertex)
	for _, v := range g.adj.Keys() {
		res = append(res, v)
	}
	return res
}

// Edges 返回所有边的迭代器
func (g *EdgeWeightedDiGraph[T]) Edges() common.Iterable[DirectedEdge[T]] {
	edgesBag := bag.NewBag[DirectedEdge[T]]()
	for _, vertex := range g.V() {
		it := g.Adj(vertex)
		for it.HasNext() {
			edgesBag.Add(it.Next()) // ps: 未实现顶点的compareTo, 所以会存在重复边
		}
	}
	return edgesBag
}

// Adj 创建某个顶点的迭代器,遍历从该顶点出发的所有边
func (g *EdgeWeightedDiGraph[T]) Adj(vertex T) common.Iterable[DirectedEdge[T]] {
	g.idxIter = g.adj.Get(vertex).NewIterator()
	return g
}

func (g *EdgeWeightedDiGraph[T]) HasNext() bool {
	return g.idxIter.HasNext()
}

func (g *EdgeWeightedDiGraph[T]) Next() DirectedEdge[T] {
	return g.idxIter.Next()
}
