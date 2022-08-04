package dfs

import (
	"go_learning/common"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// DirectedDFS 有向图的深度优先搜索
type DirectedDFS[T comparable] struct {
	marked *st.SeparateChainingHashST[T, bool] // 用于记录访问过的点（hash符号表-拉链法）
	count  int                                 // 访问过的顶点数
}

func NewDirectedDFS[T comparable](G graph.DiGraph[T]) *DirectedDFS[T] {
	// 用图的所有顶点，初始化对应的访问点
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	return &DirectedDFS[T]{
		marked: marked,
	}
}

func NewDirectedDFSWith[T comparable](G graph.DiGraph[T], source common.Iterable[T]) *DirectedDFS[T] {
	VertexArr := G.V()
	marked := st.NewSeparateChainingHashST[T, bool](len(VertexArr))
	for i := range VertexArr {
		marked.Put(VertexArr[i], false)
	}

	for source.HasNext() {
		v := source.Next()
		if !marked.Get(v) {

		}
	}
	return &DirectedDFS[T]{
		marked: marked,
	}
}

// public DirectedDFS(Digraph G, Iterable<Integer> sources) {
// marked = new boolean[G.V()];
// validateVertices(sources);
// for (int v : sources) {
// if (!marked[v]) dfs(G, v);
// }
// }

// Dfs 从顶点开始搜索
func (d *DirectedDFS[T]) Dfs(G graph.DiGraph[T], startVertex T) {
	d.Mark(startVertex)
	it := G.Adj(startVertex)
	for it.HasNext() {
		nextVertex := it.Next()
		// 未被标记就继续搜索下一个点，否则回退
		if !d.marked.Get(nextVertex) {
			d.Dfs(G, nextVertex)
		}
	}
}

// Mark 标记顶点为已访问
func (d *DirectedDFS[T]) Mark(vertex T) {
	if d.marked.Get(vertex) {
		return
	}
	d.marked.Put(vertex, true)
	d.count++
}

// Marked 顶点是否访问过
func (d *DirectedDFS[T]) Marked(vertex T) bool {
	return d.marked.Get(vertex)
}

// Count 访问过的顶点数
func (d *DirectedDFS[T]) Count() int {
	return d.count
}

// Reachable 多点可到达
func (d *DirectedDFS[T]) Reachable(g graph.DiGraph[T], sources ...T) (res []T) {
	res = make([]T, 0)
	for _, item := range sources {
		d.Dfs(g, item)
	}
	vertexes := g.V()
	for i := range vertexes {
		if d.Marked(vertexes[i]) {
			res = append(res, vertexes[i])
		}
	}
	return res
}
