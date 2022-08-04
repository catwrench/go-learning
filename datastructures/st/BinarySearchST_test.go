package st

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitBinarySearchST() *BinarySearchST[int, string] {
	arr1 := Input1()
	st := NewBinarySearchST[int, string](len(arr1))
	for i := range arr1 {
		st.Put(arr1[i], strconv.Itoa(arr1[i]))
	}
	return st
}

func TestBinarySearchST(t *testing.T) {
	st := InitBinarySearchST()
	str := ""

	st.NewIterator()
	for st.HasNext() {
		_, v := st.Next()
		str += v
	}
	assert.Equal(t, "0123456789", str)
}

func TestBinarySearchSTDel(t *testing.T) {
	st := InitBinarySearchST()
	assert.Equal(t, "0", st.Get(0))
	st.Del(0)
	assert.Equal(t, "", st.Get(0))
	assert.Equal(t, "4", st.Get(4))
	st.Del(4)
	assert.Equal(t, "", st.Get(4))
}

func TestBinarySearchSTMin(t *testing.T) {
	st := InitBinarySearchST()
	assert.Equal(t, "0", st.Min())
	st.Del(0)
	assert.Equal(t, "1", st.Min())

	st = InitBinarySearchST()
	keys := st.keys
	for i := range keys {
		st.Del(keys[i])
	}
	assert.Equal(t, "", st.Min())
}

func TestBinarySearchSTMax(t *testing.T) {
	st := InitBinarySearchST()
	assert.Equal(t, "9", st.Max())
	st.Del(9)
	assert.Equal(t, "8", st.Max())

	st = InitBinarySearchST()
	keys := st.keys
	for i := range keys {
		st.Del(keys[i])
	}
	assert.Equal(t, "", st.Max())
}
