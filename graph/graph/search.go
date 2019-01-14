package graph

import (
	"github.com/Kaiser925/algorithms4go/queue/queue"
	"github.com/Kaiser925/algorithms4go/stack"
)

type Paths struct {
	visited []bool // current vertex has been visited.
	edgeTo  []int
	start   int // start vertex.
}

// DepthFirstPaths returns a new DepthFirstPaths with start vertex.
func (g *Graph) DepthFirstPaths(start int) Paths {
	p := Paths{
		visited: make([]bool, g.Vs(), g.Vs()),
		edgeTo:  make([]int, g.Vs(), g.Vs()),
		start:   start,
	}
	p.dfs(g, start)
	return p
}

// BreadthFirstPaths returns a new BreadthFirstPaths with start vertex.
func (g *Graph) BreadthFirstPaths(start int) Paths {
	p := Paths{
		visited: make([]bool, g.Vs(), g.Vs()),
		edgeTo:  make([]int, g.Vs(), g.Vs()),
		start:   start,
	}
	p.bfs(g, start)
	return p
}

func (p *Paths) dfs(g *Graph, v int) {
	p.visited[v] = true
	iter := g.Iterator(v)

	for iter.HasNext() {
		w := iter.Value().(int)
		if !p.visited[w] {
			p.edgeTo[w] = v
			p.dfs(g, w)
		}
	}
}

func (p *Paths) bfs(g *Graph, s int) {
	q := queue.New()
	p.visited[s] = true
	q.Enqueue(s)

	for !q.Empty() {
		vi, _ := q.Dequeue()
		v := vi.(int)
		iter := g.Iterator(v)

		for iter.HasNext() {
			w := iter.Value().(int)
			if !p.visited[w] {
				p.edgeTo[w] = v
				p.visited[w] = true
				q.Enqueue(w)
			}
		}
	}
}

// HasPathTo returns true if there was a path from start to v, or false if there was not.
func (p *Paths) HasPathTo(v int) bool {
	return p.visited[v]
}

// PathTo returns a stack stored path to v, or nil if there was no path from start to v.
func (p *Paths) PathTo(v int) *stack.Stack {
	if !p.HasPathTo(v) {
		return nil
	}
	path := stack.New()

	for x := v; x != p.start; x = p.edgeTo[x] {
		path.Push(x)
	}
	path.Push(p.start)
	return path
}
