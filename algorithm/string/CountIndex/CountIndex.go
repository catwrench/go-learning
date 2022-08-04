package CountIndex

import (
	"go_learning/common"
	"strconv"
)

// CountIndex 键索引记数法
func CountIndex(data []string) {
	N := len(data)
	aux := make([]string, N, N)
	count := make([]int, common.R256+1, common.R256+1)
	group := make([]int, N, N)
	// 预处理分组,以字符串的第一位元素作为分组
	for i := 0; i < N; i++ {
		g, _ := strconv.Atoi(data[i][:1])
		group[i] = g
	}

	// 计算出现频率,对应分组计数+1
	for i := 0; i < N; i++ {
		// 以字符串的第一位元素作为分组
		// data[i][0] 为数组第一个字符串，索引为0元素对应的ascii编码值。4PGC938 第一位 为 4，对应ascii编码为52, count[53]++
		count[data[i][0]+1]++
	}
	// 将频率转换为索引，分组计数汇总，得到每个分组的开始索引
	for r := 0; r < common.R256; r++ {
		count[r+1] += count[r]
	}
	// 将元素分类，元素按所属分组索引顺序放到aux辅助数组里
	for i := 0; i < N; i++ {
		aux[count[data[i][0]]] = data[i]
		count[data[i][0]]++ // 分组开始索引移到下一位
	}
	// 回写
	for i := 0; i < N; i++ {
		data[i] = aux[i]
	}
}
