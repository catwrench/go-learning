package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiGraphAdj(t *testing.T) {
	g := CreateDiGraph()

	// 顶点1的边集
	str := ""
	it := g.Adj("1")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "8432", str)

	str = ""
	it = g.Adj("3")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "54", str)

	str = ""
	it = g.Adj("0")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "", str)

	str = ""
	it = g.Adj("9")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "8", str)
}

func TestDiGraphReverse(t *testing.T) {
	g := CreateDiGraph()
	reverseG := g.Reverse()

	// 顶点1的边集
	str := ""
	it := reverseG.Adj("1")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "", str)

	str = ""
	it = reverseG.Adj("3")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "12", str)

	str = ""
	it = reverseG.Adj("0")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "", str)

	str = ""
	it = reverseG.Adj("9")
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "7", str)
}
