package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	params
	answer
}

type params struct {
	nums   []int
	target int
}

type answer struct {
	result []int
}

var qs = []question{
	{
		params{[]int{2, 7, 11, 15}, 9},
		answer{[]int{0, 1}},
	},
	{
		params{[]int{3, 2, 4}, 6},
		answer{[]int{1, 2}},
	},
	{
		params{[]int{3, 3}, 6},
		answer{[]int{0, 1}},
	},
}

func Test_twoSum1(t *testing.T) {
	for _, q := range qs {
		a, p := q.answer, q.params
		assert.Equal(t, twoSum1(p.nums, p.target), a.result)
	}
}

func Test_twoSum2(t *testing.T) {
	for _, q := range qs {
		a, p := q.answer, q.params
		assert.Equal(t, twoSum2(p.nums, p.target), a.result)
	}
}
