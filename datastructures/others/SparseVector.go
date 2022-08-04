package others

import "go_learning/datastructures/st"

// SparseVector 稀疏向量
type SparseVector struct {
	st *st.SeparateChainingHashST[int, float64]
}

func NewSparseVector(cap int) *SparseVector {
	return &SparseVector{
		st: st.NewSeparateChainingHashST[int, float64](cap),
	}
}

// Dot 点乘
func (s *SparseVector) Dot(that []float64) (sum float64) {
	for _, i := range s.st.Keys() {
		sum += that[i] * s.st.Get(i)
	}
	return sum
}

func (s *SparseVector) Get(i int) float64 {
	return s.st.Get(i)
}

func (s *SparseVector) Put(i int, v float64) {
	s.st.Put(i, v)
}

func (s *SparseVector) Size() int {
	return s.st.Size()
}
