package sp

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
)

// ISP 有向图最短路径搜索接口
type ISP[T comparable] interface {
	HasPathTo(vertex T)
	DistTo(vertex T) float64
	PathTo(vertex T) common.Iterable[*graph.DirectedEdge[T]]
}
