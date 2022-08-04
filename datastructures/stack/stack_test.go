package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Input1() []string {
	return []string{"1", "2", "4", "9", "3", "8", "5", "7", "6", "0"}
}

func TestLIFOStack(t *testing.T) {
	arr1 := Input1()
	st := NewLIFOStack[string]()
	for i := range arr1 {
		st.Push(arr1[i])
	}

	str := ""
	st.NewIterator()
	for st.HasNext() {
		str += st.Next()
	}
	assert.Equal(t, "0675839421", str)
}
