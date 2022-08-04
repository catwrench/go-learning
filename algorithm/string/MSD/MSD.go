package MSD

import (
	"go_learning/common"
)

// MSD 高位优先 排序
type MSD struct {
	N   int      // 字符串数量
	R   int      // 字母表元素总数
	M   int      // 切换小数组排序的阈值，默认小于15时切换插入排序
	aux []string // 负责数组
}

func MSdSort(data []string) {
	N := len(data)
	msd := &MSD{
		N:   N,
		R:   common.R256,
		M:   15,
		aux: make([]string, N, N),
	}
	msd.sort(data, 0, N-1, 0)
}

func (m *MSD) sort(data []string, lo, hi, d int) {
	// 字符串长度小于切换阈值，切换插入排序
	if hi <= lo+m.M {
		m.insertSort(data, lo, hi, d)
		return
	}

	// 计算频率
	count := make([]int, m.R+2, m.R+2)
	for i := lo; i <= hi; i++ {
		count[m.charAt(data[i], d)+2]++
	}
	// 计算索引
	for r := 0; r < m.R+1; r++ {
		count[r+1] += count[r]
	}
	// 分类
	for i := lo; i <= hi; i++ {
		m.aux[count[m.charAt(data[i], d)+1]] = data[i]
		count[m.charAt(data[i], d)+1]++
	}
	// 回写
	for i := lo; i <= hi; i++ {
		data[i] = m.aux[i-lo]
	}

	// 递归计算每一组数据的子数组
	for r := 0; r < m.R; r++ {
		m.sort(data, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

func (m *MSD) charAt(str string, d int) int {
	if d < len(str) {
		return int(str[d])
	} else {
		return -1
	}
}

func (m *MSD) insertSort(data []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		// 将目标元素，插入到有序数组的合适位置
		for j := i; j > lo && m.less(data[j], data[j-1], d); j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func (m *MSD) less(str1, str2 string, d int) bool {
	len1 := len(str1)
	len2 := len(str2)
	min := len1
	if len2 < len1 {
		min = len2
	}
	for i := d; i < min; i++ {
		if m.charAt(str1, i) < m.charAt(str2, i) {
			return true
		}
		if m.charAt(str1, i) > m.charAt(str2, i) {
			return false
		}
	}
	return len1 < len2
}
