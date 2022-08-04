package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	input := []int{1, 9, 2, 3, 4, 5, 8, 7, 0, 6}
	queue := NewQueue[int]()
	for i, _ := range input {
		queue.EnQueue(input[i])
	}

	// 测试迭代
	actual := make([]int, 0, len(input))
	it := queue.NewIterator()
	for it.HasNext() {
		actual = append(actual, queue.Next())
	}
	assert.Equal(t, input, actual)

	// 测试进出队列
	actual = make([]int, 0, len(input))
	for !queue.IsEmpty() {
		actual = append(actual, queue.DeQueue())
	}
	assert.Equal(t, input, actual)
}
