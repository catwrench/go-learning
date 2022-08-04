package queue

import "go_learning/common"

type IndexMinPQ[T common.IntAny] struct {
	queue     []*T  // 数组实现的堆,0号位不使用,存放有优先级之分的元素
	idxQueue  []int // 索引队列
	idxQueueR []int // 反向索引队列, idxQueueR[i] 等于 i 在 idxQueue 的位置，即 idxQueue[j] = i
	N         int   // 元素个数
}

func NewIndexMinPQ[T common.IntAny](maxN int) *IndexMinPQ[T] {
	res := &IndexMinPQ[T]{}
	res.queue = make([]*T, maxN+1)
	res.idxQueue = make([]int, maxN+1)
	res.idxQueueR = make([]int, maxN+1)
	for i := 0; i <= maxN; i++ {
		res.idxQueueR[i] = -1
	}
	return res
}

func (i *IndexMinPQ[T]) Insert(idx int, item *T) {
	i.N++
	i.queue[idx] = item
	i.idxQueue[i.N] = idx
	i.idxQueueR[idx] = i.N
	i.swim(i.N)
}

func (i *IndexMinPQ[T]) Change(idx int, item *T) {
	i.queue[idx] = item
	// 正向索引先上浮，后下沉
	i.swim(i.idxQueueR[idx])
	i.sink(i.idxQueueR[idx])
}

func (i *IndexMinPQ[T]) Contain(idx int) bool {
	return i.idxQueueR[idx] != -1
}

func (i *IndexMinPQ[T]) Delete(idx int) {
	// 先将元素和队尾交换，然后对idx位置的新元素进行上浮和下沉，标记队尾元素删除
	i.Exchange(i.N, i.idxQueueR[idx])
	i.ExchangeR(idx, i.idxQueue[i.N])
	i.N--
	i.swim(i.idxQueueR[idx])
	i.sink(i.idxQueueR[idx])
	i.queue[i.idxQueue[i.N+1]] = nil
	i.idxQueueR[i.idxQueue[i.N+1]] = -1
}

func (i *IndexMinPQ[T]) Min() *T {
	return i.queue[i.idxQueue[1]]
}

func (i *IndexMinPQ[T]) MinIdx() int {
	return i.idxQueue[1]
}

func (i *IndexMinPQ[T]) DelMin() int {
	// 交换1位置的最小元素到末尾，删除末尾元素，然后对头部最大元素执行下沉
	// 和普通队列的区别是,需要先通过反向索引队列 idxQueueR 找到最小值对应的 idx
	idxMin := i.idxQueue[1]
	idxMax := i.idxQueue[i.N]
	i.Exchange(1, i.N)
	i.ExchangeR(idxMin, idxMax)
	i.N--
	i.sink(1)
	i.queue[i.idxQueue[i.N+1]] = nil    // 移除和元素的关联
	i.idxQueueR[i.idxQueue[i.N+1]] = -1 // 标识对应索引不存在元素
	return idxMin
}

func (i *IndexMinPQ[T]) IsEmpty() bool {
	return i.N == 0
}

func (i *IndexMinPQ[T]) Size() int {
	return i.N
}

func (i *IndexMinPQ[T]) Less(firstIdx, secondIdx int) bool {
	if i.queue[i.idxQueue[firstIdx]] != nil && i.queue[i.idxQueue[secondIdx]] != nil {
		return *i.queue[i.idxQueue[firstIdx]] < *i.queue[i.idxQueue[secondIdx]]
	}
	return false
}

func (i *IndexMinPQ[T]) Exchange(firstIdx, secondIdx int) {
	i.idxQueue[firstIdx], i.idxQueue[secondIdx] = i.idxQueue[secondIdx], i.idxQueue[firstIdx]
}

func (i *IndexMinPQ[T]) ExchangeR(firstIdx, secondIdx int) {
	i.idxQueueR[firstIdx], i.idxQueueR[secondIdx] = i.idxQueueR[secondIdx], i.idxQueueR[firstIdx]
}

// Swim 元素上浮
func (i *IndexMinPQ[T]) swim(index int) {
	for index > 1 && i.Less(index, index/2) {
		i.Exchange(index, index/2)
		index /= 2
	}
}

// Sink 元素下沉
func (i *IndexMinPQ[T]) sink(index int) {
	// 下沉需要两次比较，一次用来找出较小的子节点，一次用来确定节点是否需要下沉

	for index*2 <= i.N {
		doubleIdx := index * 2
		// 找出较小的子节点，在堆底部只有一个子节点时不用比较
		if doubleIdx < i.N && i.Less(doubleIdx+1, doubleIdx) {
			doubleIdx += 1
		}

		// 比较，确定节点是否需要下沉
		if i.Less(doubleIdx, index) {
			i.Exchange(doubleIdx, index)
			index = doubleIdx
		} else {
			break
		}
	}
}
