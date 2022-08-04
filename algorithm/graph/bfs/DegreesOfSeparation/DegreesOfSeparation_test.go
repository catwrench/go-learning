package DegreesOfSeparation

import (
	"go_learning/datastructures/graph"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initDegreesOfSeparation() *DegreesOfSeparation[string] {
	g := NewDegreesOfSeparation[string]()
	g.symbolGraph = graph.CreateSymbolGraph()
	return g
}

func TestDegreesOfSeparation(t *testing.T) {
	g := initDegreesOfSeparation()

	path, ok := g.Degree("1", "9")
	assert.True(t, ok)
	assert.Equal(t, "189", strings.Join(path, ""))

	path, ok = g.Degree("5", "0")
	assert.False(t, ok)
	assert.Equal(t, "", strings.Join(path, ""))
}
