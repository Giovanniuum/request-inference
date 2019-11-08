package automata_test

import (
	"os"
	"testing"

	"github.com/Giovanniuum/request-inference/automata"
	"github.com/Giovanniuum/request-inference/models"
	"github.com/stretchr/testify/assert"
)

var (
	fa         *automata.FiniteWordAutomata
	a, b       models.Symbol
	q0, q1, q2 *models.State
)

func TestMain(m *testing.M) {
	fa = automata.NewFiniteWordAutomata()
	a = models.Symbol{Object: "a"}
	b = models.Symbol{Object: "b"}
	q0 = fa.Q0
	q1 = models.NewState(q0)
	q2 = models.NewState(q1)
	os.Exit(m.Run())
}
func TestWordAutomata_RecognizePattern_FinishingWithB(t *testing.T) {
	defer fa.Reset()
	fa.AddRule(&models.Rule{Source: q0, Destination: q0, Symbol: a})
	fa.AddRule(&models.Rule{Source: q0, Destination: q1, Symbol: b})
	fa.AddRule(&models.Rule{Source: q1, Destination: q0, Symbol: a})
	fa.AddRule(&models.Rule{Source: q1, Destination: q1, Symbol: b})
	fa.AddFinalState(q1)
	assert.True(t, fa.RecognizePattern("aaaaab"))
	assert.False(t, fa.RecognizePattern("aaaaaba"))
	assert.True(t, fa.RecognizePattern("bbb"))
}

func TestWordAutomata_RecognizePattern_FinishingWithAB(t *testing.T) {
	defer fa.Reset()
	fa.AddRule(&models.Rule{Source: q0, Destination: q1, Symbol: a})
	fa.AddRule(&models.Rule{Source: q0, Destination: q0, Symbol: b})
	fa.AddRule(&models.Rule{Source: q1, Destination: q0, Symbol: a})
	fa.AddRule(&models.Rule{Source: q1, Destination: q2, Symbol: b})
	fa.AddRule(&models.Rule{Source: q2, Destination: q1, Symbol: a})
	fa.AddRule(&models.Rule{Source: q2, Destination: q0, Symbol: b})
	fa.AddFinalState(q2)
	assert.True(t, fa.RecognizePattern("aaaaab"))
	assert.False(t, fa.RecognizePattern("aaaaaba"))
	assert.False(t, fa.RecognizePattern("bbb"))
}
