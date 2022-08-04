package prim

import (
	"go_learning/datastructures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimMST(t *testing.T) {
	g := graph.CreateEdgeWeightGraph()

	Prim := NewPrimMST(*g)

	actual := make([]float64, 0)
	mst := Prim.Mst()
	for mst.HasNext() {
		actual = append(actual, mst.Next().Weight())
	}

	// {"1", "2", 1.0},
	// {"2", "3", 2.0},
	// {"3", "4", 3.0},
	// {"4", "5", 4.0}, X 4-5连接了节点5
	// {"5", "6", 6.0},
	// {"6", "7", 7.0},
	// {"7", "8", 8.0},
	// {"8", "9", 9.0},
	// {"1", "5", 1.0},
	// {"5", "9", 9.0}, X
	// {"9", "1", 9.0}, X
	assert.Equal(t, []float64{9, 1, 8, 1, 6, 3, 2, 7}, actual)
	assert.Equal(t, float64(37), Prim.Weight())
}
