package Huffman

import (
	"fmt"
	"go_learning/common"
)

// Huffman 霍夫曼压缩
type Huffman struct {
	r        int      // 对应字母表 字符数
	count    int      // 总字符数
	bitCount int      // 总bit数
	trieSt   []string // 单词查找树编码集，key:原始字符对应的元素表索引， val:编码后的数据
	trieBits string
	trie     *Node
}

// Node 霍夫曼单词查找树中的节点
type Node struct {
	ch          int // 该节点对应元素的字母表索引
	freq        int // 出现频率
	left, right *Node
}

func NewHuffman() *Huffman {
	return &Huffman{
		r: common.R256,
	}
}

func NewNode(ch, freq int, left, right *Node) *Node {
	return &Node{
		ch:    ch,
		freq:  freq,
		left:  left,
		right: right,
	}
}

// Compress 压缩
func (h *Huffman) Compress(txt string) (count int, trieBits, dataBits string) {
	// 1、读取输入，并转换为字母表数组
	re := make([]int, h.r, h.r)
	for i := 0; i < h.r; i++ {
		re[i] = -1
	}
	for _, i := range []rune(txt) {
		re[i] = int(i)
	}

	// 2、统计每个字符出现频率
	freq := make([]int, h.r, h.r)
	for _, i := range []rune(txt) {
		freq[i] += 1
	}

	// 3、构建霍夫曼编码树
	root := h.buildTrie(freq)
	h.trie = root

	// 4、构造编译表（递归）
	st := h.buildCode(root)

	// 5、输出解码用的单词查找树
	trieBits = h.writeTrie(root)

	// 6、输出字符总数
	h.count = root.freq

	// 7、使用霍夫曼编码处理输入
	for i := 0; i < h.count; i++ {
		code := st[txt[i]]
		dataBits += code
	}
	h.bitCount = len(dataBits)
	return h.count, trieBits, dataBits
}

// Expand 解压
func (h *Huffman) Expand(code string) (txt string) {
	// 1、读取单词查找树
	root := h.readTrie()

	// 2、读取需要解码的bit数
	count := len(code)

	// 3、使用单词查找树解码比特流
	node := root
	for i := 0; i < count; i++ {
		if string(rune(code[i])) == "0" {
			node = node.left
		} else if string(rune(code[i])) == "1" {
			node = node.right
		}
		if node.isLeaf() {
			txt += string(rune(node.ch))
			node = root
		}
	}
	return txt
}

// 构建霍夫曼单词查找树
func (h *Huffman) buildTrie(freq []int) *Node {
	// 每个元素的频率依次放入最小队列
	pq := NewNodeMinPQ[int]()
	for ch, count := range freq {
		if count > 0 {
			pq.Insert(NewNode(ch, count, nil, nil))
		}
	}

	// 合并最小俩节点为一棵树，直到队列中只剩一个节点
	for pq.Size() > 1 {
		left := pq.DelMin()
		right := pq.DelMin()
		parent := NewNode(-1, left.freq+right.freq, left, right)
		pq.Insert(parent)
	}
	return pq.DelMin()
}

// buildCode 使用单词查找树 构造编译表
func (h *Huffman) buildCode(root *Node) []string {
	st := make([]string, h.r, h.r)
	h.buildCodeRecursion(root, st, "")
	return st
}

// 构造编译表(递归)
func (h *Huffman) buildCodeRecursion(node *Node, st []string, s string) {
	if node == nil {
		return
	}
	if node.isLeaf() {
		st[node.ch] = s
		return
	}
	h.buildCodeRecursion(node.left, st, s+"0")
	h.buildCodeRecursion(node.right, st, s+"1")
}

// writeTrie 输出解码用的单词查找树
func (h *Huffman) writeTrie(root *Node) string {
	// 输出单词查找树的bit字符串
	h.writeTrieRecursion(root)
	return h.trieBits
}

// 输出解码用的单词查找树(递归)
func (h *Huffman) writeTrieRecursion(node *Node) {
	if node == nil {
		return
	}
	if node.isLeaf() {
		// TODO 将单词查找树进行编码
		h.trieBits += "1"
		h.trieBits += fmt.Sprintf("%.8b", node.ch)
		return
	}
	h.trieBits += "0"
	h.writeTrieRecursion(node.left)
	h.writeTrieRecursion(node.right)
	return
}

// TODO 从编码读取单词查找树
func (h *Huffman) readTrie() *Node {
	return h.trie
}

// TODO 从编码读取字符总数
func (h *Huffman) readCount() int {
	return h.count
}

// 是否叶子节点
func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

// 比较两个节点
// 返回值：1大于，0等于，-1小于
func (n *Node) compareTo(node *Node) int {
	if n.freq > node.freq {
		return 1
	} else if n.freq < node.freq {
		return -1
	}
	return 0
}
