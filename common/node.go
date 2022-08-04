package common

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		value: data,
	}
}

func (n *Node[T]) Get() T {
	return n.value
}

func (n *Node[T]) Set(data T) *Node[T] {
	n.value = data
	return n
}

func (n *Node[T]) SetNext(node *Node[T]) *Node[T] {
	n.next = node
	return n
}

func (n *Node[T]) HasNext() bool {
	// 因为是栈，所以是从最后一位向前遍历的
	return n.next != nil
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}
