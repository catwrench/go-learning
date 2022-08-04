package TST

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTST() *TST[bool] {
	data := []string{
		"she",
		"sells",
		"sea",
		"shells",
		"by",
		"the",
		"sea",
		"shore",
	}
	st := NewTST[bool]()
	for _, key := range data {
		st.Put(key, true)
	}
	return st
}

func TestKeysWithPrefix(t *testing.T) {
	st := createTST()

	expect := make([]string, 0)
	it := st.KeysWithPrefix("se")
	for it.HasNext() {
		expect = append(expect, it.Next())
	}
	assert.Equal(t, []string{"sea", "sells"}, expect)
}

func TestLongestPrefixOf(t *testing.T) {
	st := createTST()

	str := st.LongestPrefixOf("shell")
	assert.Equal(t, "she", str)

	str = st.LongestPrefixOf("shells")
	assert.Equal(t, "shells", str)
}

func TestKeysThatMatch(t *testing.T) {
	st := createTST()

	expect := make([]string, 0)
	it := st.KeysThatMatch("se.")
	for it.HasNext() {
		expect = append(expect, it.Next())
	}
	assert.Equal(t, []string{"sea"}, expect)

	expect = make([]string, 0)
	it = st.KeysThatMatch("se.ls")
	for it.HasNext() {
		expect = append(expect, it.Next())
	}
	assert.Equal(t, []string{"sells"}, expect)
}
