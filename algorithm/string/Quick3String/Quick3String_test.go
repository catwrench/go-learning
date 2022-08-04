package Quick3String

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuick3StringSort(t *testing.T) {
	data := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"sea",
		"shore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
	}
	Quick3StringSort(data)
	expect := []string{
		"are",
		"by",
		"sea",
		"seashells",
		"seashells",
		"sells",
		"sells",
		"she",
		"she",
		"shells",
		"shore",
		"surely",
		"the",
		"the",
	}
	assert.Equal(t, expect, data)
}
