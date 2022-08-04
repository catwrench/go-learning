package st

import "go_learning/common"

// BinarySearchST 二分查找符号表
type BinarySearchST[K common.IntAny, V any] struct {
	keys    []K
	values  []V
	n       int // 当前元素数量
	idxIter int // 迭代索引
}

func NewBinarySearchST[K common.IntAny, V any](cap int) *BinarySearchST[K, V] {
	return &BinarySearchST[K, V]{
		keys:    make([]K, cap, cap),
		values:  make([]V, cap, cap),
		idxIter: -1,
	}
}

func (b *BinarySearchST[K, V]) Put(k K, v V) {
	i := b.Rank(k)
	// 命中，更新
	if i < b.n && b.keys[i] == k {
		b.values[i] = v
		return
	}
	// 未命中，添加
	for j := b.n; j > i; j-- {
		b.keys[j] = b.keys[j-1]
		b.values[j] = b.values[j-1]
	}
	b.keys[i] = k
	b.values[i] = v
	b.n++
}

func (b *BinarySearchST[K, V]) Get(k K) (res V) {
	i := b.Rank(k)
	if i < b.n && b.keys[i] == k {
		return b.values[i]
	}
	return
}

func (b *BinarySearchST[K, V]) Del(k K) {
	i := b.Rank(k)
	if b.n == 0 {
		return
	}
	// 不在末尾，索引后面的元素前移
	if i < b.n-1 && b.keys[i] == k {
		for j := i; j < b.n-1; j++ {
			b.keys[j] = b.keys[j+1]
			b.values[j] = b.values[j+1]
		}
	}
	b.n--
}

func (b *BinarySearchST[K, V]) Contains(k K) bool {
	i := b.Rank(k)
	if i < b.n && b.keys[i] == k {
		return true
	} else {
		return false
	}
}

func (b *BinarySearchST[K, V]) IsEmpty() bool {
	return b.n == 0
}

func (b *BinarySearchST[K, V]) Size() int {
	return b.n
}

// Rank 小于k的键的数量
func (b *BinarySearchST[K, V]) Rank(k K) int {
	// 通过二分查找找到键所在索引位置
	l, r := 0, b.n-1

	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if b.keys[mid] > k {
			r = mid - 1
		} else if b.keys[mid] < k {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}

func (b *BinarySearchST[K, V]) Keys() []K {
	return b.keys
}

func (b *BinarySearchST[K, V]) Values() []V {
	return b.values
}

func (b *BinarySearchST[K, V]) Min() (res V) {
	if b.n > 0 {
		return b.values[0]
	}
	return
}

func (b *BinarySearchST[K, V]) Max() (res V) {
	if b.n > 0 {
		return b.values[b.n-1]
	}
	return
}

func (b *BinarySearchST[K, V]) NewIterator() *BinarySearchST[K, V] {
	b.idxIter = -1
	return b
}

func (b *BinarySearchST[K, V]) HasNext() bool {
	return b.idxIter < b.n-1 && b.n > 0
}

func (b *BinarySearchST[K, V]) Next() (k K, v V) {
	if b.idxIter < b.n-1 {
		b.idxIter++
		return b.keys[b.idxIter], b.values[b.idxIter]
	}
	return
}
