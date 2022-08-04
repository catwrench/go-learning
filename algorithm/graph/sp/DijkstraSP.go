package sp

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/others"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
	"math"
)

// DijkstraSP 单点最短路径-迪杰斯特拉算法
type DijkstraSP[T comparable] struct {
	distTo *st.SeparateChainingHashST[T, float64]                // 从起点到顶点最短路径的权重
	edgeTo *st.SeparateChainingHashST[T, *graph.DirectedEdge[T]] // 从起点到顶点最短路径的边
	minPQ  *others.SymbolIndexMinPQ[T, float64]                  // 索引最小优先队列，用来存储下一次需要遍历的顶点，以及到该顶点最短路径的权重
}

func NewDijkstraSP[T comparable](G graph.EdgeWeightedDiGraph[T], startVertex T) *DijkstraSP[T] {
	vertexArr := G.V()
	lenVertexArr := len(vertexArr)
	distTo := st.NewSeparateChainingHashST[T, float64](lenVertexArr)
	edgeTo := st.NewSeparateChainingHashST[T, *graph.DirectedEdge[T]](lenVertexArr)
	for _, vertex := range vertexArr {
		distTo.Put(vertex, math.MaxFloat64) // 初始化每个顶点的最短路径权重为max
	}
	res := &DijkstraSP[T]{
		distTo: distTo,
		edgeTo: edgeTo,
		minPQ:  others.NewSymbolIndexMinPQ[T, float64](lenVertexArr),
	}

	// 从起始顶点开始放松每个顶点
	res.distTo.Put(startVertex, 0)
	res.minPQ.Insert(startVertex, 0)
	for !res.minPQ.IsEmpty() {
		res.Relax(G, res.minPQ.DelMin())
	}
	return res
}

// HasPathTo 是否存在到顶点vertex的路径
func (d *DijkstraSP[T]) HasPathTo(vertex T) bool {
	return d.distTo.Get(vertex) != math.MaxFloat64
}

// DistTo 到顶点vertex的权重
func (d *DijkstraSP[T]) DistTo(vertex T) float64 {
	return d.distTo.Get(vertex)
}

// PathTo 到顶点vertex的路径
func (d *DijkstraSP[T]) PathTo(vertex T) common.Iterable[*graph.DirectedEdge[T]] {
	iter := stack.NewLIFOStack[*graph.DirectedEdge[T]]()
	for e := d.edgeTo.Get(vertex); e != nil; e = d.edgeTo.Get(e.From()) {
		iter.Push(e)
	}
	return iter.NewIterator()
}

// Relax 放松顶点 (找到一条比从起点到vertex更短的路径)
// 类似于橡皮筋，绷紧的橡皮筋两点之间距离较远，放松的橡皮筋两点之间较近
func (d *DijkstraSP[T]) Relax(G graph.EdgeWeightedDiGraph[T], vertex T) {
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

			// 相邻顶点入队，如果已在队列就更新权重
			if d.minPQ.Contain(to) {
				d.minPQ.Change(to, d.distTo.Get(to))
			} else {
				d.minPQ.Insert(to, d.distTo.Get(to))
			}
		}
	}
}
