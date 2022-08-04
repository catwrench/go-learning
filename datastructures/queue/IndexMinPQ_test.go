package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexMinPQ(t *testing.T) {
	input := []int{1, 9, 2, 3, 4, 5, 8, 7, 0, 6}
	mpq := NewIndexMinPQ[int](len(input))
	for i, _ := range input {
		mpq.Insert(i+1, &input[i])
	}

	actualIdx := make([]int, 0, len(input))
	actual := make([]int, 0, len(input))
	for mpq.Size() > 0 {
		actualIdx = append(actualIdx, mpq.MinIdx())
		actual = append(actual, *mpq.Min())
		mpq.DelMin()
	}
	assert.Equal(t, []int{9, 1, 3, 4, 5, 6, 10, 8, 7, 2}, actualIdx)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, actual)
}
