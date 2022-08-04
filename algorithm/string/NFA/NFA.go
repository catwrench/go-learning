package NFA

import (
	"fmt"
	dfs2 "go_learning/algorithm/graph/ddfs"
	"go_learning/common"
	"go_learning/datastructures/bag"
	"go_learning/datastructures/graph"
	"go_learning/datastructures/stack"
)

// NFA 非确定状态自动转换机
type NFA struct {
	re []int               // 匹配转换，对应regexp每个元素在元素表的位置
	G  *graph.DiGraph[int] // 有向图
	M  int                 // 状态数量
}

func NewNFA(regexp string) *NFA {
	M := len(regexp)
	res := &NFA{}
	// 初始化M+1个顶点的图
	G := graph.NewDiGraph[int]()
	for i := 0; i <= M; i++ {
		G.AddVertex(i)
	}
	// 初始化元素数组
	re := make([]int, common.R256, common.R256)
	for i := 0; i < common.R256; i++ {
		re[i] = -1
	}
	for i, v := range []rune(regexp) {
		re[i] = int(v)
	}

	ops := stack.NewLIFOStack[int]()
	for i := 0; i < M; i++ {
		lp := i
		// 使用栈存储 (、| 两种元素
		// 匹配到 ) 时，进行出栈，直到出栈元素为 ( 为止
		if res.check(re[i], "(") || res.check(re[i], "|") {
			ops.Push(i)
		} else if res.check(re[i], ")") {
			or := ops.Pop()
			if res.check(or, "|") {
				lp = ops.Pop()
				G.AddEdge(lp, or+1)
				G.AddEdge(or, i)
			} else {
				lp = or
			}
		}
		// 查找下一个元素，加入到有向图
		if i < M-1 || res.check(re[i+1], "*") {
			G.AddEdge(lp, i+1)
			G.AddEdge(i+1, lp)
		}
		// （*）三种元素 加入到有向图
		if res.check(re[i], "(") || res.check(re[i], "*") || res.check(re[i], ")") {
			G.AddEdge(i, i+1)
		}
	}

	res.re = re
	res.G = G
	res.M = M
	return res
}

// Recognizes 识别文本
func (n *NFA) Recognizes(txt string) bool {
	pc := bag.NewBag[int]()
	dfs := dfs2.NewDirectedDFS[int](*n.G)
	dfs.Dfs(*n.G, 0)

	vertexArr := n.G.V()
	for _, vertex := range vertexArr {
		if dfs.Marked(vertex) {
			pc.Add(vertex)
		}
	}
	// 计算txt[i+1]能到达的所有状态
	for i := 0; i < len(txt); i++ {
		match := bag.NewBag[int]()
		it := pc.NewIterator()
		for it.HasNext() {
			v := it.Next()
			if v < n.M {
				if n.re[v] == int(txt[i]) || n.check(n.re[v], ".") {
					match.Add(v + 1)
				}
			}
		}
		// 基于匹配到的字符，创建新的dfs遍历
		pc = bag.NewBag[int]()
		dfs = dfs2.NewDirectedDFS[int](*n.G)
		it = match.NewIterator()
		for it.HasNext() {
			v := it.Next()
			dfs.Dfs(*n.G, v)
		}
		for _, vertex := range vertexArr {
			if dfs.Marked(vertex) {
				pc.Add(vertex)
			}
		}
	}

	it := pc.NewIterator()
	for it.HasNext() {
		v := it.Next()
		if v == n.M {
			return true
		}
	}
	return false
}

func Grep(pat string) string {
	regex := "(.*" + pat + ".*)"
	naf := NewNFA(regex)
	if naf.Recognizes(pat) {
		return pat
	}
	return ""
}

// 对比字符表元素
func (n *NFA) check(re int, char string) bool {
	return fmt.Sprintf("%c", re) == char
}
