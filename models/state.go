package models

var nextStateID int64 // package incrementation

func init() {
	nextStateID = 0
}

// State is an abstract state an automata can be on. It's just a basic ID,
// plus we add its ancestor for tracking purpose on some algorithms needing it.
type State struct {
	ID       int64
	Ancestor *State
}

// NewState creates a new `State`, by setting its father and incrementing the ID.
func NewState(father *State) *State {
	nextStateID++
	return &State{ID: nextStateID, Ancestor: father}
}

// ResetStates resets the current ID for the next created `State`.
func ResetStates() {
	nextStateID = 0
}
