package DegreesOfSeparation

import (
	"go_learning/algorithm/graph/bfs/BfsPath"
	"go_learning/datastructures/graph"
)

// DegreesOfSeparation 分隔的度
// 通过符号图建立的索引图，进行广度优先查找，找到两个顶点之间最短的路径
type DegreesOfSeparation[T comparable] struct {
	symbolGraph *graph.SymbolGraph[T] // 符号图
}

func NewDegreesOfSeparation[T comparable]() *DegreesOfSeparation[T] {
	return &DegreesOfSeparation[T]{
		symbolGraph: graph.NewSymbolGraph[T](),
	}
}

func (d *DegreesOfSeparation[T]) AddVertex(vertex T) {
	d.symbolGraph.AddVertex(vertex)
}

func (d *DegreesOfSeparation[T]) AddEdge(vertex1, vertex2 T) {
	d.symbolGraph.AddEdge(vertex1, vertex2)
}

// Degree 通过路径计算度
func (d *DegreesOfSeparation[T]) Degree(source, sink T) (path []T, ok bool) {
	path = make([]T, 0)
	if !d.symbolGraph.Contains(source) || !d.symbolGraph.Contains(sink) {
		return path, false
	}

	// 通过符号图建立的索引图，进行广度优先查找，找到两个顶点之间最短的路径
	g := d.symbolGraph.Graph()
	bfs := BfsPath.NewBreadthFirstPaths[int](*g)
	idxSource := d.symbolGraph.Index(source)
	idxSink := d.symbolGraph.Index(sink)

	bfs.Search(*g, idxSource)
	if !bfs.HasPathTo(idxSink) {
		return path, false
	}
	it := bfs.PathTo(idxSink)
	for it.HasNext() {
		path = append(path, d.symbolGraph.Name(it.Next()))
	}
	return path, true
}
