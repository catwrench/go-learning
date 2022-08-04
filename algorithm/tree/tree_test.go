package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	tree := NewTree("5")
	tree.AppendLeft(tree.root, "4")
	tree.AppendLeft(tree.root.Left, "2")
	tree.AppendRight(tree.root.Left, "3")

	tree.AppendRight(tree.root, "6")
	tree.AppendLeft(tree.root.Right, "7")
	tree.AppendRight(tree.root.Right, "8")

	tree.PrevTraverse(tree.root)
	assert.Equal(t, []string{"5", "4", "2", "3", "6", "7", "8"}, tree.Values)

	tree.Values = make([]string, 0)
	tree.MiddleTraverse(tree.root)
	assert.Equal(t, []string{"2", "4", "3", "5", "7", "6", "8"}, tree.Values)

	tree.Values = make([]string, 0)
	tree.NextTraverse(tree.root)
	assert.Equal(t, []string{"2", "3", "4", "7", "8", "6", "5"}, tree.Values)
}
