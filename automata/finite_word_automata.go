package automata

import (
	"github.com/Giovanniuum/request-inference/models"
)

// FiniteWordAutomata is a specific finite automata used to recognize word patterns only.
// Its symbols must be characters - a string of length 1 doesn't matter the value - and
// the pattern it recognizes must be strings.
type FiniteWordAutomata struct {
	*FiniteAutomata
}

// NewFiniteWordAutomata initializes a new FiniteWordAutomata, which is globally nothing
// more than a `FiniteAutomata`.
func NewFiniteWordAutomata() *FiniteWordAutomata {
	return &FiniteWordAutomata{FiniteAutomata: NewFiniteAutomata()}
}

// AddRule adds a new rule to the automata. Rule's symbol must be a character to be valid for
// this kind of automata. Returns true if the rule was added to the set of rules,
// or false if it already exists or the rule is invalid.
func (fwa *FiniteWordAutomata) AddRule(r *models.Rule) bool {
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
