package kruskal

import (
	"go_learning/algorithm/graph/unionfind"
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/others"
	"go_learning/datastructures/queue"
)

// KruskalMST 最小生成树-克鲁斯卡尔算法
type KruskalMST[T comparable] struct {
	mst   *queue.Queue[*graph.Edge[T]] // 队列，用来保存最小生成树的所有边
	minPQ *others.EdgeMinPQ[T]         // 最小优先队列，将所有边按权重从低到高排列
	uf    *unionfind.QuickUnion        // QuickUnion 连通分量查找算法，用于判断选择的边是否在同一个连通分量里
}

// NewKruskalMST
// ps: 懒狗不想改 QuickUnion 了 ಥ_ಥ，所以这里全部用 int
func NewKruskalMST[T int](G graph.EdgeWeightedGraph[int]) *KruskalMST[int] {
	vertexArr := G.V()
	res := &KruskalMST[int]{
		mst:   queue.NewQueue[*graph.Edge[int]](),
		minPQ: others.NewEdgeMinPQ[int](),
		uf:    unionfind.NewQuickUnion(len(vertexArr)),
	}

	// 将所有边加入最小优先队列
	edges := G.Edges()
	for edges.HasNext() {
		edge := edges.Next()
		res.minPQ.Insert(&edge)
	}

	// 通过quickUnion算法判断边的两个顶点是否为同一连通分量，如果不是的话就加入最小生成树
	for !res.minPQ.IsEmpty() {
		edge := res.minPQ.DelMin()
		vertex := edge.Either()
		other, _ := edge.Other(vertex)
		if res.uf.Connected(vertex, other) {
			continue
		}
		res.uf.Union(vertex, other)
		res.mst.EnQueue(edge)
	}
	return res
}

// Mst 返回最小生成树
func (k *KruskalMST[T]) Mst() common.Iterable[*graph.Edge[T]] {
	return k.mst.NewIterator()
}

// Weight 返回最小生成树的权重之和
func (k *KruskalMST[T]) Weight() (weight float64) {
	it := k.Mst()
	for it.HasNext() {
		edge := it.Next()
		weight += edge.Weight()
	}
	return
}
