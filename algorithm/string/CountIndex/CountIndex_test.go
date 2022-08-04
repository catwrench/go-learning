package CountIndex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIndex(t *testing.T) {
	arr := []string{
		"4PGC938",
		"2IYE230",
		"3CIO720",
		"1ICK750",
		"1OHV845",
		"4JZY524",
		"1ICK750",
		"3CIO720",
		"1OHV845",
		"1OHV845",
		"2RLA629",
		"2RLA629",
		"3ATW723",
	}

	CountIndex(arr)

	expect := []string{
		"1ICK750",
		"1OHV845",
		"1ICK750",
		"1OHV845",
		"1OHV845",
		"2IYE230",
		"2RLA629",
		"2RLA629",
		"3CIO720",
		"3CIO720",
		"3ATW723",
		"4PGC938",
		"4JZY524",
	}
	assert.Equal(t, expect, arr)
}
