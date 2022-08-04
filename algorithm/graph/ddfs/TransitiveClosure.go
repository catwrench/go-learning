package dfs

import (
	"go_learning/datastructures/graph"
	"go_learning/datastructures/st"
)

// TransitiveClosure 传递闭包
// 需要V^2的空间来做图的预处理，V个 DirectedDFS 对象都需要 marked 符号表来检测E条边
type TransitiveClosure[T comparable] struct {
	all *st.SeparateChainingHashST[T, *DirectedDFS[T]]
}

func NewTransitiveClosure[T comparable](G graph.DiGraph[T]) *TransitiveClosure[T] {
	V := G.V()
	res := &TransitiveClosure[T]{
		all: st.NewSeparateChainingHashST[T, *DirectedDFS[T]](len(V)),
	}

	// 完成每个图的预处理
	for _, vertex := range V {
		ddfs := NewDirectedDFS[T](G)
		ddfs.Dfs(G, vertex)
		res.all.Put(vertex, ddfs)
	}
	return res
}

func (t *TransitiveClosure[T]) Reachable(vertex1, vertex2 T) bool {
	return t.all.Get(vertex1).Marked(vertex2)
}
