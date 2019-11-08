package models

// Rule is the representation of a transition rule the automata uses to
// go from one state to another. It's a generic type composing of three fields:
// `Source` is the state the automata must be to use the rule;
// `Destination` is the state we reach when the source and the symbol are the ones provided;
// `Symbol` is the symbol we should read when automata is in the state Source to reach the
// Destination state.
type Rule struct {
	Source      *State
	Destination *State
	Symbol      Symbol
}

// Symbol is a model composed of an `Object`, which can be any type as long as the automata
// manages it. It's what the automata reads to go from one state to another.
type Symbol struct {
	Object interface{}
}
