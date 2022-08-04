package tree

type Tree[T any] struct {
	root   *Node[T]
	Values []T // 遍历结果
}

type Node[T any] struct {
	Left  *Node[T]
	Right *Node[T]
	Value T
}

func NewTree[T any](str T) *Tree[T] {
	return &Tree[T]{
		root: &Node[T]{Value: str},
	}
}

func (t *Tree[T]) AppendLeft(node *Node[T], data T) {
	if node == nil {
		return
	}
	node.Left = &Node[T]{Value: data}
}

func (t *Tree[T]) AppendRight(node *Node[T], data T) {
	if node == nil {
		return
	}
	node.Right = &Node[T]{Value: data}
}

// PrevTraverse 前序遍历
func (t *Tree[T]) PrevTraverse(node *Node[T]) {
	if node == nil {
		return
	}
	t.Values = append(t.Values, node.Value)
	t.PrevTraverse(node.Left)
	t.PrevTraverse(node.Right)
}

// MiddleTraverse 中序遍历
func (t *Tree[T]) MiddleTraverse(node *Node[T]) {
	if node == nil {
		return
	}
	t.MiddleTraverse(node.Left)
	t.Values = append(t.Values, node.Value)
	t.MiddleTraverse(node.Right)
}

// NextTraverse 后序遍历
func (t *Tree[T]) NextTraverse(node *Node[T]) {
	if node == nil {
		return
	}
	t.NextTraverse(node.Left)
	t.NextTraverse(node.Right)
	t.Values = append(t.Values, node.Value)
}
