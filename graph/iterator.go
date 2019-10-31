package graph

import "github.com/Kaiser925/algorithms4go/bag"

func (g *Graph) Iterator(v int) *bag.Iterator {
	return g.adj[v].Iterator()
}
