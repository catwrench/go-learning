package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPQ(t *testing.T) {
	input := []int{1, 9, 2, 3, 4, 5, 8, 7, 0, 6}
	mpq := NewMinPQ()
	for _, v := range input {
		mpq.Insert(v)
	}

	actual := make([]int, 0, len(input))
	for mpq.Size() > 0 {
		actual = append(actual, mpq.Min())
		mpq.DelMin()
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, actual)
}
