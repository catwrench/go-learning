package others

import (
	"go_learning/common"
	"go_learning/datastructures/st"
)

// SymbolIndexMinPQ 符号的索引最小优先队列
// K 顶点，V权重值
type SymbolIndexMinPQ[K comparable, V common.IntAny] struct {
	queue     []V                            // 数组实现的堆,0号位不使用,存放有优先级之分的元素
	keyIdxST  *st.SequentialSearchST[K, int] // k 到int索引的符号表
	idxKeyST  []K                            // int索引到 k的关联
	idxQueue  []int                          // 正向索引,索引二叉堆 1->int索引
	idxQueueR []int                          // 反向索引 int索引->1
	N         int                            // 元素个数
}

func NewSymbolIndexMinPQ[K comparable, V common.IntAny](maxN int) *SymbolIndexMinPQ[K, V] {
	res := &SymbolIndexMinPQ[K, V]{
		queue:     make([]V, maxN+1),
		keyIdxST:  st.NewSequentialSearchST[K, int](),
		idxKeyST:  make([]K, maxN+1),
		idxQueue:  make([]int, maxN+1),
		idxQueueR: make([]int, maxN+1),
	}
	for i := 0; i <= maxN; i++ {
		res.idxQueueR[i] = -1
	}
	return res
}

func (i *SymbolIndexMinPQ[K, V]) Insert(key K, value V) {
	i.N++
	var idx int
	idx = i.keyIdxST.Size() + 1
	i.keyIdxST.Put(key, idx)
	i.idxKeyST[idx] = key
	i.queue[idx] = value
	i.idxQueue[i.N] = idx
	i.idxQueueR[idx] = i.N
	i.swim(i.N)
}

func (i *SymbolIndexMinPQ[K, V]) Change(key K, value V) {
	idx := i.getIdx(key)
	i.queue[idx] = value
	// 正向索引先上浮，后下沉
	i.swim(i.idxQueueR[idx])
	i.sink(i.idxQueueR[idx])
}

func (i *SymbolIndexMinPQ[K, V]) Contain(key K) bool {
	return i.idxQueueR[i.getIdx(key)] != -1
}

func (i *SymbolIndexMinPQ[K, V]) Delete(key K) {
	// 先将元素和队尾交换，然后对idx位置的新元素进行上浮和下沉，标记队尾元素删除
	idx := i.getIdx(key)
	i.exchange(idx, i.N)
	i.exchangeR(idx, i.idxQueue[i.N])
	i.N--
	i.swim(i.idxQueueR[idx])
	i.sink(i.idxQueueR[idx])
	i.queue[i.idxQueue[i.N+1]] = 0
	i.idxQueueR[i.idxQueue[i.N+1]] = -1
}

func (i *SymbolIndexMinPQ[K, V]) Min() V {
	return i.queue[i.idxQueue[1]]
}

func (i *SymbolIndexMinPQ[K, V]) MinIdx() K {
	return i.idxKeyST[i.idxQueue[1]]
}

func (i *SymbolIndexMinPQ[K, V]) DelMin() K {
	// 交换1位置的最小元素到末尾，删除末尾元素，然后对头部最大元素执行下沉
	// 和普通队列的区别是,需要先通过反向索引队列 idxQueueR 找到最小值对应的 idx
	// 删除时需要处理对应的 idxKeyST
	idxMin := i.idxQueue[1]
	idxMax := i.idxQueue[i.N]
	i.exchange(1, i.N)
	i.exchangeR(idxMin, idxMax)
	i.N--
	i.sink(1)
	i.queue[i.idxQueue[i.N+1]] = 0      // 移除和元素的关联
	i.idxQueueR[i.idxQueue[i.N+1]] = -1 // 标识对应索引不存在元素
	return i.idxKeyST[idxMin]
}

func (i *SymbolIndexMinPQ[K, V]) IsEmpty() bool {
	return i.N == 0
}

func (i *SymbolIndexMinPQ[K, V]) Size() int {
	return i.N
}

func (i *SymbolIndexMinPQ[K, V]) less(firstIdx, secondIdx int) bool {
	if i.queue[i.idxQueue[firstIdx]] != 0 && i.queue[i.idxQueue[secondIdx]] != 0 {
		return i.queue[i.idxQueue[firstIdx]] < i.queue[i.idxQueue[secondIdx]]
	}
	return false
}

func (i *SymbolIndexMinPQ[K, V]) exchange(firstIdx, secondIdx int) {
	// i.queue[firstIdx], i.queue[secondIdx] = i.queue[secondIdx], i.queue[firstIdx]
	i.idxQueue[firstIdx], i.idxQueue[secondIdx] = i.idxQueue[secondIdx], i.idxQueue[firstIdx]
}

func (i *SymbolIndexMinPQ[K, V]) exchangeR(firstIdx, secondIdx int) {
	i.idxQueueR[firstIdx], i.idxQueueR[secondIdx] = i.idxQueueR[secondIdx], i.idxQueueR[firstIdx]
}

// Swim 元素上浮
func (i *SymbolIndexMinPQ[K, V]) swim(index int) {
	for index > 1 && i.less(index, index/2) {
		i.exchange(index, index/2)
		index /= 2
	}
}

// Sink 元素下沉
func (i *SymbolIndexMinPQ[K, V]) sink(index int) {
	// 下沉需要两次比较，一次用来找出较小的子节点，一次用来确定节点是否需要下沉

	for index*2 <= i.N {
		doubleIdx := index * 2
		// 找出较小的子节点，在堆底部只有一个子节点时不用比较
		if doubleIdx < i.N && i.less(doubleIdx+1, doubleIdx) {
			doubleIdx += 1
		}

		// 比较，确定节点是否需要下沉
		if i.less(doubleIdx, index) {
			i.exchange(doubleIdx, index)
			index = doubleIdx
		} else {
			break
		}
	}
}

func (i *SymbolIndexMinPQ[K, V]) getIdx(key K) int {
	return i.keyIdxST.Get(key)
}
