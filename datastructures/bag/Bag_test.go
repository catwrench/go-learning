package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag(t *testing.T) {
	input := []string{"1", "3", "4", "2", "8", "9", "6", "0", "5", "7"}
	b := NewBag[string]()
	for i := range input {
		b.Add(input[i])
	}

	str := ""
	it := b.NewIterator()
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "7506982431", str)

	str = ""
	b.Del()
	it = b.NewIterator()
	for it.HasNext() {
		str += it.Next()
	}
	assert.Equal(t, "506982431", str)
}
