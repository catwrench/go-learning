package KMP

// KMP 算法
// 参考：https://juejin.cn/post/6854573206896918542
// 思想：
// 	当出现不匹配时，就能知晓一部分内容，可以利用这些信息将指针回退到所有这些已知字符之前
// 场景：
// 	适合查找具有重复键的模式串
// ps
// 	1、部分匹配表pmt 等于 next表 整体左移一位，没有开始的-1位置
// 	2、这里主要是通过计算最长前后缀，计算的next表

type KMP struct {
	baseStr string // 原始字符串
	baseLen int    // 原始字符串长度
}

func NewKMP(baseStr string) *KMP {
	return &KMP{
		baseStr: baseStr,
		baseLen: len(baseStr),
	}
}

// KMPSearch kmp搜索算法
// return 子串在第pos个字符后的位置
func (k *KMP) KMPSearch(subStr string, pos int) int {
	baseIndex, subIndex := pos, 0 // 原始串起始位置pos,子串起始10

	subLen := len(subStr)
	next := make([]int, k.baseLen+1)
	k.getNext(subStr, next)

	for baseIndex < k.baseLen && subIndex < subLen {
		if subIndex == -1 || k.baseStr[baseIndex] == subStr[subIndex] {
			// 两串元素对应位置相等时，索引都+1
			baseIndex++
			subIndex++
		} else {
			// baseIndex位置不变,subIndex跳到下一个位置
			subIndex = next[subIndex]
		}
	}

	// 退出条件: subIndex 大于子串长度
	if subIndex == subLen {
		return baseIndex - subIndex
	}
	return -1
}

// getNext 计算下一次开始的位置
func (k *KMP) getNext(subStr string, next []int) {
	i, j := 1, 0
	subLen := len(subStr)

	next[0] = -1 // 默认首位-1，避免后面回溯到开头时 进入死循环

	for i < subLen {
		if j == -1 || subStr[i] == subStr[j] {
			i++
			j++
			next[i] = j
		} else {
			// 如果字符串不相等，进行回溯
			j = next[j]
		}
	}

	return
}

// KMPSearch2 kmp搜索算法 (优化了子串元素连续相同的情况)
func (k *KMP) KMPSearch2(subStr string, pos int) int {
	baseIndex, subIndex := pos, 0 // 原始串起始位置pos,子串起始10

	subLen := len(subStr)
	next := make([]int, k.baseLen+1)
	k.getNextVal(subStr, next)

	for baseIndex < k.baseLen && subIndex < subLen {
		if subIndex == -1 || k.baseStr[baseIndex] == subStr[subIndex] {
			baseIndex++
			subIndex++
		} else {
			subIndex = next[subIndex]
		}
	}

	if subIndex == subLen {
		return baseIndex - subIndex
	}
	return -1
}

// getNextVal 计算下一次开始的位置
// ps:优化了子串元素连续相同的情况
func (k *KMP) getNextVal(subStr string, next []int) {
	i, j := 1, 0
	subLen := len(subStr)

	next[0] = -1

	for i < subLen-1 {
		if j == -1 || subStr[i] == subStr[j] {
			i++
			j++
			if subStr[i] == subStr[j] { // 因为这里是在i++后进行的判断，所以for循环跳出条件需要-1
				// 前后缀字符相等时，让靠后的i位置赋值为靠前的j位置，减少回溯次数
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			// 如果字符串不相等，进行回溯
			j = next[j]
		}
	}

	return
}
