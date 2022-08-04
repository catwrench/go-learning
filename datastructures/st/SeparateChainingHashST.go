package st

import (
	"go_learning/common"
)

// SeparateChainingHashST 散列表（基于拉链法）
type SeparateChainingHashST[K comparable, V any] struct {
	n   int                         // 元素个数
	cap int                         // 链表数量
	st  []*SequentialSearchST[K, V] // 处理冲突用的顺序符号表
}

func NewSeparateChainingHashST[K comparable, V any](cap int) *SeparateChainingHashST[K, V] {
	res := &SeparateChainingHashST[K, V]{
		cap: cap,
		st:  make([]*SequentialSearchST[K, V], 0, cap),
	}
	for i := 0; i < cap; i++ {
		res.st = append(res.st, NewSequentialSearchST[K, V]())
	}
	return res
}

func (s *SeparateChainingHashST[K, V]) Hash(key K) int {
	return common.HashInt(key) % s.cap
}

func (s *SeparateChainingHashST[K, V]) Contains(key K) bool {
	return s.st[s.Hash(key)].Contains(key)
}

func (s *SeparateChainingHashST[K, V]) Get(key K) V {
	return s.st[s.Hash(key)].Get(key)
}

func (s *SeparateChainingHashST[K, V]) Put(key K, value V) {
	s.st[s.Hash(key)].Put(key, value)
	s.n++
	return
}

func (s *SeparateChainingHashST[K, V]) Del(key K) {
	hash := s.Hash(key)
	if s.st[hash].Contains(key) {
		s.st[hash].Del(key)
		s.n--
	}
}

func (s *SeparateChainingHashST[K, V]) Cap() int {
	return s.cap
}

func (s *SeparateChainingHashST[K, V]) Size() int {
	return s.n
}

func (s *SeparateChainingHashST[K, V]) Keys() []K {
	keys := make([]K, 0)
	for i := 0; i < len(s.st); i++ {
		keys = append(keys, s.st[i].Keys()...)
	}
	return keys
}
