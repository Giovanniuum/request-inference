package automata

import (
	"github.com/Giovanniuum/request-inference/models"
)

// FiniteAutomata is an abstract model representing a finite automata mathematical structure.
// It's nothing more than a quintuplet.
// For more theorical information about the automata models, see in French :
// https://fr.wikipedia.org/wiki/Automate_fini
type FiniteAutomata struct {
	Q            map[*models.State]bool
	Q0           *models.State
	Qf           map[*models.State]bool
	E            map[models.Symbol]bool
	R            map[*models.Rule]bool
	CurrentState *models.State
}

// NewFiniteAutomata inits a new FiniteWordAutomata with the correct default values.
func NewFiniteAutomata(q0 *models.State) *FiniteAutomata {
	if q0 == nil {
		q0 = models.NewState(nil)
	}
	return &FiniteAutomata{
		Q:            map[*models.State]bool{q0: true},
		Q0:           q0,
		Qf:           map[*models.State]bool{},
		E:            map[models.Symbol]bool{},
		R:            map[*models.Rule]bool{},
		CurrentState: q0,
	}
}

// AddFinalState adds a new final state to the final states set of the automata.
// Also adds it to the set of states if it's not present already.
func (fa *FiniteAutomata) AddFinalState(fs *models.State) {
	if !fa.Qf[fs] {
		fa.Qf[fs] = true
	}
	if !fa.Q[fs] {
		fa.Q[fs] = true
	}
}

// Reset makes the given automata a brand new-like created `FiniteAutomata`.
func (fa *FiniteAutomata) Reset(q0 *models.State) {
	fa = NewFiniteAutomata(q0)
}

// Restart makes the current state of the automata being its starting state.
func (fa *FiniteAutomata) Restart() {
	fa.CurrentState = fa.Q0
}

// ReachNextState searches in automata set of rules if one fits the given
// symbol. Depending on automata current state and the provided symbol, it can
// reach or not a new state, which can eventually be the same state.
// Returns true when it finds a rule fitting the current configuration, then
// updates automata current state.
// Else returns false, meaning automata didn't change its state because it didn't
// find a rule fitting the needs.
func (fa *FiniteAutomata) reachNextState(s models.Symbol) bool {
	if !fa.E[s] {
		return false
	}
	for r := range fa.R {
		if r.Source == fa.CurrentState && r.Symbol == s {
			fa.CurrentState = r.Destination
			return true
		}
	}
	return false
}

// RecognizePattern tries, with the current configuration of the automata,
// to reach a final state by following the given pattern.
// It progresses rule by rule until it reaches the end of the pattern, or if the
// `reachNextState` function returned false, meaning it found a not managed
// configuration. Returns true if it ended up on a final state, else returns false.
func (fa *FiniteAutomata) RecognizePattern(p string) bool {
	if fa.CurrentState != fa.Q0 {
		fa.Restart()
	}
	defer fa.Restart()
	for i := 0; i < len(p); i++ {
		if !fa.reachNextState(models.Symbol{Object: p[i : i+1]}) {
			return false
		}
	}
	return fa.Qf[fa.CurrentState]
}
