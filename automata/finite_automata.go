package automata

import (
	"sort"
	"strings"

	"github.com/Giovanniuum/request-inference/data-structures/basics"
)

// FiniteAutomata is an abstract model representing a finite automata mathematical structure.
// It's nothing more than a quintuplet.
// For more theorical information about the automata basics, see in French :
// https://fr.wikipedia.org/wiki/Automate_fini
type FiniteAutomata struct {
	Q            map[*basics.State]bool
	Q0           *basics.State
	Qf           map[*basics.State]bool
	E            map[basics.Symbol]bool
	R            map[*basics.Rule]bool
	CurrentState *basics.State
}

// NewFiniteAutomata inits a new FiniteWordAutomata with the correct default values.
func NewFiniteAutomata(q0 *basics.State) *FiniteAutomata {
	if q0 == nil {
		q0 = basics.NewState(nil)
	}
	return &FiniteAutomata{
		Q:            map[*basics.State]bool{q0: true},
		Q0:           q0,
		Qf:           map[*basics.State]bool{},
		E:            map[basics.Symbol]bool{},
		R:            map[*basics.Rule]bool{},
		CurrentState: q0,
	}
}

// Rules is the reprensentation of the rules of the automata as a sorted array.
type Rules []*basics.Rule

// Len implements `sort.Interface` interface to be able to sort Rules.
func (r Rules) Len() int {
	return len(r)
}

// Less implements `sort.Interface` interface to be able to sort Rules.
func (r Rules) Less(i, j int) bool {
	return r[i].String() <= r[j].String()
}

// Swap implements `sort.Interface` interface to be able to sort Rules.
func (r Rules) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Rules returns the rules of the automata as a rule array.
func (fa *FiniteAutomata) Rules() Rules {
	rules := Rules{}
	for r := range fa.R {
		rules = append(rules, r)
	}
	sort.Sort(rules)
	return rules
}

// AddFinalState adds a new final state to the final states set of the automata.
// Also adds it to the set of states if it's not present already.
func (fa *FiniteAutomata) AddFinalState(fs *basics.State) {
	if !fa.Qf[fs] {
		fa.Qf[fs] = true
	}
	if !fa.Q[fs] {
		fa.Q[fs] = true
	}
}

// Reset makes the given automata a brand new-like created `FiniteAutomata`.
func (fa *FiniteAutomata) Reset(q0 *basics.State) {
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
func (fa *FiniteAutomata) reachNextState(s basics.Symbol) bool {
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

// GetRule searches in automata rules if one with source and symbol already exists, and returns it.
// Else it returns nil. Mainly used for determinization of automata.
func (fa *FiniteAutomata) GetRule(source *basics.State, symbol basics.Symbol) *basics.Rule {
	for _, r := range fa.Rules() {
		if r.Source == source && r.Symbol.Object == symbol.Object {
			return r
		}
	}
	return nil
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
		if !fa.reachNextState(basics.Symbol{Object: p[i : i+1]}) {
			return false
		}
	}
	return fa.Qf[fa.CurrentState]
}

func (fa *FiniteAutomata) String() string {
	var sb strings.Builder
	for _, r := range fa.Rules() {
		sb.WriteString(r.String())
	}
	for q := range fa.Q {
		sb.WriteString(q.String())
	}
	return sb.String()
}
