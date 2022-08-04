package sp

import (
	dfs "go_learning/algorithm/graph/ddfs"
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
	"math"
)

// AcyclicSP 无环有向图 单点最短路径
// 以拓扑排序的顶点顺序，依次放松所有顶点，每个顶点只会放松一次
type AcyclicSP[T comparable] struct {
	distTo *st.SeparateChainingHashST[T, float64]                // 从起点到顶点最短路径的权重
	edgeTo *st.SeparateChainingHashST[T, *graph.DirectedEdge[T]] // 从起点到顶点最短路径的边
}

func NewAcyclicSP[T comparable](G graph.EdgeWeightedDiGraph[T], startVertex T) *AcyclicSP[T] {
	vertexArr := G.V()
	lenVertexArr := len(vertexArr)
	distTo := st.NewSeparateChainingHashST[T, float64](lenVertexArr)
	edgeTo := st.NewSeparateChainingHashST[T, *graph.DirectedEdge[T]](lenVertexArr)
	for _, vertex := range vertexArr {
		distTo.Put(vertex, math.MaxFloat64) // 初始化每个顶点的最短路径权重为max
	}
	res := &AcyclicSP[T]{
		distTo: distTo,
		edgeTo: edgeTo,
	}

	// 以拓扑排序的顶点顺序，依次放松所有顶点
	res.distTo.Put(startVertex, 0)
	topo := dfs.NewTopologicalDWG[T](G)
	order := topo.Order()
	for order.HasNext() {
		res.Relax(G, order.Next())
	}
	return res
}

// HasPathTo 是否存在到顶点vertex的路径
func (d *AcyclicSP[T]) HasPathTo(vertex T) bool {
	return d.distTo.Get(vertex) != math.MaxFloat64
}

// DistTo 到顶点vertex的权重
func (d *AcyclicSP[T]) DistTo(vertex T) float64 {
	return d.distTo.Get(vertex)
}

// PathTo 到顶点vertex的路径
func (d *AcyclicSP[T]) PathTo(vertex T) common.Iterable[*graph.DirectedEdge[T]] {
	iter := stack.NewLIFOStack[*graph.DirectedEdge[T]]()
	for e := d.edgeTo.Get(vertex); e != nil; e = d.edgeTo.Get(e.From()) {
		iter.Push(e)
	}
	return iter.NewIterator()
}

// Relax 放松顶点 (找到一条比从起点到vertex更短的路径)
// 类似于橡皮筋，绷紧的橡皮筋两点之间距离较远，放松的橡皮筋两点之间较近
func (d *AcyclicSP[T]) Relax(G graph.EdgeWeightedDiGraph[T], vertex T) {
	// 依次放松相邻的边
	edges := G.Adj(vertex)
	for edges.HasNext() {
		edge := edges.Next()
		to := edge.To()

		// 如果到 to 的权重比 到vertex的权重+当前边的权重更大，那么就是新的到to的最短路径。
		// 将到 to 的路径更新，之前的路径就失效了
		if d.distTo.Get(to) > d.distTo.Get(vertex)+edge.Weight() {
			d.distTo.Put(to, d.distTo.Get(vertex)+edge.Weight())
			d.edgeTo.Put(to, &edge)
		}
	}
}
