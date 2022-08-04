package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPQ(t *testing.T) {
	input := []int{1, 9, 2, 3, 4, 5, 8, 7, 0, 6}
	mpq := NewMaxPQ()
	for i, _ := range input {
		mpq.Insert(input[i])
	}

	actual := make([]int, 0, len(input))
	for mpq.Size() > 0 {
		actual = append(actual, mpq.Max())
		mpq.DelMax()
	}
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, actual)
}
