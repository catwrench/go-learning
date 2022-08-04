package st

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialSearchST(t *testing.T) {
	arr1 := Input1()
	st := NewSequentialSearchST[int, string]()
	for i := range arr1 {
		st.Put(i, strconv.Itoa(arr1[i]))
	}

	str, v := "", ""
	st.NewIterator()
	for st.HasNext() {
		_, v = st.Next()
		str += v
	}
	assert.Equal(t, "0675839421", str)

	assert.Equal(t, "1", st.Get(0))
	st.Del(0)
	assert.Equal(t, 9, st.Size())
	assert.Equal(t, "", st.Get(0))

	assert.Equal(t, "3", st.Get(4))
	st.Del(4)
	assert.Equal(t, 8, st.Size())
	assert.Equal(t, "", st.Get(4))
}
