package st

import (
	"go_learning/common"
	"go_learning/datastructures/queue"
)

// BST 符号表（基于二叉查找树）
type BST[K common.IntAny, V any] struct {
	queue *queue.Queue[K]
	root  *Node[K, V]
}

func NewBST[K common.IntAny, V any]() *BST[K, V] {
	return &BST[K, V]{}
}

func (b *BST[K, V]) Put(k K, v V) {
	b.root = b.put(b.root, k, v)
	return
}

func (b *BST[K, V]) Get(k K) (res V) {
	node := b.get(b.root, k)
	if node != nil {
		return node.value
	}
	return res
}

func (b *BST[K, V]) Del(k K) {
	b.root = b.del(b.root, k)
}

func (b *BST[K, V]) Contains(k K) bool {
	return b.contains(b.root, k)
}

func (b *BST[K, V]) IsEmpty() bool {
	return b.size(b.root) == 0
}

func (b *BST[K, V]) Size() int {
	return b.size(b.root)
}

// Rank 返回小于 k 的元素个数
func (b *BST[K, V]) Rank(k K) int {
	return b.rank(b.root, k)
}

// Select 返回排名为rank的键
func (b *BST[K, V]) Select(rank int) (res K) {
	return b.selectR(b.root, rank)
}

func (b *BST[K, V]) Min() (res K) {
	node := b.min(b.root)
	if node != nil {
		return node.key
	}
	return
}

func (b *BST[K, V]) Max() (res K) {
	node := b.max(b.root)
	if node != nil {
		return node.key
	}
	return
}

func (b *BST[K, V]) DeleteMin() {
	b.root = b.deleteMin(b.root)
	return
}

func (b *BST[K, V]) DeleteMax() {
	b.root = b.deleteMax(b.root)
}

func (b *BST[K, V]) Keys() {
	b.queue = queue.NewQueue[K]()
	b.keys(b.root, b.Min(), b.Max())
}

// NewIterator 返回一个key的迭代器
func (b *BST[K, V]) NewIterator() *queue.Queue[K] {
	b.Keys()
	return b.queue
}

// ---------------------------------------------------------------

func (b *BST[K, V]) put(node *Node[K, V], k K, v V) *Node[K, V] {
	if node == nil {
		return NewNode[K, V](k, v).SetN(1)
	}
	// k 比当前节点小就加在左子节点，比当前节点大就加载又子节点，等于当前节点就替换value
	if k < node.key {
		node.left = b.put(node.left, k, v)
	} else if k > node.key {
		node.right = b.put(node.right, k, v)
	} else {
		node.value = v
	}
	node.N = 1 + b.size(node.left) + b.size(node.right)
	return node
}

func (b *BST[K, V]) get(node *Node[K, V], k K) *Node[K, V] {
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

func (b *BST[K, V]) del(node *Node[K, V], k K) *Node[K, V] {
	if node == nil {
		return nil
	}

	// 如果当前节点比k大，则递归左子树
	// 如果当前节点比k小，则递归右子树
	// 如果等于当前节点是目标节点则进入下面的情况
	if k < node.key {
		node.left = b.del(node.left, k)
	} else if k > node.key {
		node.right = b.del(node.right, k)
	} else {
		// 当前为目标节点
		// 如果只有一个子节点，那么直接用子节点代替当前节点即可
		// 如果有两个子节点，那么默认选择直接后继节点来替换当前节点（实际上随机选择前驱或者后继节点来替换更好，避免树失衡）
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}
		oldNode := node
		// 将直接后继元素作为被删除节点的替换节点
		node = b.min(oldNode.right)
		node.right = b.deleteMin(oldNode.right)
		node.left = oldNode.left
	}
	node.N = 1 + b.size(node.left) + b.size(node.right)
	return node
}

func (b *BST[K, V]) contains(node *Node[K, V], k K) bool {
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
func (b *BST[K, V]) size(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return 1 + b.size(node.left) + b.size(node.right)
}

func (b *BST[K, V]) rank(node *Node[K, V], k K) (res int) {
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

func (b *BST[K, V]) selectR(node *Node[K, V], rank int) (res K) {
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

func (b *BST[K, V]) min(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return b.min(node.left)
	}
	return node
}

func (b *BST[K, V]) max(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return b.max(node.right)
	}
	return node
}

func (b *BST[K, V]) deleteMin(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node.left
	}
	node.left = b.deleteMin(node.left)
	if node.left == nil {
		node.N = 1 + b.size(node.right)
	} else {
		node.left.N = 1 + b.size(node.left) + b.size(node.right)
	}
	return node
}

func (b *BST[K, V]) deleteMax(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node.left
	}
	// 向右查找到最大子节点
	node.right = b.deleteMax(node.right)
	node.right.N = 1 + b.size(node.left) + b.size(node.right)
	return node
}

func (b *BST[K, V]) keys(node *Node[K, V], low, high K) {
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

type Node[K common.IntAny, V any] struct {
	key         K
	value       V
	left, right *Node[K, V] // 左右子树
	N           int         // 以该节点为root的树中元素的个数
}

func NewNode[K common.IntAny, V any](k K, v V) *Node[K, V] {
	return &Node[K, V]{
		key:   k,
		value: v,
		N:     1,
	}
}

func (n *Node[K, V]) SetN(N int) *Node[K, V] {
	n.N = N
	return n
}
