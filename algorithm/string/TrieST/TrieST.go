package TrieST

import (
	"fmt"
	"go_learning/common"
	"go_learning/datastructures/queue"
)

// TrieST 基于单词查找树的符号表
type TrieST[T comparable] struct {
	R    int      // 字母表元素总数
	root *Node[T] // 根节点
}

type Node[T comparable] struct {
	value T
	next  []*Node[T]
}

func NewTrieST[T comparable]() *TrieST[T] {
	return &TrieST[T]{
		R: common.R256,
	}
}

func NewNode[T comparable](cap int) *Node[T] {
	return &Node[T]{
		next: make([]*Node[T], cap, cap),
	}
}

// Get 查找指定key
func (t *TrieST[T]) Get(key string) (res T) {
	node := t.get(t.root, key, 0)
	if node != nil {
		return node.value
	}
	return res
}

// Put 添加key
func (t *TrieST[T]) Put(key string, val T) {
	t.root = t.put(t.root, key, val, 0)
}

// Del 删除key
func (t *TrieST[T]) Del(key string) {
	t.del(t.root, key, 0)
}

// Keys 获取所有key
func (t *TrieST[T]) Keys() common.Iterable[string] {
	return t.KeysWithPrefix("")
}

// KeysWithPrefix 获取指定前缀的所有key
func (t *TrieST[T]) KeysWithPrefix(pre string) common.Iterable[string] {
	q := queue.NewQueue[string]()
	t.collect(t.get(t.root, pre, 0), pre, q)
	return q
}

// KeysThatMatch 获取通配符匹配的所有key
func (t *TrieST[T]) KeysThatMatch(pat string) common.Iterable[string] {
	q := queue.NewQueue[string]()
	t.collectPat(t.root, "", pat, q)
	return q
}

// LongestPrefixOf 获取能匹配到的最长前缀
func (t *TrieST[T]) LongestPrefixOf(key string) string {
	long := t.collectLong(t.root, key, 0, 0)
	return key[:long]
}

// 前缀匹配
func (t *TrieST[T]) collect(node *Node[T], pre string, queue *queue.Queue[string]) {
	if node == nil {
		return
	}
	var val T
	if node.value != val {
		queue.EnQueue(pre)
	}
	for c := 0; c < t.R; c++ {
		t.collect(node.next[c], fmt.Sprintf(pre+"%c", c), queue)
	}
}

// 通配符模式匹配
func (t *TrieST[T]) collectPat(node *Node[T], pre, pat string, queue *queue.Queue[string]) {
	if node == nil {
		return
	}

	d := len(pre)
	var val T
	if len(pat) == d && node.value != val {
		queue.EnQueue(pre)
	}
	if len(pat) == d {
		return
	}

	next := t.charAt(pat, d)
	for c := 0; c < t.R; c++ {
		// 下一位是通配符或者目标字符，进行递归
		if fmt.Sprintf("%c", next) == "." || next == c {
			t.collectPat(node.next[c], fmt.Sprintf(pre+"%c", c), pat, queue)
		}
	}
}

func (t *TrieST[T]) collectLong(node *Node[T], key string, d, length int) int {
	// node为空,未命中
	if node == nil {
		return length
	}
	// 命中时更新最长前缀长度
	var val T
	if node.value != val {
		length = d
	}
	// 查找完key所有字符结束
	if d == len(key) {
		return length
	}

	// 未命中，递归
	ch := t.charAt(key, d)
	return t.collectLong(node.next[ch], key, d+1, length)
}

func (t *TrieST[T]) get(node *Node[T], key string, d int) *Node[T] {
	// node为空,未命中
	if node == nil {
		return node
	}
	// 命中
	if len(key) == d {
		return node
	}
	// 未命中，递归
	ch := t.charAt(key, d)
	return t.get(node.next[ch], key, d+1)
}

func (t *TrieST[T]) put(node *Node[T], key string, val T, d int) *Node[T] {
	if node == nil {
		node = NewNode[T](t.R)
	}
	if d == len(key) {
		node.value = val
		return node
	}
	ch := t.charAt(key, d)
	node.next[ch] = t.put(node.next[ch], key, val, d+1)
	return node
}

func (t *TrieST[T]) charAt(str string, d int) int {
	if d < len(str) {
		return int(str[d])
	} else {
		return -1
	}
}

func (t *TrieST[T]) del(node *Node[T], key string, d int) *Node[T] {
	if node == nil {
		return nil
	}

	var val T
	if d == len(key) {
		// 命中,值设置为空
		node.value = val
	} else {
		// 未命中，递归查找
		c := t.charAt(key, d)
		node.next[c] = t.del(node.next[c], key, d+1)
	}

	// 检查是否为最终节点，是最终节点且没有后续节点，就删除当前节点
	if node.value != val {
		return node
	}
	for r := 0; r < t.R; r++ {
		if node.next[r] != nil {
			return node
		}
	}
	return nil
}
