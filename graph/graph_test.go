package graph

import (
	"github.com/Kaiser925/algorithms4go/bag"
	"testing"
)

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
		t.Errorf("Es error: excepted %v, got %v", es_len, es)
	}

	if vs := g.Vs(); vs != 13 {
		t.Errorf("Es error: excepted %v, got %v", 13, vs)
	}
}

func TestGraph_Iterator(t *testing.T) {
	g := NewGraphWithV(13)
	for _, e := range test {
		g.AddEdge(e[0], e[1])
	}

	for i := range g.adj {
		iter := g.Iter(i)

		j := 0
		for iter.HasNext() {
			if item, _ := iter.Value(); item != adjs[i][j] {
				t.Errorf("Iterator error: excepted %v, got %v", adjs[i][j], item)
			}
			j++
		}
	}
}

var es = [][]int{
	{0, 5},
	{2, 4},
	{2, 3},
	{1, 2},
	{0, 1},
	{3, 4},
	{3, 5},
	{0, 2},
}

var dfspaths = [][]int{
	{0},
	{0, 2, 1},
	{0, 2},
	{0, 2, 3},
	{0, 2, 3, 4},
	{0, 2, 3, 5},
}

var bfspaths = [][]int{
	{0},
	{0, 1},
	{0, 2},
	{0, 2, 3},
	{0, 2, 4},
	{0, 5},
}

func TestGraph_DepthFirstPaths(t *testing.T) {
	g := NewGraphWithV(6)
	for _, e := range es {
		g.AddEdge(e[0], e[1])
	}

	paths := g.DepthFirstPaths(0)
	for v := 0; v < g.Vs(); v++ {
		if paths.HasPathTo(v) {
			iter := paths.PathTo(v).Iter()
			j := 0
			for iter.HasNext() {
				if val, _ := iter.Value(); val != dfspaths[v][j] {
					t.Errorf("DFS error: excepted %v, got %v", dfspaths[v][j], val)
				}
				j++
			}
		}
	}
}

func TestGraph_BreadthFirstPaths(t *testing.T) {
	g := NewGraphWithV(6)
	for _, e := range es {
		g.AddEdge(e[0], e[1])
	}
	paths := g.BreadthFirstPaths(0)
	for v := 0; v < g.Vs(); v++ {
		if paths.HasPathTo(v) {
			iter := paths.PathTo(v).Iter()
			j := 0
			for iter.HasNext() {
				if val, _ := iter.Value(); val != bfspaths[v][j] {
					t.Errorf("DFS error: excepted %v, got %v", dfspaths[v][j], val)
				}
				j++
			}
		}
	}
}

func TestGraph_ConnectedComponent(t *testing.T) {
	g := NewGraphWithV(13)
	for _, e := range test {
		g.AddEdge(e[0], e[1])
	}

	ccs := [][]int{
		{6, 5, 4, 3, 2, 1, 0},
		{8, 7},
		{12, 11, 10, 9},
	}

	cc := g.ConnectedComponent()
	m := cc.Count()
	if m != 3 {
		t.Errorf("Count error: excepted %v, got %v", 3, m)
	}

	coms := make([]bag.Bag, m)

	for v := 0; v < g.Vs(); v++ {
		coms[cc.Id(v)].Add(v)
	}

	for i := 0; i < m; i++ {
		iter := coms[i].Iter()
		j := 0
		for iter.HasNext() {
			if val, _ := iter.Value(); val.(int) != ccs[i][j] {
				t.Errorf("CC error: excepted %v, got %v", ccs[i][j], val)
			}
			j++
		}
	}
}
