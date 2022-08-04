package KMP

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMPBase(t *testing.T) {
	kmp2 := NewKMP("XXYXXYXXY")
	index2 := kmp2.KMPSearch("XXYXXYXXX", 0)
	assert.Equal(t, -1, index2)
}

func TestKMP(t *testing.T) {
	kmp := NewKMP("Hello world !")

	index := kmp.KMPSearch("wor", 0)
	assert.Equal(t, 6, index)
	index = kmp.KMPSearch("ll", 0)
	assert.Equal(t, 2, index)
	index = kmp.KMPSearch("lasd", 0)
	assert.Equal(t, -1, index)
	index = kmp.KMPSearch("Hello", 0)
	assert.Equal(t, 0, index)
	index = kmp.KMPSearch("Hello", 3)
	assert.Equal(t, -1, index)
	return
}

func TestKMP2(t *testing.T) {
	kmp := NewKMP("Hello world !")

	index := kmp.KMPSearch2("wor", 0)
	assert.Equal(t, 6, index)
	index = kmp.KMPSearch2("ll", 0)
	assert.Equal(t, 2, index)
	index = kmp.KMPSearch2("lasd", 0)
	assert.Equal(t, -1, index)
	index = kmp.KMPSearch2("Hello", 0)
	assert.Equal(t, 0, index)
	index = kmp.KMPSearch2("Hello", 3)
	assert.Equal(t, -1, index)
	return
}
