package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstOrderDWG(t *testing.T) {
	g := graph.CreateEdgeWeightDiGraph()
	dfo := NewDepthFirstOrderDWG(*g)

	pre := make([]string, 0)
	it := dfo.Pre()
	for it.HasNext() {
		pre = append(pre, it.Next())
	}
	assert.Equal(t, "987654321", strings.Join(pre, ""))

	post := make([]string, 0)
	it = dfo.Post()
	for it.HasNext() {
		post = append(post, it.Next())
	}
	assert.Equal(t, "987654321", strings.Join(post, ""))

	reversePost := make([]string, 0)
	it = dfo.ReversePost()
	for it.HasNext() {
		reversePost = append(reversePost, it.Next())
	}
	assert.Equal(t, "123456789", strings.Join(reversePost, ""))
}
