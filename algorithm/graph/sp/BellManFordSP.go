package sp

import (
	dfs "go_learning/algorithm/graph/ddfs"
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/queue"
	"go_learning/datastructures/st"
	"go_learning/datastructures/stack"
	"math"
)

// BellManFordSP 基于队列的BellMan-Ford算法
// 在任意一轮中，有许多边的放松都不会成功，只有上一轮中distTo值发生变化的顶点 指出的边才能改变其他distTo元素的值，
// BellMan-Ford算法用一个队列来记录这些顶点
type BellManFordSP[T comparable] struct {
	distTo *st.SeparateChainingHashST[T, float64]                // 从起点到某个顶点的路径长度
	edgeTo *st.SeparateChainingHashST[T, *graph.DirectedEdge[T]] // 从起点到某个顶点的最后一条边
	pq     *queue.Queue[T]                                       // 正在被放松的顶点
	onPQ   *st.SeparateChainingHashST[T, bool]                   // 顶点是否在被放松
	cost   int                                                   // 放松次数
	cycle  common.Iterable[*graph.DirectedEdge[T]]               // edgeTo 中是否存在负权重环
}

func NewBellManFordSP[T comparable](G graph.EdgeWeightedDiGraph[T], startVertex T) *BellManFordSP[T] {
	vertexArr := G.V()
	lenVertexArr := len(vertexArr)
	distTo := st.NewSeparateChainingHashST[T, float64](lenVertexArr)
	edgeTo := st.NewSeparateChainingHashST[T, *graph.DirectedEdge[T]](lenVertexArr)
	onPQ := st.NewSeparateChainingHashST[T, bool](lenVertexArr)
	for _, vertex := range vertexArr {
		distTo.Put(vertex, math.Inf(1)) // 初始化每个顶点的最短路径权重为max
		onPQ.Put(vertex, false)
	}

	res := &BellManFordSP[T]{
		distTo: distTo,
		edgeTo: edgeTo,
		pq:     queue.NewQueue[T](),
		onPQ:   onPQ,
	}

	// 从起始顶点开始放松 lenVertexArr 轮
	distTo.Put(startVertex, 0)
	res.pq.EnQueue(startVertex)
	res.onPQ.Put(startVertex, true)
	for !res.pq.IsEmpty() && !res.HasNegativeCycle() {
		nextVertex := res.pq.DeQueue()
		res.onPQ.Put(nextVertex, false)
		res.Relax(G, nextVertex)
	}

	return res
}

// Relax 放松顶点 (找到一条比从起点到vertex更短的路径)
func (b *BellManFordSP[T]) Relax(G graph.EdgeWeightedDiGraph[T], vertex T) {
	// 依次放松相邻的边
	edges := G.Adj(vertex)
	for edges.HasNext() {
		edge := edges.Next()
		to := edge.To()

		// 如果到 to 的权重比 到vertex的权重+当前边的权重更大，那么就是新的到to的最短路径。
		// 将到 to 的路径更新，之前的路径就失效了
		if b.distTo.Get(to) > b.distTo.Get(vertex)+edge.Weight() {
			b.distTo.Put(to, b.distTo.Get(vertex)+edge.Weight())
			b.edgeTo.Put(to, &edge)

			// 如果相邻顶点未在放松队列里，标记入队
			if !b.onPQ.Get(to) {
				b.pq.EnQueue(to)
				b.onPQ.Put(to, true)
			}
		}

		// 查找负权重环
		if b.cost%len(G.V()) == 0 {
			b.FindNegativeCycle()
			if b.HasNegativeCycle() { // 找到负权重环返回
				return
			}
		}
	}
}

// FindNegativeCycle 寻找负权重环
func (b *BellManFordSP[T]) FindNegativeCycle() {
	// 如果存在负权重环，那么edgeTo所表现的子图中必然含有这个负权重环，
	// 所以基于edgeTo中的边构建一幅加权有向图，然后在图中检测有向环。
	spt := graph.NewEdgeWeightedDiGraph[T]()
	keys := b.edgeTo.Keys()
	for _, key := range keys {
		if b.edgeTo.Get(key) != nil {
			spt.AddEdge(*b.edgeTo.Get(key))
		}
	}
	finder := dfs.NewEdgeWeightedDirectedCycle[T](*spt)
	if finder.HasCycle() {
		b.cycle = finder.Cycle()
	}
}

// HasNegativeCycle 是否存在负权重环
func (b *BellManFordSP[T]) HasNegativeCycle() bool {
	return b.cycle != nil
}

// HasPathTo 是否存在到顶点vertex的路径
func (b *BellManFordSP[T]) HasPathTo(vertex T) bool {
	return !math.IsInf(b.distTo.Get(vertex), -1)
}

// DistTo 到顶点vertex的权重
func (b *BellManFordSP[T]) DistTo(vertex T) (float64, bool) {
	if b.HasNegativeCycle() {
		return 0, false
	}
	return b.distTo.Get(vertex), true
}

// PathTo 到顶点vertex的路径
func (b *BellManFordSP[T]) PathTo(vertex T) (common.Iterable[*graph.DirectedEdge[T]], bool) {
	iter := stack.NewLIFOStack[*graph.DirectedEdge[T]]()
	if b.HasNegativeCycle() {
		return iter, false
	}
	for e := b.edgeTo.Get(vertex); e != nil; e = b.edgeTo.Get(e.From()) {
		iter.Push(e)
	}
	return iter.NewIterator(), true
}
