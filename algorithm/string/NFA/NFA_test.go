package NFA

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNFA(t *testing.T) {
	nfa := NewNFA("(A*B|AC)D")
	assert.True(t, nfa.Recognizes("AAAABD"))

	nfa = NewNFA("(A*B|AC)D")
	assert.False(t, nfa.Recognizes("AAAAC"))

	nfa = NewNFA("(a|(bc)*d)*")
	assert.True(t, nfa.Recognizes("abcbcd"))

	nfa = NewNFA("(a|(bc)*d)*")
	assert.True(t, nfa.Recognizes("abcbcbcdaaaabcbcdaaaddd"))
}
