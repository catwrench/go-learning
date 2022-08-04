package graph

import "go_learning/datastructures/st"

// SymbolGraph 符号图
type SymbolGraph[T comparable] struct {
	keySt *st.SequentialSearchST[T, int] // 正向索引：符号名 -> idx 的符号表
	keys  []T                            // 反向索引：idx索引 -> 符号名
	graph *Graph[int]                    // 图，只存储索引
}

func NewSymbolGraph[T comparable]() *SymbolGraph[T] {
	res := &SymbolGraph[T]{
		keySt: st.NewSequentialSearchST[T, int](),
		keys:  make([]T, 0),
		graph: NewGraph[int](),
	}
	return res
}

// AddVertex 添加顶点
func (s *SymbolGraph[T]) AddVertex(vertex T) {
	if !s.keySt.Contains(vertex) {
		idxKey := s.keySt.Size()
		// 添加正向索引
		s.keySt.Put(vertex, idxKey)
		// 添加反向索引
		s.keys = append(s.keys, vertex)
		// 添加图顶点
		s.graph.AddVertex(idxKey)
	}
}

// AddEdge 添加边
func (s *SymbolGraph[T]) AddEdge(vertex1, vertex2 T) {
	s.graph.AddEdge(s.keySt.Get(vertex1), s.keySt.Get(vertex2))
}

// Index 获取顶点索引
func (s *SymbolGraph[T]) Index(vertex T) int {
	return s.keySt.Get(vertex)
}

// Name 获取顶点名称
func (s *SymbolGraph[T]) Name(idxKey int) T {
	return s.keys[idxKey]
}

// Contains 顶点是否存在
func (s *SymbolGraph[T]) Contains(vertex T) bool {
	return s.keySt.Contains(vertex)
}

// Graph 获取索引图
func (s *SymbolGraph[T]) Graph() *Graph[int] {
	return s.graph
}
