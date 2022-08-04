package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr1 := Input()
	arr2 := Input2()
	arr3 := Input3()
	QuickSort(arr1)
	QuickSort(arr2)
	QuickSort(arr3)
	assert.True(t, IsSorted(arr1))
	assert.True(t, IsSorted(arr2))
	assert.True(t, IsSorted(arr3))
}
