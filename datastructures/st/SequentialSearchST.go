package st

import (
	"go_learning/common"
)

// SequentialSearchST 顺序查找符号表（基于链表）
type SequentialSearchST[K comparable, V any] struct {
	first   *common.KVNode[K, V]
	size    int
	idxIter *common.KVNode[K, V] // 迭代索引
}

func NewSequentialSearchST[K comparable, V any]() *SequentialSearchST[K, V] {
	return &SequentialSearchST[K, V]{}
}

func (s *SequentialSearchST[K, V]) Put(k K, v V) {
	node := s.first
	for ; node != nil; node = node.Next() {
		//  命中key
		if node.GetKey() == k {
			node.SetValue(v)
			return
		}
	}
	// 未命中key,在头部添加新节点
	s.first = common.NewKVNode[K, V](k, v).AddNext(s.first)
	s.size++
}

func (s *SequentialSearchST[K, V]) Get(k K) (res V) {
	node := s.first
	for ; node != nil; node = node.Next() {
		if node.GetKey() == k {
			return node.GetValue()
		}
	}
	return
}

func (s *SequentialSearchST[K, V]) Del(k K) {
	var pNode *common.KVNode[K, V]
	node := s.first
	for ; node != nil; node = node.Next() {
		if node.GetKey() == k {
			if pNode == nil {
				s.first = nil
			} else {
				pNode.AddNext(node.Next())
			}
			s.size--
		}
		pNode = node
	}
}

func (s *SequentialSearchST[K, V]) Contains(k K) bool {
	node := s.first
	for ; node != nil; node = node.Next() {
		if node.GetKey() == k {
			return true
		}
	}
	return false
}

func (s *SequentialSearchST[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SequentialSearchST[K, V]) Size() int {
	return s.size
}

func (s *SequentialSearchST[K, V]) NewIterator() *SequentialSearchST[K, V] {
	s.idxIter = nil
	return s
}

func (s *SequentialSearchST[K, V]) HasNext() bool {
	if s.idxIter == nil {
		if s.first != nil {
			return true
		}
	} else {
		if s.idxIter.Next() != nil {
			return true
		}
	}
	return false
}

func (s *SequentialSearchST[K, V]) Next() (k K, v V) {
	if s.idxIter == nil {
		if s.first != nil {
			s.idxIter = s.first
		}
	} else {
		s.idxIter = s.idxIter.Next()
	}
	if s.idxIter != nil {
		return s.idxIter.GetKey(), s.idxIter.GetValue()
	}
	return
}

func (s *SequentialSearchST[K, V]) Keys() []K {
	keys := make([]K, 0)
	s.NewIterator()
	for s.HasNext() {
		k, _ := s.Next()
		keys = append(keys, k)
	}
	return keys
}
