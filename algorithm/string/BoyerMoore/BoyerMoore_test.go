package BoyerMoore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoyerMoore(t *testing.T) {
	bm := NewBoyerMoore("abracadabra")
	i := bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 14, i)

	bm = NewBoyerMoore("rab")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 8, i)

	bm = NewBoyerMoore("bcara")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, len("abacadabrabracabracadabrabrabracad"), i)

	bm = NewBoyerMoore("rabrabracad")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, len("abacadabrabracabracadabrabrabracad")-len("rabrabracad"), i)

	bm = NewBoyerMoore("abacad")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 0, i)
}
