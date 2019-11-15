package graph

import (
	"github.com/Giovanniuum/request-inference/data-structures/basics"
)

// Signature is the signature of a node, which can be positive
// or negative.
type Signature int

const (
	// Positive is a positive node
	Positive Signature = iota
	// Negative is a negative node
	Negative
)

// Node is the type we use in our graph to implement the needed algorithms.
// It's composed of a value, which is of `basics.Symbol` type, and a signature
// which is used in some learning algorithms, like RPNI.
type Node struct {
	value     basics.Symbol
	signature Signature
}

// NewNode creates a new ptr to `Node` given a symbol instance.
// By default a node is positive.
func NewNode(v basics.Symbol) *Node {
	return &Node{value: v, signature: Positive}
}

// NewNodeWithSignature creates a new ptr to `Node` given a symbol instance,
// and a bool representing its signing.
func NewNodeWithSignature(v basics.Symbol, s Signature) *Node {
	return &Node{value: v, signature: s}
}

// Nodes is an array of ptr to `Node`, mainly used to list all nodes of a graph
// as an array instead of a map.
type Nodes []*Node
