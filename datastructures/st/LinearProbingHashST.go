package st

import (
	"go_learning/common"
)

// LinearProbingHashST 散列表,线性探测
type LinearProbingHashST struct {
	n      int           // 元素个数
	cap    int           // 线性探测表容量
	keys   []string      // 键数组
	values []interface{} // 值数组
}

func NewLinearProbingHashST(cap int) *LinearProbingHashST {
	return &LinearProbingHashST{
		cap:    cap,
		keys:   make([]string, cap, cap),
		values: make([]interface{}, 0, cap),
	}
}

func (s *LinearProbingHashST) Hash(str string) int {
	return common.HashInt(str) % s.cap
}

func (s *LinearProbingHashST) Get(key string) (res interface{}) {
	for idx := s.Hash(key); s.keys[idx] != ""; idx = (idx + 1) % s.cap {
		if s.keys[idx] == key {
			return s.values[idx]
		}
	}
	return
}

func (s *LinearProbingHashST) Put(key string, v interface{}) {
	// 扩容
	if s.n > s.cap/2 {
		s.resize(s.cap * 2)
	}

	idx := s.Hash(key)
	for ; s.keys[idx] != ""; idx = (idx + 1) % s.cap {
		// key命中
		if s.keys[idx] == key {
			s.values[idx] = v
			return
		}
	}

	// key未命中
	s.keys[idx] = key
	s.values[idx] = v
	s.n++
}

func (s *LinearProbingHashST) Del(key string) {
	idx := s.Hash(key)
	// 删除指定idx
	for ; s.keys[idx] != ""; idx = (idx + 1) % s.cap {
		if s.keys[idx] == key {
			s.keys[idx] = ""
			s.values[idx] = nil
		}
	}
	// idx后连续的key重新入列
	idx = (idx + 1) % s.cap
	for ; s.keys[idx] != ""; idx = (idx + 1) % s.cap {
		s.keys[idx] = ""
		s.values[idx] = nil
		s.n--
		s.Put(s.keys[idx], s.values[idx])
	}
	s.n--

	// 缩绒
	if s.n > 0 && s.n == s.cap/8 {
		s.resize(s.cap / 2)
	}
}

func (s *LinearProbingHashST) resize(cap int) {
	// 创建新的表，并将旧数据迁移
	st := NewLinearProbingHashST(cap)
	for i := 0; i < s.cap; i++ {
		if s.keys[i] != "" {
			st.Put(s.keys[i], s.values[i])
		}
	}
	// 数据替换
	s.cap = st.cap
	s.keys = st.keys
	s.values = st.values
}
