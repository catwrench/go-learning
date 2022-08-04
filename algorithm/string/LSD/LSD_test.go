package LSD

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLSD(t *testing.T) {
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

	LSD(arr, 7)

	expect := []string{
		"1ICK750",
		"1ICK750",
		"1OHV845",
		"1OHV845",
		"1OHV845",
		"2IYE230",
		"2RLA629",
		"2RLA629",
		"3ATW723",
		"3CIO720",
		"3CIO720",
		"4JZY524",
		"4PGC938",
	}
	assert.Equal(t, expect, arr)
}
