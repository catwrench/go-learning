package sp

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"math"
)

// Arbitrage 套汇获利问题
// 计算加权有向图的负权重环
type Arbitrage[T comparable] struct {
	bell *BellManFordSP[T]
}

func NewArbitrage[T comparable](G graph.EdgeWeightedDiGraph[T]) *Arbitrage[T] {
	// 对汇率取自然对数并取反，将汇率之积问题转换为汇率之和问题
	newG := graph.NewEdgeWeightedDiGraph[T]()
	vertexArr := G.V()
	for _, vertex := range vertexArr {
		it := G.Adj(vertex)
		for it.HasNext() {
			edge := it.Next()
			newG.AddEdge(*graph.NewDirectedEdge[T](edge.From(), edge.To(), -math.Log(edge.Weight())))
		}
	}

	return &Arbitrage[T]{
		bell: NewBellManFordSP[T](*newG, vertexArr[0]),
	}
}

// GetPath 获取套汇路径
// 通过BellmanFord算法获取到的负权重环，即为套汇获利路径
func (a *Arbitrage[T]) GetPath() common.Iterable[*graph.DirectedEdge[T]] {
	return a.bell.cycle
}
