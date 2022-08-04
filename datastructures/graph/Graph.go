package graph

import (
	"go_learning/common"
	"go_learning/datastructures/bag"
	"go_learning/datastructures/st"
)

// Graph 图（基于背包数据结构实现）
type Graph[T comparable] struct {
	vertex  int                                    // 顶点数
	edge    int                                    // 边数量
	adj     *st.SequentialSearchST[T, *bag.Bag[T]] // 邻接表（数组索引不方便用泛型，所以用的顺序查找符号表+背包，也可以换成其他符号表实现）
	idxIter common.Iterable[T]                     // 遍历背包的迭代器
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adj: st.NewSequentialSearchST[T, *bag.Bag[T]](),
	}
}

// AddVertex 添加顶点
func (g *Graph[T]) AddVertex(vertex T) {
	if vertexBag := g.adj.Get(vertex); vertexBag == nil {
		g.adj.Put(vertex, bag.NewBag[T]())
	}
}

// AddEdge 添加一条边
func (g *Graph[T]) AddEdge(vertex1, vertex2 T) {
	g.adj.Get(vertex1).Add(vertex2)
	g.adj.Get(vertex2).Add(vertex1)
	g.edge++
}

// V 返回顶点数组
func (g *Graph[T]) V() []T {
	res := make([]T, g.vertex)
	for _, v := range g.adj.Keys() {
		res = append(res, v)
	}
	return res
}

// Adj 创建某个顶点的迭代器,遍历从该顶点出发的所有边
func (g *Graph[T]) Adj(vertex T) common.Iterable[T] {
	g.idxIter = g.adj.Get(vertex).NewIterator()
	return g
}

func (g *Graph[T]) HasNext() bool {
	return g.idxIter.HasNext()
}

func (g Graph[T]) Next() T {
	return g.idxIter.Next()
}
