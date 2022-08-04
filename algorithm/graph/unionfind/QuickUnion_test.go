package unionfind

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func QuickUnionInput() [][]int {
	input := "9-0 3-4 5-8 7-2 2-1 5-7 0-3 4-2" // 没有6
	arr := strings.Split(input, " ")
	res := make([][]int, 0)
	for _, item := range arr {
		arr2 := strings.Split(item, "-")
		k, _ := strconv.Atoi(arr2[0])
		v, _ := strconv.Atoi(arr2[1])
		res = append(res, []int{k, v})
	}
	return res
}

func TestQuickUnion(t *testing.T) {
	input := QuickUnionInput()

	q := NewQuickUnion(9)
	for _, item := range input {
		q.Union(item[0], item[1])
	}

	// 测试数据为0-9，没有6
	// 测试数据为完全连通的，所以连通分量为1，加上不存在的节点6，所以值为2
	assert.Equal(t, 2, q.Count())
}
