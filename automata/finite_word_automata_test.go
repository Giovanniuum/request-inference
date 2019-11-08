package automata_test

import (
	"os"
	"testing"

	"github.com/Giovanniuum/request-inference/automata"
	"github.com/Giovanniuum/request-inference/models"
	"github.com/stretchr/testify/assert"
)

var (
	fwa        *automata.FiniteWordAutomata
	a, b       models.Symbol
	q0, q1, q2 *models.State
)

func TestMain(m *testing.M) {
	q0 = models.NewState(nil)
	fwa = automata.NewFiniteWordAutomata(q0)
	a = models.Symbol{Object: "a"}
	b = models.Symbol{Object: "b"}
	q1 = models.NewState(q0)
	q2 = models.NewState(q1)
	os.Exit(m.Run())
}
func TestWordAutomata_RecognizePattern_FinishingWithB(t *testing.T) {
	defer fwa.Reset(q0)
	fwa.AddRule(&models.Rule{Source: q0, Destination: q0, Symbol: a})
	fwa.AddRule(&models.Rule{Source: q0, Destination: q1, Symbol: b})
	fwa.AddRule(&models.Rule{Source: q1, Destination: q0, Symbol: a})
	fwa.AddRule(&models.Rule{Source: q1, Destination: q1, Symbol: b})
	fwa.AddFinalState(q1)
	assert.True(t, fwa.RecognizePattern("aaaaab"))
	assert.False(t, fwa.RecognizePattern("aaaaaba"))
	assert.True(t, fwa.RecognizePattern("bbb"))
}

func TestWordAutomata_RecognizePattern_FinishingWithAB(t *testing.T) {
	defer fwa.Reset(q0)
	fwa.AddRule(&models.Rule{Source: q0, Destination: q1, Symbol: a})
	fwa.AddRule(&models.Rule{Source: q0, Destination: q0, Symbol: b})
	fwa.AddRule(&models.Rule{Source: q1, Destination: q0, Symbol: a})
	fwa.AddRule(&models.Rule{Source: q1, Destination: q2, Symbol: b})
	fwa.AddRule(&models.Rule{Source: q2, Destination: q1, Symbol: a})
	fwa.AddRule(&models.Rule{Source: q2, Destination: q0, Symbol: b})
	fwa.AddFinalState(q2)
	assert.True(t, fwa.RecognizePattern("aaaaab"))
	assert.False(t, fwa.RecognizePattern("aaaaaba"))
	assert.False(t, fwa.RecognizePattern("bbb"))
}
