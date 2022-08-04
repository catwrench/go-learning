package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolGraph(t *testing.T) {
	g := CreateSymbolGraph()

	assert.True(t, g.Contains("9"))
	assert.True(t, g.Contains("0"))

	assert.Equal(t, "1", g.Name(0))
	assert.Equal(t, "0", g.Name(9))

	assert.Equal(t, 0, g.Index("1"))
	assert.Equal(t, 9, g.Index("0"))

	assert.True(t, g.Contains("0"))
	assert.False(t, g.Contains("11"))
}
