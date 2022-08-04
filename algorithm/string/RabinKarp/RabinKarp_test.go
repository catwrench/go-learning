package RabinKarp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRabinKarp(t *testing.T) {
	bm := NewRabinKarp("abracadabra")
	i := bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 14, i)

	bm = NewRabinKarp("rab")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 8, i)

	bm = NewRabinKarp("bcara")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, len("abacadabrabracabracadabrabrabracad"), i)

	bm = NewRabinKarp("rabrabracad")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, len("abacadabrabracabracadabrabrabracad")-len("rabrabracad"), i)

	bm = NewRabinKarp("abacad")
	i = bm.Search("abacadabrabracabracadabrabrabracad")
	assert.Equal(t, 0, i)
}
