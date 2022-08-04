package st

import (
	"go_learning/datastructures/queue"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitTestBST() *BST[int, string] {
	arr1 := Input1()
	st := NewBST[int, string]()
	for i := range arr1 {
		st.Put(arr1[i], strconv.Itoa(arr1[i]))
	}
	return st
}

func TestBST(t *testing.T) {
	st := InitTestBST()

	str := ""
	it := st.NewIterator()
	for it.HasNext() {
		k := it.Next()
		str += st.Get(k)
	}
	assert.Equal(t, "0123456789", str)
}

func TestBSTDel(t *testing.T) {
	st := InitTestBST()
	assert.Equal(t, "0", st.Get(0))
	st.Del(0)
	assert.Equal(t, "", st.Get(0))
	assert.Equal(t, "4", st.Get(4))
	st.Del(4)
	assert.Equal(t, "", st.Get(4))
}

func TestBSTMin(t *testing.T) {
	st := InitTestBST()
	assert.Equal(t, 0, st.Min())
	st.Del(0)
	assert.Equal(t, 1, st.Min())

	st = InitTestBST()
	it := st.NewIterator()
	for it.HasNext() {
		st.Del(it.Next())
	}
	assert.Equal(t, 0, st.Min())
}

func TestBSTMax(t *testing.T) {
	st := InitTestBST()
	assert.Equal(t, 9, st.Max())
	st.Del(9)
	assert.Equal(t, 8, st.Max())

	st = InitTestBST()
	it := st.NewIterator()
	for it.HasNext() {
		st.Del(it.Next())
	}
	assert.Equal(t, 0, st.Max())
}

func TestBSTKeys(t *testing.T) {
	st := InitTestBST()
	str := ""
	st.queue = queue.NewQueue[int]()
	st.keys(st.root, 2, 8)
	it := st.queue
	for it.HasNext() {
		k := it.Next()
		str += st.Get(k)
	}
	assert.Equal(t, "2345678", str)

	str = ""
	st.queue = queue.NewQueue[int]()
	st.keys(st.root, 1, 9)
	it = st.queue
	for it.HasNext() {
		k := it.Next()
		str += st.Get(k)
	}
	assert.Equal(t, "123456789", str)
}

func TestBSTSelect(t *testing.T) {
	st := InitTestBST()
	assert.Equal(t, 0, st.Select(0))
	assert.Equal(t, 1, st.Select(1))
	assert.Equal(t, 9, st.Select(9))
	assert.Equal(t, 5, st.Select(5))

	st.Del(0)
	assert.Equal(t, 1, st.Select(0))
	assert.Equal(t, 9, st.Select(8))
	assert.Equal(t, 0, st.Select(9))
}
