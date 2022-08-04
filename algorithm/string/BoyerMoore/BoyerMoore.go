package BoyerMoore

import "go_learning/common"

// BoyerMoore 字符串搜索
type BoyerMoore struct {
	right []int  // 辅助数组
	pat   string // 模式串
	R     int    // 字母表数量
}

func NewBoyerMoore(pat string) *BoyerMoore {
	M := len(pat)
	R := common.R256
	right := make([]int, R, R)
	for i := 0; i < R; i++ {
		right[i] = -1 // 将辅助数组初始为-1，即不存在模式串的值默认索引-1
	}
	for i := 0; i < M; i++ {
		right[pat[i]] = i // 计算模式串每个字符出现的最右的位置
	}
	return &BoyerMoore{
		right: right,
		pat:   pat,
		R:     R,
	}
}

func (b *BoyerMoore) Search(txt string) int {
	N := len(txt)
	M := len(b.pat)

	skip := 0
	for i := 0; i <= N-M; i += skip {
		skip = 0
		for j := M - 1; j >= 0; j-- {
			// 模式从右向左对比，当串模式串和文本对应位置不匹配时,使用辅助数组计算需要向右移动的距离
			if b.pat[j] != txt[i+j] {
				skip = j - b.right[txt[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}

		if skip == 0 { // 匹配成功
			return i
		}
	}
	return N // 未匹配成功
}
