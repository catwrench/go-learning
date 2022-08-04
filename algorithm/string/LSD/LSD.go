package LSD

import "go_learning/common"

// LSD 低位优先 排序
// 适用场景：长度相同的字符串排序
// 从右向左：如果将字符看作是一个256位的数，那么低位优先即从右向左检查
func LSD(data []string, Width int) {
	N := len(data)
	aux := make([]string, N, N)

	// 从右向左遍历字符
	for w := Width - 1; w >= 0; w-- {
		count := make([]int, common.R256+1, common.R256+1)
		// 计算频率
		for i := 0; i < N; i++ {
			count[data[i][w]+1]++
		}
		// 计算索引
		for r := 0; r < common.R256; r++ {
			count[r+1] += count[r]
		}
		// 分类
		for i := 0; i < N; i++ {
			aux[count[data[i][w]]] = data[i]
			count[data[i][w]]++
		}

		// 回填
		for i := 0; i < N; i++ {
			data[i] = aux[i]
		}
	}
}
