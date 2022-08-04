package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func input() (data []string) {
	data = []string{"1,一", "3,三", "5,五", "6,六", "4,四", "2,二", "0,零", "8,八", "9,九", "7,七"}
	return
}

func TestLookUpIndex(t *testing.T) {
	data := input()
	i := NewLookUpIndex(len(data))
	i.Add(data, ",")
	assert.Equal(t, "0", i.GetKey("零"))
	assert.Equal(t, "零", i.GetVal("0"))
	assert.Equal(t, "9", i.GetKey("九"))
	assert.Equal(t, "九", i.GetVal("9"))
}
