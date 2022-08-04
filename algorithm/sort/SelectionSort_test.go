package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert.True(t, IsSorted(SelectionSort(Input())))
	assert.True(t, IsSorted(SelectionSort(Input2())))
	assert.True(t, IsSorted(SelectionSort(Input3())))
}
