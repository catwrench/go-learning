package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolIndexMinPQ(t *testing.T) {
	inputKey := []string{"A", "9", "B", "C", "4", "5", "8", "7", "6"}
	inputVal := []int{1, 9, 2, 3, 4, 5, 8, 7, 6}
	mpq := NewSymbolIndexMinPQ[string, int](len(inputKey))
	for i, _ := range inputKey {
		mpq.Insert(inputKey[i], inputVal[i])
	}

	actualKey := make([]string, 0, len(inputKey))
	actualValue := make([]int, 0, len(inputVal))

	for mpq.Size() > 0 {
		k, v := mpq.MinIdx(), mpq.Min()
		actualKey = append(actualKey, k)
		actualValue = append(actualValue, v)
		mpq.DelMin()
	}
	assert.Equal(t, []string{"A", "B", "C", "4", "5", "6", "7", "8", "9"}, actualKey)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, actualValue)
}
