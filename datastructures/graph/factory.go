package graph

import (
	"strings"
)

func Input() (V []string, edge []string) {
	V = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	// 没有到0的边
	edge = []string{
		"1,2", "2,3", "3,4", "4,5", "5,6", "6,7", "7,8",
		"1,3", "3,5", "5,7", "7,9",
		"1,4", "1,8",
		"9,8",
	}
	return
}

func InputCycle() (V []string, edge []string) {
	V = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	// 没有到0的边
	edge = []string{
		"1,2", "2,3", "3,4", "4,5", "5,6", "6,7", "7,8", "8,9",
		"9,1",
	}
	return
}

func InputEdge() []Edge[string] {
	return []Edge[string]{
		{"1", "2", 1.0},
		{"2", "3", 2.0},
		{"3", "4", 3.0},
		{"4", "5", 4.0},
		{"5", "6", 6.0},
		{"6", "7", 7.0},
		{"7", "8", 8.0},
		{"8", "9", 9.0},
		{"1", "5", 1.0},
		{"5", "9", 9.0},
		{"9", "1", 9.0},
	}
}

func InputEdge2() []Edge[int] {
	return []Edge[int]{
		{1, 2, 1.0},
		{2, 3, 2.0},
		{3, 4, 3.0},
		{4, 5, 4.0},
		{5, 6, 6.0},
		{6, 7, 7.0},
		{7, 8, 8.0},
		{8, 9, 9.0},
		{1, 5, 1.0},
		{5, 9, 9.0},
		{9, 1, 9.0},
	}
}

func InputDiEdge() []DirectedEdge[string] {
	return []DirectedEdge[string]{
		{"1", "2", 1.0},
		{"2", "3", 2.0},
		{"3", "4", 3.0},
		{"4", "5", 4.0},
		{"5", "6", 6.0},
		{"6", "7", 7.0},
		{"7", "8", 8.0},
		{"8", "9", 9.0},
		{"1", "5", 1.0},
		{"5", "9", 9.0},
		// {"9", "1", 9.0},
	}
}

func InputDiEdge2() []DirectedEdge[string] {
	return []DirectedEdge[string]{
		{"1", "2", 1.0},
		{"2", "3", 2.0},
		{"3", "4", 3.0},
		{"4", "5", 4.0},
		{"5", "6", 6.0},
		{"6", "7", 7.0},
		{"7", "8", 8.0},
		{"8", "9", 9.0},
		{"1", "5", 1.0},
		{"5", "9", 9.0},
		{"9", "1", 9.0},
	}
}
func inputBellmanFordEdge() []DirectedEdge[int] {
	// *  % java BellmanFordSP tinyEWDn.txt 0
	// *  0 to 0 ( 0.00)
	// *  0 to 1 ( 0.93)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25   4->5  0.35   5->1  0.32
	// *  0 to 2 ( 0.26)  0->2  0.26
	// *  0 to 3 ( 0.99)  0->2  0.26   2->7  0.34   7->3  0.39
	// *  0 to 4 ( 0.26)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25
	// *  0 to 5 ( 0.61)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25   4->5  0.35
	// *  0 to 6 ( 1.51)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52
	// *  0 to 7 ( 0.60)  0->2  0.26   2->7  0.34
	// *
	// 	*  % java BellmanFordSP tinyEWDnc.txt 0
	// *  4->5  0.35
	// *  5->4 -0.66
	return []DirectedEdge[int]{
		{4, 5, 0.35},
		{5, 4, 0.35},
		{4, 7, 0.37},
		{5, 7, 0.28},
		{7, 5, 0.28},
		{5, 1, 0.32},
		{0, 4, 0.38},
		{0, 2, 0.26},
		{7, 3, 0.39},
		{1, 3, 0.29},
		{2, 7, 0.34},
		{6, 2, -1.20},
		{3, 6, 0.52},
		{6, 0, -1.40},
		{6, 4, -1.25},
	}
}

// CreateGraph 创建无向图
func CreateGraph() *Graph[string] {
	g := NewGraph[string]()
	V, edgeArr := Input()
	for i := range V {
		g.AddVertex(V[i])
	}
	for _, edgeItem := range edgeArr {
		vArr := strings.Split(edgeItem, ",")
		g.AddEdge(vArr[0], vArr[1])
	}
	return g
}

// CreateSymbolGraph 创建无向符号图
func CreateSymbolGraph() *SymbolGraph[string] {
	g := NewSymbolGraph[string]()
	V, edgeArr := Input()
	for i := range V {
		g.AddVertex(V[i])
	}
	for _, edgeItem := range edgeArr {
		vArr := strings.Split(edgeItem, ",")
		g.AddEdge(vArr[0], vArr[1])
	}
	return g
}

// CreateDiGraph 创建有向图
func CreateDiGraph() *DiGraph[string] {
	g := NewDiGraph[string]()
	V, edgeArr := Input()
	for i := range V {
		g.AddVertex(V[i])
	}
	for _, edgeItem := range edgeArr {
		vArr := strings.Split(edgeItem, ",")
		g.AddEdge(vArr[0], vArr[1])
	}
	return g
}

// InitCycleDiGraph 创建有向图 带环
func InitCycleDiGraph() *DiGraph[string] {
	g := NewDiGraph[string]()
	V, edgeArr := InputCycle()
	for i := range V {
		g.AddVertex(V[i])
	}
	for _, edgeItem := range edgeArr {
		vArr := strings.Split(edgeItem, ",")
		g.AddEdge(vArr[0], vArr[1])
	}
	return g
}

// CreateEdgeWeightGraph 创建加权无向图 string
func CreateEdgeWeightGraph() *EdgeWeightedGraph[string] {
	G := NewEdgeWeightedGraph[string]()
	edges := InputEdge()
	for _, edge := range edges {
		G.AddEdge(edge)
	}
	return G
}

// CreateEdgeWeightGraph2 创建加权无向图 int
func CreateEdgeWeightGraph2() *EdgeWeightedGraph[int] {
	G := NewEdgeWeightedGraph[int]()
	edges := InputEdge2()
	for _, edge := range edges {
		G.AddEdge(edge)
	}
	return G
}

// CreateEdgeWeightDiGraph 创建加权有向图 string 无环
func CreateEdgeWeightDiGraph() *EdgeWeightedDiGraph[string] {
	G := NewEdgeWeightedDiGraph[string]()
	edges := InputDiEdge()
	for _, edge := range edges {
		G.AddEdge(edge)
	}
	return G
}

// CreateEdgeWeightDiGraph2 创建加权有向图 string 带环
func CreateEdgeWeightDiGraph2() *EdgeWeightedDiGraph[string] {
	G := NewEdgeWeightedDiGraph[string]()
	edges := InputDiEdge2()
	for _, edge := range edges {
		G.AddEdge(edge)
	}
	return G
}

// CreateEdgeWeightDiGraphForBell 创建加权有向图int
func CreateEdgeWeightDiGraphForBell() *EdgeWeightedDiGraph[int] {
	G := NewEdgeWeightedDiGraph[int]()
	edges := inputBellmanFordEdge()
	for _, edge := range edges {
		G.AddEdge(edge)
	}
	return G
}
