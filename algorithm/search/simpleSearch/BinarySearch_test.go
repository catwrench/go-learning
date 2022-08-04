package simpleSearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 6, 8, 10, 12, 13}

	tests := []struct {
		name   string
		key    int
		expect int
	}{
		{name: "0", key: 0, expect: 0},
		{name: "1", key: 100, expect: -1},
		{name: "2", key: 2, expect: 2},
		{name: "3", key: 13, expect: 9},
		{name: "4", key: 6, expect: 5},
	}
	for _, test := range tests {
		t.Run(test.name, func(k *testing.T) {
			assert.Equal(k, test.expect, BinarySearch(arr, test.key))
		})
	}
}
