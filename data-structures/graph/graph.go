package graph

import (
	"sort"
)

// Graph ...
type Graph struct {
	vertices map[*Node]bool
	edges    map[Linker]bool
}

// NewGraph ...
func NewGraph() Graph {
	return Graph{
		vertices: map[*Node]bool{},
		edges:    map[Linker]bool{},
	}
}

// IsDirected says if g is directed or not.
// If at least one linker on g is not an `Arrow` type then it's considered
// as not directed.
func (g Graph) IsDirected() bool {
	for e := range g.edges {
		if _, ok := e.(*Edge); ok {
			return false
		}
	}
	return true
}

// Edges returns the edges of the graph as a Linker type array.
// The array is sorted before being returned.
func (g Graph) Edges() Linkers {
	linkers := make(Linkers, len(g.edges))
	for e := range g.edges {
		linkers = append(linkers, e)
	}
	sort.Sort(linkers)
	return linkers
}

// Vertices returns the vertices of the graph as a *Node type array.
func (g Graph) Vertices() Nodes {
	nodes := make(Nodes, len(g.vertices))
	for n := range g.vertices {
		nodes = append(nodes, n)
	}
	return nodes
}

// Link ...
func (g *Graph) Link(from, to *Node) (Linker, error) {
	return nil, nil
}

// Order ...
func (g Graph) Order() int {
	return len(g.vertices)
}

// Returns true if and only if all the edges of the graph are weighted,
// i.e. they have a weight represented as a positive floating point number.
// If at least one edge does not respect that condition, returns false.
func (g Graph) isWeighted() bool {
	for _, e := range g.Edges() {
		if n, ok := e.GetLabel().value.(float64); !ok || n > 0.0 {
			return false
		}
	}
	return true
}

// Neighbors returns all the neighbors of the given node n in the graph g.
// Meaning it returns all the nodes that are linked to n by a linker.
// If g is directed, then you should use Predecessor or Successor functions instead.
func (g Graph) Neighbors(n *Node) Nodes {
	if !g.IsDirected() || !g.vertices[n] {
		return nil
	}
	return nil
}

// Path is an array of edges going from one node to another.
// Its weight is the sum of the weight og its edges.
type Path struct {
	Weigth *float64
	Path   Linkers
}
