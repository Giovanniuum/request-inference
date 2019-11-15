package basics

// Term is our custom way to implement a simple tree structure.
// It's just composed of a root, which is a `Symbol` instance,
// and an array of children trees, which can be empty when the current `Term`
// is a leaf of the general tree.
type Term struct {
	Root     Symbol
	Children []*Term
}

// NewTerm creates a new `Term` and initializes its fields.
func NewTerm(r Symbol) *Term {
	return &Term{
		Root:     r,
		Children: []*Term{},
	}
}

// AppendChild adds an already existing `Term` ptr to the list of children of
// the given `Term`.
func (t *Term) AppendChild(c *Term) {
	t.Children = append(t.Children, c)
}
