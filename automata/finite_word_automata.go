package automata

import (
	"github.com/Giovanniuum/request-inference/data-structures/basics"
)

// FiniteWordAutomata is a specific finite automata used to recognize word patterns only.
// Its symbols must be characters - a string of length 1 doesn't matter the value - and
// the pattern it recognizes must be strings.
type FiniteWordAutomata struct {
	*FiniteAutomata
}

// NewFiniteWordAutomata initializes a new FiniteWordAutomata, which is globally nothing
// more than a `FiniteAutomata`.
func NewFiniteWordAutomata(q0 *basics.State) *FiniteWordAutomata {
	return &FiniteWordAutomata{FiniteAutomata: NewFiniteAutomata(q0)}
}

// AddRule adds a new rule to the automata. Rule's symbol must be a character to be valid for
// this kind of automata. Returns true if the rule was added to the set of rules,
// or false if it already exists or the rule is invalid.
func (fwa *FiniteWordAutomata) AddRule(r *basics.Rule) bool {
	if fwa.R[r] {
		return false
	}
	if s, ok := r.Symbol.Object.(string); !ok {
		return false
	} else if len(s) == 0 || len(s) > 1 {
		return false
	}
	fwa.R[r] = true
	if !fwa.E[r.Symbol] {
		fwa.E[r.Symbol] = true
	}
	if !fwa.Q[r.Source] {
		fwa.Q[r.Source] = true
	}
	if !fwa.Q[r.Destination] {
		fwa.Q[r.Destination] = true
	}
	return true
}

// Reset makes the given automata a brand new-like created `FiniteWordAutomata`
// by resetting its embedded `FiniteAutomata`field.
func (fwa *FiniteWordAutomata) Reset(q0 *basics.State) {
	fwa.FiniteAutomata = NewFiniteAutomata(q0)
}

// BuildAutomataFromPositiveExemples reads a list of word considered as positives exemples and build
// a finite automata that recognizes exactly the given set of strings. If a string is provided several
// times it'll only be checked once.
func BuildAutomataFromPositiveExemples(words []string) *FiniteWordAutomata {
	q0 := basics.NewState(nil)
	fta := NewFiniteWordAutomata(q0)
	for _, w := range words {
		qf := fta.buildForWord(w, fta.Q0)
		fta.AddFinalState(qf)
	}
	return fta
}

func (fwa *FiniteWordAutomata) buildForWord(word string, source *basics.State) *basics.State {
	if len(word) == 0 { // Reached the end of the word, then `source` is a final state recognizing the word
		return source
	}
	prefix, suffix := word[0:1], word[1:]

	var ancestor *basics.State // TODO: make it usefull when using tree automata?
	if source.Ancestor == nil {
		ancestor = source
	} else {
		ancestor = source.Ancestor
	}
	rule := &basics.Rule{
		Source:      source,
		Symbol:      basics.Symbol{Object: prefix},
		Destination: basics.NewState(ancestor),
	}
	fwa.AddRule(rule)
	return fwa.buildForWord(suffix, rule.Destination)
}
