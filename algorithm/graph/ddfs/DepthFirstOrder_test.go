package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstOrder(t *testing.T) {
	g := graph.CreateDiGraph()
	dfo := NewDepthFirstOrder(*g)

	pre := make([]string, 0)
	it := dfo.Pre()
	for it.HasNext() {
		pre = append(pre, it.Next())
	}
	// 遍历顺序 0-9-8 7-9-8 6-7 5-7-6 4-5 3-5-4 2-3 1-8-4-3-2
	assert.Equal(t, "0987654321", strings.Join(pre, ""))

	post := make([]string, 0)
	it = dfo.Post()
	for it.HasNext() {
		post = append(post, it.Next())
	}
	assert.Equal(t, "0897654321", strings.Join(post, ""))

	reversePost := make([]string, 0)
	it = dfo.ReversePost()
	for it.HasNext() {
		reversePost = append(reversePost, it.Next())
	}
	assert.Equal(t, "1234567980", strings.Join(reversePost, ""))
}
