package TST

import (
	"fmt"
	"go_learning/common"
	"go_learning/datastructures/queue"
)

// TST 三向单词查找树
type TST[T comparable] struct {
	root *Node[T]
}

type Node[T comparable] struct {
	c                int
	value            T
	left, mid, right *Node[T]
}

func NewTST[T comparable]() *TST[T] {
	return &TST[T]{}
}

func NewNode[T comparable]() *Node[T] {
	return &Node[T]{}
}

// Get 查找指定key
func (t *TST[T]) Get(key string) (res T) {
	node := t.get(t.root, key, 0)
	if node != nil {
		return node.value
	}
	return res
}

// Put 添加key
func (t *TST[T]) Put(key string, val T) {
	t.root = t.put(t.root, key, val, 0)
}

// Del 删除key
func (t *TST[T]) Del(key string) {
	t.del(t.root, key, 0)
}

// KeysWithPrefix 获取指定前缀的所有key
func (t *TST[T]) KeysWithPrefix(pre string) common.Iterable[string] {
	var val T
	q := queue.NewQueue[string]()
	node := t.get(t.root, pre, 0)
	if node == nil {
		return q
	}
	if node.value != val {
		q.EnQueue(pre)
	}

	t.collect(node.mid, pre, q)
	return q
}

// KeysThatMatch 获取通配符匹配的所有key
func (t *TST[T]) KeysThatMatch(pat string) common.Iterable[string] {
	q := queue.NewQueue[string]()
	t.collectPat(t.root, "", pat, q)
	return q
}

// LongestPrefixOf 获取能匹配到的最长前缀
func (t *TST[T]) LongestPrefixOf(key string) string {
	long := t.collectLong(t.root, key, 0, 0)
	return key[:long]
}

func (t *TST[T]) get(node *Node[T], key string, d int) *Node[T] {
	if node == nil {
		return nil
	}
	ch := t.charAt(key, d)
	if ch < node.c {
		return t.get(node.left, key, d)
	} else if ch > node.c {
		return t.get(node.right, key, d)
	} else if d < len(key)-1 {
		return t.get(node.mid, key, d+1)
	}
	return node
}

func (t *TST[T]) put(node *Node[T], key string, val T, d int) *Node[T] {
	ch := t.charAt(key, d)
	if node == nil {
		node = NewNode[T]()
		node.c = ch
	}

	// 根据当前字符和node节点字符比较，进行三向切分
	if ch < node.c {
		node.left = t.put(node.left, key, val, d)
	} else if ch > node.c {
		node.right = t.put(node.right, key, val, d)
	} else {
		// 相等时有两种情况，字符为key末尾字符和非末尾字符
		// 非末尾字符：从mid向下递归
		// 末尾字符：更新node值
		if d < len(key)-1 {
			node.mid = t.put(node.mid, key, val, d+1)
		} else {
			node.value = val
		}
	}
	return node
}

func (t *TST[T]) del(node *Node[T], key string, d int) *Node[T] {
	if node == nil {
		return nil
	}
	ch := t.charAt(key, d)

	var val T
	if d == len(key) {
		// 命中,值设置为空
		node.value = val
	} else {
		// 未命中，递归查找
		if ch < node.c {
			node = t.del(node.left, key, d)
		} else if ch > node.c {
			node = t.del(node.right, key, d)
		} else if d < len(key)-1 {
			node = t.del(node.mid, key, d+1)
		}
	}

	// 检查是否为最终节点，是最终节点且没有后续节点，就删除当前节点
	if node.value != val {
		return node
	}
	if node.left != nil || node.mid != nil || node.right != nil {
		return node
	}
	return nil
}

func (t *TST[T]) charAt(str string, d int) int {
	if d < len(str) {
		return int(str[d])
	} else {
		return -1
	}
}

// 前缀匹配
func (t *TST[T]) collect(node *Node[T], pre string, queue *queue.Queue[string]) {
	if node == nil {
		return
	}
	var val T
	if node.value != val {
		queue.EnQueue(fmt.Sprintf(pre+"%c", node.c))
	}

	if node.left != nil {
		t.collect(node.left, pre, queue)
	}
	if node.right != nil {
		t.collect(node.right, pre, queue)
	}
	if node.mid != nil {
		t.collect(node.mid, fmt.Sprintf(pre+"%c", node.c), queue)
	}
}

// 通配符模式匹配
func (t *TST[T]) collectPat(node *Node[T], pre, pat string, queue *queue.Queue[string]) {
	if node == nil {
		return
	}

	var val T
	d := len(pre)
	next := t.charAt(pat, d)

	// 下一位是通配符或者目标字符，进行递归
	if fmt.Sprintf("%c", next) == "." || next < node.c {
		t.collectPat(node.left, pre, pat, queue)
	}
	if fmt.Sprintf("%c", next) == "." || next > node.c {
		t.collectPat(node.right, pre, pat, queue)
	}
	if fmt.Sprintf("%c", next) == "." || next == node.c {
		if d == len(pat)-1 && node.value != val {
			queue.EnQueue(fmt.Sprintf(pre+"%c", node.c))
		}
		if d < len(pat)-1 {
			t.collectPat(node.mid, fmt.Sprintf(pre+"%c", node.c), pat, queue)
		}
	}
}

func (t *TST[T]) collectLong(node *Node[T], key string, d, length int) int {
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
	if d < len(key) {
		if ch < node.c {
			return t.collectLong(node.left, key, d, length)
		} else if ch > node.c {
			return t.collectLong(node.right, key, d, length)
		} else if d < len(key)-1 {
			return t.collectLong(node.mid, key, d+1, length)
		} else {
			length += 1
		}
	}
	return length
}
