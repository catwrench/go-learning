package dfs

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReachable(t *testing.T) {
	g := graph.CreateDiGraph()
	ddfs := NewDirectedDFS(*g)

	reach1 := ddfs.Reachable(*g, "0")
	assert.Equal(t, "0", strings.Join(reach1, ""))

	reach2 := ddfs.Reachable(*g, "0", "9")
	assert.Equal(t, "098", strings.Join(reach2, ""))

	reach3 := ddfs.Reachable(*g, "5", "0", "9")
	assert.Equal(t, "098765", strings.Join(reach3, ""))

	reach4 := ddfs.Reachable(*g, "1", "5", "0", "9")
	assert.Equal(t, "0987654321", strings.Join(reach4, ""))

}
