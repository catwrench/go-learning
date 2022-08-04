package graph

import "fmt"

// Edge 带权重的边
type Edge[T comparable] struct {
	vertex1, vertex2 T       // 边的两个顶点
	weight           float64 // 权重
}

func NewEdge[T comparable](vertex1, vertex2 T, weight float64) *Edge[T] {
	return &Edge[T]{
		vertex1: vertex1,
		vertex2: vertex2,
		weight:  weight,
	}
}

// Either 随机返回边的一个顶点
func (e *Edge[T]) Either() T {
	return e.vertex1
}

// Other 返回边的另一个顶点
func (e *Edge[T]) Other(vertex T) (res T, ok bool) {
	if vertex == e.vertex1 {
		return e.vertex2, true
	}
	if vertex == e.vertex2 {
		return e.vertex1, true
	}
	return res, false
}

// Weight 获取边的权重
func (e *Edge[T]) Weight() float64 {
	return e.weight
}

// CompareTo 和另一条边比较权重大小
// 大于：1，小于：-1，等于：0
func (e *Edge[T]) CompareTo(edge Edge[T]) int {
	if e.Weight() > edge.Weight() {
		return 1
	}
	if e.Weight() < edge.Weight() {
		return -1
	}
	return 0
}

func (e *Edge[T]) ToString() string {
	return fmt.Sprintf("%v-%v %f", e.vertex1, e.vertex2, e.weight)
}
