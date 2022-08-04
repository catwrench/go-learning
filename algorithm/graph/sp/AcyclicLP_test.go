package sp

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAcyclicLP(t *testing.T) {
	G := graph.CreateEdgeWeightDiGraph()

	sp := NewAcyclicLP(*G, "5")
	assert.False(t, sp.HasPathTo("1"))
	assert.True(t, sp.HasPathTo("9"))
	assert.True(t, sp.HasPathTo("7"))

	path := make([]string, 0)
	it := sp.PathTo("9")
	for it.HasNext() {
		edge := it.Next()
		path = append(path, edge.To())
	}
	assert.Equal(t, "6789", strings.Join(path, ""))
	assert.Equal(t, float64(30), sp.DistTo("9"))

	path = make([]string, 0)
	it = sp.PathTo("7")
	for it.HasNext() {
		edge := it.Next()
		path = append(path, edge.To())
	}
	assert.Equal(t, "67", strings.Join(path, ""))
	assert.Equal(t, float64(13), sp.DistTo("7"))
}
