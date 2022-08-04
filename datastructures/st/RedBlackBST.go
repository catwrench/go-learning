package st

import (
	"go_learning/common"
	"go_learning/datastructures/queue"
)

type Color bool // 红黑树颜色

const (
	Red   Color = true // 默认创建的节点为红色
	Black Color = false
)

// RedBlackBST 符号表（基于二叉查找树）
type RedBlackBST[K common.IntAny, V any] struct {
	queue *queue.Queue[K]
	root  *RBNode[K, V]
}

func NewRedBlackBST[K common.IntAny, V any]() *RedBlackBST[K, V] {
	return &RedBlackBST[K, V]{}
}

// RBNode 红黑树节点
type RBNode[K common.IntAny, V any] struct {
	key         K
	value       V
	left, right *RBNode[K, V] // 左右子树
	N           int           // 以该节点为root的树中元素的个数
	Color       Color
}

func NewRBNode[K common.IntAny, V any](k K, v V) *RBNode[K, V] {
	return &RBNode[K, V]{
		key:   k,
		value: v,
		N:     1,
	}
}

// ---------------------------------------------------------------

func (b *RedBlackBST[K, V]) Get(k K) (res V) {
	node := b.get(b.root, k)
	if node != nil {
		return node.value
	}

	return res
}

func (b *RedBlackBST[K, V]) Contains(k K) bool {
	return b.contains(b.root, k)
}

func (b *RedBlackBST[K, V]) IsEmpty() bool {
	return b.size(b.root) == 0
}

func (b *RedBlackBST[K, V]) Size() int {
	return b.size(b.root)
}

// Rank 返回小于 k 的元素个数
func (b *RedBlackBST[K, V]) Rank(k K) int {
	return b.rank(b.root, k)
}

// Select 返回排名为rank的键
func (b *RedBlackBST[K, V]) Select(rank int) (res K) {
	return b.selectR(b.root, rank)
}

func (b *RedBlackBST[K, V]) Min() (res K) {
	node := b.min(b.root)
	if node != nil {
		return node.key
	}
	return
}

func (b *RedBlackBST[K, V]) Max() (res K) {
	node := b.max(b.root)
	if node != nil {
		return node.key
	}
	return
}

func (b *RedBlackBST[K, V]) Keys() {
	b.queue = queue.NewQueue[K]()
	b.keys(b.root, b.Min(), b.Max())
}

// NewIterator 返回一个key的迭代器
func (b *RedBlackBST[K, V]) NewIterator() *queue.Queue[K] {
	b.Keys()
	return b.queue
}

// ----------------红黑树 相比 二叉搜索树 有变化的逻辑-----------------

func (b *RedBlackBST[K, V]) Put(k K, v V) {
	b.root = b.put(b.root, k, v)
	b.root.Color = Black
	return
}

func (b *RedBlackBST[K, V]) Del(k K) {
	if b.root == nil {
		return
	}
	// 左右均为黑色节点时，将根节点置为红色
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.Color = Red
	}
	b.root = b.del(b.root, k)
	if b.root != nil {
		b.root.Color = Black
	}
}

func (b *RedBlackBST[K, V]) DeleteMin() {
	if b.root == nil {
		return
	}
	// 左右均为黑色节点时，将根节点置为红色
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.Color = Red
	}
	b.root = b.deleteMin(b.root)
	if b.root != nil {
		b.root.Color = Black
	}
	return
}

func (b *RedBlackBST[K, V]) DeleteMax() {
	if b.root == nil {
		return
	}
	// 左右均为黑色节点时，将根节点置为红色
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.Color = Red
	}
	b.root = b.deleteMax(b.root)
	if b.root != nil {
		b.root.Color = Black
	}
}

func (b *RedBlackBST[K, V]) put(node *RBNode[K, V], k K, v V) *RBNode[K, V] {
	if node == nil {
		return NewRBNode[K, V](k, v).SetN(1).SetColor(Red)
	}
	// k 比当前节点小就加在左子节点，比当前节点大就加载又子节点，等于当前节点就替换value
	if k < node.key {
		node.left = b.put(node.left, k, v)
	} else if k > node.key {
		node.right = b.put(node.right, k, v)
	} else {
		node.value = v
	}

	// 对节点进行平衡（旋转+变色）
	node = b.balance(node)
	return node
}

func (b *RedBlackBST[K, V]) del(node *RBNode[K, V], k K) *RBNode[K, V] {
	if node == nil {
		return nil
	}

	if k < node.key {
		if !b.isRed(node.left) && !b.isRed(node.left.left) {
			node = b.moveRedLeft(node)
		}
		node.left = b.del(node.left, k)
	} else {
		if b.isRed(node.left) {
			node = b.rotateRight(node)
		}
		if k == node.key && node.right == nil {
			return nil
		}
		if !b.isRed(node.right) && !b.isRed(node.right.left) {
			node = b.moveRedRight(node)
		}
		if k == node.key {
			node.value = b.get(node.right, b.min(node.right).key).value
			node.key = b.min(node.right).key
			node.right = b.deleteMin(node.right)
		} else {
			node.right = b.del(node.right, k)
		}
	}
	return b.balance(node)
}

func (b *RedBlackBST[K, V]) deleteMin(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node.left
	}
	// 如果左右节点都是黑色，那么将红色节点左移
	if !b.isRed(node.left) && !b.isRed(node.left.left) {
		node = b.moveRedLeft(node)
	}

	node.left = b.deleteMin(node.left)
	return b.balance(node)
}

func (b *RedBlackBST[K, V]) deleteMax(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	if b.isRed(node.left) {
		node = b.rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	// 如果右节点是黑色，右节点的左孩子是黑色，那么将红色节点右移
	if !b.isRed(node.right) && !b.isRed(node.right.left) {
		node = b.moveRedRight(node)
	}

	node.right = b.deleteMax(node.right)
	return b.balance(node)
}

// moveRedLeft 将红色节点左移
func (b *RedBlackBST[K, V]) moveRedLeft(node *RBNode[K, V]) *RBNode[K, V] {
	// 假设 node 为红色节点，node.left 和 node.left.left 均为黑色节点
	// 将 node.left 或者 node.left 其中一个孩子节点变红
	b.flipColors(node)
	if node.right != nil && b.isRed(node.right.left) {
		node.right = b.rotateRight(node.right)
		node = b.rotateLeft(node)
	}
	return node
}

// moveRedRight 将红色节点右移
func (b *RedBlackBST[K, V]) moveRedRight(node *RBNode[K, V]) *RBNode[K, V] {
	// 假设 node 为红色节点，node.right 和 node.right.left 均为黑色节点
	// 将 node.right 或者 node.right 其中一个孩子节点变红
	b.flipColors(node)
	if node.left != nil && b.isRed(node.left.left) {
		node = b.rotateRight(node)
	}
	return node
}

// 对节点进行平衡（旋转+变色）
func (b *RedBlackBST[K, V]) balance(node *RBNode[K, V]) *RBNode[K, V] {
	// 进行旋转和染色
	// 1、如果左孩子为黑色，右孩子为红色：进行左旋
	// 2、如果左孩子为红色，左孩子的孩子也为红色：进行右旋
	// 3、如果左右孩子都为红色：进行 染色
	// ps: 1和2执行完后最后都会转换为3的情况，所以不用重复执行染色
	if !b.isRed(node.left) && b.isRed(node.right) {
		node = b.rotateLeft(node)
	}
	if b.isRed(node.left) && b.isRed(node.left.left) {
		node = b.rotateRight(node)
	}
	if b.isRed(node.left) && b.isRed(node.right) {
		b.flipColors(node)
	}

	node.N = 1 + b.size(node.left) + b.size(node.right)
	return node
}

// rotateLeft 节点左旋
func (b *RedBlackBST[K, V]) rotateLeft(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	newNode := node.right
	node.right = newNode.left
	newNode.left = node

	newNode.Color = node.Color
	node.Color = Red

	newNode.N = node.N
	node.N = 1 + b.size(node.left) + b.size(node.right)
	return newNode
}

// rotateRight 节点右旋
func (b *RedBlackBST[K, V]) rotateRight(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	newNode := node.left
	node.left = newNode.right
	newNode.right = node

	newNode.Color = node.Color
	node.Color = Red

	newNode.N = node.N
	node.N = 1 + b.size(node.left) + b.size(node.right)
	return newNode
}

// flipColors 颜色翻转
func (b *RedBlackBST[K, V]) flipColors(node *RBNode[K, V]) {
	if node == nil {
		return
	}
	node.Color = Red
	if node.left != nil {
		node.left.Color = Black
	}
	if node.right != nil {
		node.right.Color = Black
	}
}

func (b *RedBlackBST[K, V]) isRed(node *RBNode[K, V]) bool {
	if node == nil {
		return false
	}
	return node.Color == Red
}

// ---------------------------------------------------------------

func (b *RedBlackBST[K, V]) get(node *RBNode[K, V], k K) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	// 比k小查找做子树，比k大查找又子树，等于k就返回对应的node
	if k < node.key {
		return b.get(node.left, k)
	} else if k > node.key {
		return b.get(node.right, k)
	} else {
		return node
	}
}

func (b *RedBlackBST[K, V]) contains(node *RBNode[K, V], k K) bool {
	if node == nil {
		return false
	}
	// 比k小查找做子树，比k大查找又子树，等于k就返回对应的v
	if k < node.key {
		return b.contains(node.left, k)
	}
	if k > node.key {
		return b.contains(node.right, k)
	}
	return true
}

// size 返回以node为根节点的树 的元素个数
func (b *RedBlackBST[K, V]) size(node *RBNode[K, V]) int {
	if node == nil {
		return 0
	}
	return 1 + b.size(node.left) + b.size(node.right)
}

func (b *RedBlackBST[K, V]) rank(node *RBNode[K, V], k K) (res int) {
	if node == nil {
		return 0
	}
	// k比当前节点小，递归计算左子树，左子树里可能存在比k大的元素
	// k比当前节点大，直接计算当前节点左子树的size + 递归计算右子树
	// k等于当前节点，直接计算当前节点左子树的size即可
	if k < node.key {
		return b.rank(node.left, k)
	} else if k > node.key {
		return 1 + b.size(node.left) + b.rank(node.right, k)
	} else {
		return b.size(node.left)
	}
}

func (b *RedBlackBST[K, V]) selectR(node *RBNode[K, V], rank int) (res K) {
	if node == nil {
		return
	}
	curRank := b.size(node.left)
	if curRank > rank {
		return b.selectR(node.left, rank)
	} else if curRank < rank {
		return b.selectR(node.right, rank-curRank-1)
	} else {
		return node.key
	}
}

func (b *RedBlackBST[K, V]) min(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return b.min(node.left)
	}
	return node
}

func (b *RedBlackBST[K, V]) max(node *RBNode[K, V]) *RBNode[K, V] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return b.max(node.right)
	}
	return node
}

func (b *RedBlackBST[K, V]) keys(node *RBNode[K, V], low, high K) {
	if node == nil {
		return
	}
	// 二叉查找树的范围查找
	if low < node.key {
		b.keys(node.left, low, high)
	}
	if low <= node.key && node.key <= high {
		b.queue.EnQueue(node.key)
	}
	if high > node.key {
		b.keys(node.right, low, high)
	}
}

// ---------------------------------------------------------------

func (n *RBNode[K, V]) SetN(N int) *RBNode[K, V] {
	n.N = N
	return n
}

func (n *RBNode[K, V]) SetColor(color Color) *RBNode[K, V] {
	n.Color = color
	return n
}
