package automata_test

import (
	"testing"

	"github.com/Giovanniuum/request-inference/automata"
	"github.com/Giovanniuum/request-inference/data-structures/basics"
	"github.com/stretchr/testify/assert"
)

var (
	fwa        *automata.FiniteWordAutomata
	a, b       basics.Symbol
	q0, q1, q2 *basics.State
)

func InitTests() {
	q0 = basics.NewState(nil)
	fwa = automata.NewFiniteWordAutomata(q0)
	a = basics.Symbol{Object: "a"}
	b = basics.Symbol{Object: "b"}
	q1 = basics.NewState(q0)
	q2 = basics.NewState(q1)
}
func TestWordAutomata_RecognizePattern_FinishingWithB(t *testing.T) {
	InitTests()
	defer fwa.Reset(q0)
	fwa.AddRule(&basics.Rule{Source: q0, Destination: q0, Symbol: a})
	fwa.AddRule(&basics.Rule{Source: q0, Destination: q1, Symbol: b})
	fwa.AddRule(&basics.Rule{Source: q1, Destination: q0, Symbol: a})
	fwa.AddRule(&basics.Rule{Source: q1, Destination: q1, Symbol: b})
	fwa.AddFinalState(q1)
	assert.True(t, fwa.RecognizePattern("aaaaab"))
	assert.False(t, fwa.RecognizePattern("aaaaaba"))
	assert.True(t, fwa.RecognizePattern("bbb"))
}

func TestWordAutomata_RecognizePattern_FinishingWithAB(t *testing.T) {
	InitTests()
	defer fwa.Reset(q0)
	fwa.AddRule(&basics.Rule{Source: q0, Destination: q1, Symbol: a})
	fwa.AddRule(&basics.Rule{Source: q0, Destination: q0, Symbol: b})
	fwa.AddRule(&basics.Rule{Source: q1, Destination: q0, Symbol: a})
	fwa.AddRule(&basics.Rule{Source: q1, Destination: q2, Symbol: b})
	fwa.AddRule(&basics.Rule{Source: q2, Destination: q1, Symbol: a})
	fwa.AddRule(&basics.Rule{Source: q2, Destination: q0, Symbol: b})
	fwa.AddFinalState(q2)
	assert.True(t, fwa.RecognizePattern("aaaaab"))
	assert.False(t, fwa.RecognizePattern("aaaaaba"))
	assert.False(t, fwa.RecognizePattern("bbb"))
}

func TestWordAutomata_BuildAutomataFromWord(t *testing.T) {
	samples := []string{"hello", "world", "how", "are", "you"}
	fwa := automata.BuildAutomataFromPositiveExemples(samples)
	for _, s := range samples {
		assert.True(t, fwa.RecognizePattern(s))
	}
}
