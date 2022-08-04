package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	arr1 := Input()
	ShellSort(arr1)
	assert.True(t, IsSorted(arr1))

	arr2 := Input2()
	ShellSort(arr2)
	assert.True(t, IsSorted(arr2))

	arr3 := Input3()
	ShellSort(arr3)
	assert.True(t, IsSorted(arr3))
}
