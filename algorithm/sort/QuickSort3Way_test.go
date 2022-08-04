package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort3Way(t *testing.T) {
	arr1 := Input()
	arr2 := Input2()
	arr3 := Input3()
	QuickSort3Way(arr1)
	QuickSort3Way(arr2)
	QuickSort3Way(arr3)
	assert.True(t, IsSorted(arr1))
	assert.True(t, IsSorted(arr2))
	assert.True(t, IsSorted(arr3))
}
