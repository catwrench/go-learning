package graph

import (
	"go_learning/common"
	"go_learning/datastructures/bag"
	"go_learning/datastructures/st"
)

// DiGraph 有向图
type DiGraph[T comparable] struct {
	vertex  int                                    // 顶点数
	edge    int                                    // 边数量
	adj     *st.SequentialSearchST[T, *bag.Bag[T]] // 邻接表（数组索引不方便用泛型，所以用的顺序查找符号表+背包，也可以换成其他符号表实现）
	idxIter common.Iterable[T]                     // 遍历背包的迭代器
}

func NewDiGraph[T comparable]() *DiGraph[T] {
	res := &DiGraph[T]{}
	res.adj = st.NewSequentialSearchST[T, *bag.Bag[T]]()
	return res
}

// AddVertex 添加顶点
func (g *DiGraph[T]) AddVertex(vertex T) {
	if vertexBag := g.adj.Get(vertex); vertexBag == nil {
		g.adj.Put(vertex, bag.NewBag[T]())
	}
}

// AddEdge 添加一条边
func (g *DiGraph[T]) AddEdge(vertex1, vertex2 T) {
	if !g.adj.Contains(vertex1) {
		g.AddVertex(vertex1)
	}
	if !g.adj.Contains(vertex2) {
		g.AddVertex(vertex2)
	}
	g.adj.Get(vertex1).Add(vertex2)
	g.edge++
}

// V 返回顶点数组 (遍历时为逆序)
func (g *DiGraph[T]) V() []T {
	res := make([]T, g.vertex)
	for _, v := range g.adj.Keys() {
		res = append(res, v)
	}
	return res
}

// Adj 创建某个顶点的迭代器,遍历从该顶点出发的所有边
func (g *DiGraph[T]) Adj(vertex T) common.Iterable[T] {
	if !g.adj.Contains(vertex) {
		g.idxIter = bag.NewBag[T]().NewIterator()
	} else {
		g.idxIter = g.adj.Get(vertex).NewIterator()
	}
	return g
}

func (g *DiGraph[T]) HasNext() bool {
	return g.idxIter.HasNext()
}

func (g DiGraph[T]) Next() T {
	return g.idxIter.Next()
}

// Reverse 翻转图
func (g *DiGraph[T]) Reverse() *DiGraph[T] {
	reverseG := NewDiGraph[T]()
	// 重新生成顶点
	vertexes := g.V()
	for i := range vertexes {
		reverseG.AddVertex(vertexes[i])
	}
	// 反向生成边
	for i := range vertexes {
		edgeIt := g.Adj(vertexes[i])
		for edgeIt.HasNext() {
			reverseG.AddEdge(edgeIt.Next(), vertexes[i])
		}
	}
	return reverseG
}
