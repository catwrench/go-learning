package queue

import (
	"go_learning/common"
)

// Queue 队列(先进先出)
type Queue[T any] struct {
	first   *common.Node[T]
	last    *common.Node[T]
	num     int
	idxIter *common.Node[T] // 迭代索引
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// EnQueue 向队列尾部添加元素
func (q *Queue[T]) EnQueue(data T) *Queue[T] {
	temp := q.last
	q.last = common.NewNode[T](data)
	q.last.SetNext(nil)
	if q.IsEmpty() {
		q.first = q.last
	} else {
		temp.SetNext(q.last)
	}
	q.num++
	return q
}

func (q *Queue[T]) DeQueue() (res T) {
	// 队列先进先出，所以删除表头元素
	if q.IsEmpty() {
		return
	}
	res = q.first.Get()
	q.first = q.first.Next()
	q.num--
	return
}

func (q *Queue[T]) IsEmpty() bool {
	return q.num == 0
}

func (q *Queue[T]) NewIterator() common.Iterable[T] {
	q.idxIter = nil
	return q
}

func (q *Queue[T]) HasNext() bool {
	if q.num == 0 {
		return false
	}
	if q.idxIter == nil {
		return true
	}
	if q.idxIter.Next() != nil {
		return true
	}
	return false
}

func (q *Queue[T]) Next() T {
	if q.idxIter == nil {
		q.idxIter = q.first
		return q.idxIter.Get()
	}
	q.idxIter = q.idxIter.Next()
	return q.idxIter.Get()
}
