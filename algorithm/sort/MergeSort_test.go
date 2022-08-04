package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeBUSort(t *testing.T) {
	arr1 := Input()
	arr2 := Input2()
	arr3 := Input3()
	ms := NewMergeSort()

	ms.MergeBUSort(arr1)
	assert.True(t, IsSorted(arr1))

	ms.MergeBUSort(arr2)
	assert.True(t, IsSorted(arr2))

	ms.MergeBUSort(arr3)
	assert.True(t, IsSorted(arr3))
}

func TestMergeUBSort(t *testing.T) {
	arr1 := Input()
	arr2 := Input2()
	arr3 := Input3()
	ms := NewMergeSort()

	ms.MergeUBSort(arr1)
	assert.True(t, IsSorted(arr1))

	ms.MergeUBSort(arr2)
	assert.True(t, IsSorted(arr2))

	ms.MergeUBSort(arr3)
	assert.True(t, IsSorted(arr3))
}
