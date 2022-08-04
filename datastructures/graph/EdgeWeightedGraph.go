package graph

import (
	"go_learning/common"
	"go_learning/datastructures/bag"
	"go_learning/datastructures/st"
)

// EdgeWeightedGraph 加权无向图
type EdgeWeightedGraph[T comparable] struct {
	vertex  int                                          // 顶点数
	edge    int                                          // 边数量
	adj     *st.SequentialSearchST[T, *bag.Bag[Edge[T]]] // 邻接表，key:顶点，value:使用背包存放顶点的边集
	idxIter common.Iterable[Edge[T]]                     // 遍历背包的迭代器
}

func NewEdgeWeightedGraph[T comparable]() *EdgeWeightedGraph[T] {
	return &EdgeWeightedGraph[T]{
		adj: st.NewSequentialSearchST[T, *bag.Bag[Edge[T]]](),
	}
}

// AddEdge 添加带权重的边
func (g *EdgeWeightedGraph[T]) AddEdge(edge Edge[T]) {
	vertex1 := edge.Either()
	vertex2, _ := edge.Other(vertex1)

	// 第一次添加顶点的边集时，需要先初始化背包
	if !g.adj.Contains(vertex1) {
		g.vertex++
		g.adj.Put(vertex1, bag.NewBag[Edge[T]]().Add(edge))
	} else {
		g.adj.Get(vertex1).Add(edge)
	}
	if !g.adj.Contains(vertex2) {
		g.vertex++
		g.adj.Put(vertex2, bag.NewBag[Edge[T]]().Add(edge))
	} else {
		g.adj.Get(vertex2).Add(edge)
	}
	g.edge++
}

// V 返回顶点数组
func (g *EdgeWeightedGraph[T]) V() []T {
	res := make([]T, 0, g.vertex)
	for _, v := range g.adj.Keys() {
		res = append(res, v)
	}
	return res
}

// Edges 返回所有边的迭代器
func (g *EdgeWeightedGraph[T]) Edges() common.Iterable[Edge[T]] {
	edgesBag := bag.NewBag[Edge[T]]()
	for _, vertex := range g.V() {
		it := g.Adj(vertex)
		for it.HasNext() {
			edgesBag.Add(it.Next()) // ps: 未实现顶点的compareTo, 所以会存在重复边
		}
	}
	return edgesBag
}

// Adj 创建某个顶点的迭代器,遍历从该顶点出发的所有边
func (g *EdgeWeightedGraph[T]) Adj(vertex T) common.Iterable[Edge[T]] {
	g.idxIter = g.adj.Get(vertex).NewIterator()
	return g
}

func (g *EdgeWeightedGraph[T]) HasNext() bool {
	return g.idxIter.HasNext()
}

func (g *EdgeWeightedGraph[T]) Next() Edge[T] {
	return g.idxIter.Next()
}
