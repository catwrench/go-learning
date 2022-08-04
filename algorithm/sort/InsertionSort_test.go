package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr1 := Input()
	InsertSort(arr1)
	assert.True(t, IsSorted(arr1))

	arr2 := Input2()
	InsertSort(arr2)
	assert.True(t, IsSorted(arr2))

	arr3 := Input3()
	InsertSort(arr3)
	assert.True(t, IsSorted(arr3))
}
