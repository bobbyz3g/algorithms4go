package graph

import "testing"

//Test Graph
//  [0]-------- [6]     [7] ---- [8]
// | \  \        |
// | [1]  [2]    |       [9] ---- [10]
// |             |       |   \
// |     [3]     |       |    \
// |    /    \   |      [11] -- [12]
// [5] -------- [4]

var test = [][]int{
	{0, 5},
	{4, 3},
	{0, 1},
	{9, 12},
	{6, 4},
	{5, 4},
	{0, 2},
	{11, 12},
	{9, 10},
	{0, 6},
	{7, 8},
	{9, 11},
	{5, 3},
}
var adjs = [][]int{
	{6, 2, 1, 5},
	{0},
	{0},
	{5, 4},
	{5, 6, 3},
	{3, 4, 0},
	{0, 4},
	{8},
	{7},
	{11, 10, 12},
	{9},
	{9, 12},
	{11, 9},
}

func TestGraph(t *testing.T) {
	g := NewGraphWithV(13)

	es_len := len(test)

	for _, e := range test {
		g.AddEdge(e[0], e[1])
	}

	if es := g.Es(); es != es_len {
		t.Errorf("Es Error: excepted %v, got %v", es_len, es)
	}

	if vs := g.Vs(); vs != 13 {
		t.Errorf("Es Error: excepted %v, got %v", 13, vs)
	}
}

func TestGraph_Iterator(t *testing.T) {
	g := NewGraphWithV(13)
	for _, e := range test {
		g.AddEdge(e[0], e[1])
	}

	for i := range g.adj {
		iter := g.Iterator(i)

		j := 0
		for iter.HasNext() {
			if item := iter.Value(); item != adjs[i][j] {
				t.Errorf("Iterator Error: excepted %v, got %v", adjs[i][j], item)
			}
			j++
		}
	}
}
