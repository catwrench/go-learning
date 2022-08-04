package graph

import "fmt"

// DirectedEdge 带权重的有向边
type DirectedEdge[T comparable] struct {
	vertex1, vertex2 T       // 边的两个顶点 v1 -> v2
	weight           float64 // 权重
}

func NewDirectedEdge[T comparable](from, to T, weight float64) *DirectedEdge[T] {
	return &DirectedEdge[T]{
		vertex1: from,
		vertex2: to,
		weight:  weight,
	}
}

// From 返回边的起始顶点
func (e *DirectedEdge[T]) From() (res T) {
	return e.vertex1
}

// To 随机返回边的结束顶点
func (e *DirectedEdge[T]) To() T {
	return e.vertex2
}

// Weight 获取边的权重
func (e *DirectedEdge[T]) Weight() float64 {
	return e.weight
}

// CompareTo 和另一条边比较权重大小
// 大于：1，小于：-1，等于：0
func (e *DirectedEdge[T]) CompareTo(edge DirectedEdge[T]) int {
	if e.Weight() > edge.Weight() {
		return 1
	}
	if e.Weight() < edge.Weight() {
		return -1
	}
	return 0
}

func (e *DirectedEdge[T]) ToString() string {
	return fmt.Sprintf("%v-%v %f", e.vertex1, e.vertex2, e.weight)
}
