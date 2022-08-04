package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	arr1 := Input()
	HeapSort(arr1)
	assert.True(t, IsSorted(arr1))

	arr2 := Input2()
	HeapSort(arr2)
	assert.True(t, IsSorted(arr2))

	arr3 := Input3()
	HeapSort(arr3)
	assert.True(t, IsSorted(arr3))
}
