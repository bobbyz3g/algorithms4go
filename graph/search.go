package graph

import (
	"github.com/Kaiser925/algorithms4go/queue"
	"github.com/Kaiser925/algorithms4go/stack"
)

type Paths struct {
	visited []bool // current vertex has been visited.
	edgeTo  []int
	start   int // start vertex.
}

// DepthFirstPaths returns a new DepthFirstPaths with start vertex.
func (g *Graph) DepthFirstPaths(start int) *Paths {
	p := &Paths{
		visited: make([]bool, g.Vs(), g.Vs()),
		edgeTo:  make([]int, g.Vs(), g.Vs()),
		start:   start,
	}
	p.dfs(g, start)
	return p
}

// BreadthFirstPaths returns a new BreadthFirstPaths with start vertex.
func (g *Graph) BreadthFirstPaths(start int) *Paths {
	p := &Paths{
		visited: make([]bool, g.Vs(), g.Vs()),
		edgeTo:  make([]int, g.Vs(), g.Vs()),
		start:   start,
	}
	p.bfs(g, start)
	return p
}

func (p *Paths) dfs(g *Graph, v int) {
	p.visited[v] = true
	iter := g.Iter(v)

	for iter.HasNext() {
		w, _ := iter.Value()
		index := w.(int)
		if !p.visited[index] {
			p.edgeTo[index] = v
			p.dfs(g, index)
		}
	}
}

func (p *Paths) bfs(g *Graph, s int) {
	q := queue.NewQueue[int]()
	p.visited[s] = true
	q.Enqueue(s)

	for !q.Empty() {
		vi, _ := q.Dequeue()
		v := vi
		iter := g.Iter(v)

		for iter.HasNext() {
			w, _ := iter.Value()
			index := w.(int)
			if !p.visited[index] {
				p.edgeTo[index] = v
				p.visited[index] = true
				q.Enqueue(index)
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

// CC represents connected component.
type CC struct {
	visited []bool
	id      []int // store the connected components.
	count   int   // number of connected components.
}

func (g *Graph) ConnectedComponent() *CC {
	vs := g.Vs()
	cc := &CC{
		visited: make([]bool, vs, vs),
		id:      make([]int, vs, vs),
		count:   0,
	}

	for s := 0; s < vs; s++ {
		if !cc.visited[s] {
			cc.dfs(g, s)
			cc.count++
		}
	}

	return cc
}

func (cc *CC) dfs(g *Graph, v int) {
	cc.visited[v] = true
	cc.id[v] = cc.count

	iter := g.Iter(v)

	for iter.HasNext() {
		w, _ := iter.Value()
		index := w.(int)
		if !cc.visited[index] {
			cc.dfs(g, index)
		}
	}
}

// Connected return true if v is connected to w, or false if not.
func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *CC) Id(v int) int {
	return cc.id[v]
}

func (cc *CC) Count() int {
	return cc.count
}
