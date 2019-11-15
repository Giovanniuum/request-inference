package graph

import (
	"fmt"
)

// Label is an information we can put on the edges of the graph.
type Label struct{ value interface{} }

// NewLabel creates a new Label instance. For now it can only
// be string or float.
func NewLabel(v interface{}) (*Label, error) {
	switch t := v.(type) {
	case string, float64:
		return &Label{value: v}, nil
	default:
		return nil, fmt.Errorf("a label can only be of type string or float64, but got : %v", t)
	}
}

// Linker ...
type Linker interface {
	FromNode() *Node
	ToNode() *Node
	Label(interface{}) (*Label, error)
	GetLabel() *Label
}

// Linkers ...
type Linkers []Linker

// Len implements `sort.Interface` interface to be able to sort Linkers.
func (l Linkers) Len() int {
	return len(l)
}

// Less implements `sort.Interface` interface to be able to sort Linkers.
func (l Linkers) Less(i, j int) bool {
	return true
}

// Swap implements `sort.Interface` interface to be able to sort Linkers.
func (l Linkers) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Edge is a simple representation of an edge in an undirected graph.
// Meaning, it has no direction as it's not an arrow.
type Edge struct {
	x, y  *Node
	label *Label
}

// FromNode returns either of the nodes at the extremities of the edge, as
// it should not be directed.
func (e *Edge) FromNode() *Node {
	return e.x
}

// ToNode returns either of the nodes at the extremities of the edge, as
// it should not be directed.
func (e *Edge) ToNode() *Node {
	return e.y
}

// Label tries to label an edge with the given value.
// It can only be labeled with string or float64 values.
func (e *Edge) Label(v interface{}) (*Label, error) {
	l, err := NewLabel(v)
	if err != nil {
		return nil, err
	}
	e.label = l
	return l, nil
}

// GetLabel ...
func (e *Edge) GetLabel() *Label {
	return e.label
}

// Arrow is the structure representing an edge in a directed graph.
// Meaning it has a direction, from X to Y.
type Arrow struct {
	x, y  *Node
	label *Label
}

// FromNode returns the X node of the arrow, as
// it should be directed, and point from X to Y.
func (a *Arrow) FromNode() *Node {
	return a.x
}

// ToNode returns the Y node of the arrow, as
// it should be directed, and point from X to Y.
func (a *Arrow) ToNode() *Node {
	return a.y
}

// Label tries to label an edge with the given value.
// It can only be labeled with string or float64 values.
func (a *Arrow) Label(v interface{}) (*Label, error) {
	l, err := NewLabel(v)
	if err != nil {
		return nil, err
	}
	a.label = l
	return l, nil
}

// GetLabel ...
func (a *Arrow) GetLabel() *Label {
	return a.label
}
