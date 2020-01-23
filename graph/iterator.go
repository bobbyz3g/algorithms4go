package graph

import "github.com/Kaiser925/algorithms4go/bag"

// Iter returns the iterator of graph.
func (g *Graph) Iter(v int) *bag.Iterator {
	return g.adj[v].Iter()
}
