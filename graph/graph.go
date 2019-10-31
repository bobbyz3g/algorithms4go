package graph

import "github.com/Kaiser925/algorithms4go/bag"

type Graph struct {
	v   int       // number of vertexes.
	e   int       // number of edges.
	adj []bag.Bag // Adjacent list.
}

// NewGraphWithV returns a graph with v vertexes.
func NewGraphWithV(v int) *Graph {
	return &Graph{
		v:   v,
		e:   0,
		adj: make([]bag.Bag, v, v),
	}
}

// Vs returns number of vertexes.
func (g *Graph) Vs() int {
	return g.v
}

// Es returns number of edges.
func (g *Graph) Es() int {
	return g.e
}

// AddEdge adds the new edge between v and w to graph.
func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.e++
}
