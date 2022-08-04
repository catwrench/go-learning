package bag

import (
	"go_learning/common"
)

type Bag[T any] struct {
	first   *common.Node[T]
	idxIter *common.Node[T]
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{}
}

func (b *Bag[T]) Add(data T) *Bag[T] {
	if b.first == nil {
		b.first = common.NewNode[T](data)
		return b
	}

	temp := common.NewNode[T](b.first.Get())
	temp.SetNext(b.first.Next())
	b.first.Set(data)
	b.first.SetNext(temp)
	return b
}

func (b *Bag[T]) Del() *Bag[T] {
	if b.IsEmpty() {
		return b
	}
	b.first = b.first.Next()
	return b
}

func (b *Bag[T]) IsEmpty() bool {
	return b.first == nil
}

// NewIterator 当前背包的实现方式，迭代时为添加的逆序，并非无序
func (b *Bag[T]) NewIterator() common.Iterable[T] {
	b.idxIter = nil
	return b
}

func (b *Bag[T]) HasNext() bool {
	if b.idxIter == nil {
		return b.first != nil
	}
	return b.idxIter.HasNext()
}

func (b *Bag[T]) Next() (res T) {
	if b.idxIter == nil {
		b.idxIter = b.first
		return b.idxIter.Get()
	}
	b.idxIter = b.idxIter.Next()
	return b.idxIter.Get()
}
