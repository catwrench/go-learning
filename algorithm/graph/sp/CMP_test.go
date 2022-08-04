package sp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func inputJobs() (count int, jobs []Jobs) {
	jobs = make([]Jobs, 0)
	jobs = []Jobs{
		{Weight: 41.0, Next: []int{1, 7, 9}},
		{Weight: 51.0, Next: []int{2}},
		{Weight: 50.0, Next: []int{}},
		{Weight: 36.0, Next: []int{}},
		{Weight: 38.0, Next: []int{}},
		{Weight: 45.0, Next: []int{}},
		{Weight: 21.0, Next: []int{3, 8}},
		{Weight: 32.0, Next: []int{3, 8}},
		{Weight: 32.0, Next: []int{2}},
		{Weight: 29.0, Next: []int{4, 6}},
	}
	return 10, jobs
}

func TestNewCMP(t *testing.T) {
	tests := []struct {
		name   string
		key    int
		expect float64
	}{
		{name: "0", key: 0, expect: 0},
		{name: "1", key: 1, expect: 41.0},
		{name: "2", key: 2, expect: 123.0},
		{name: "3", key: 3, expect: 91.0},
		{name: "4", key: 4, expect: 70.0},
		{name: "5", key: 5, expect: 0.0},
		{name: "6", key: 6, expect: 70.0},
		{name: "7", key: 7, expect: 41.0},
		{name: "8", key: 8, expect: 91.0},
		{name: "9", key: 9, expect: 41.0},
	}

	_, jobs := inputJobs()
	cmp := NewCMP(jobs)

	startTimes := cmp.StartTimes()
	for _, test := range tests {
		t.Run(test.name, func(k *testing.T) {
			assert.Equal(k, test.expect, startTimes[test.key])
		})
	}

	finishTime := cmp.FinishTime()
	assert.Equal(t, 173.0, finishTime)
}
