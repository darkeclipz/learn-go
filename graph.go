package foohandler

var MatrixGraph = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var AdjacencyList = map[int][]int{
	1: {2, 4},
	2: {3, 5, 1},
	3: {6, 2},
	4: {1, 5, 7},
	5: {2, 6, 8, 4},
	6: {3, 0, 9, 5}, // why is 0 in here???
	7: {4, 8},
	8: {5, 9, 7},
	9: {6, 0, 8},
}

var AdjacencyMatrix = [][]int{
	{0, 1, 0, 1, 0, 0, 0, 0, 0},
	{1, 0, 1, 0, 1, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 1, 0, 0, 0},
	{1, 0, 0, 0, 1, 0, 1, 0, 0},
	{0, 1, 0, 1, 0, 1, 0, 1, 0},
	{0, 0, 1, 0, 1, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 1, 0, 1, 0, 1},
	{0, 0, 0, 0, 0, 1, 0, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 1, 0},
}

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Val   int
	Edges map[int]*Edge
}

type Edge struct {
	Weight int
	Vertex *Vertex
}
