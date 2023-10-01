package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum/description/

// method 1 use backtracking (DFS), scan all the combination, brute force
// 1) use backtracking (DFS) to find the combination
// 2) base case is remain < 0
// 3) convergence is remain-candidates[i]
// k = target / min(candidates), n = len(candidates)
// TC = O(k * n^k) ==> O(n^k)
// SC = O(k) ==> O(n)
// * this is the best solution for me currently
func combinationSum1(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}

	backTrack(&result, current, candidates, target, 0)

	return result
}

func backTrack(result *[][]int, current []int, candidates []int, remain int, startIndex int) {
	if remain < 0 { // base case: remain < 0
		return
	} else if remain == 0 {
		// this copy is essential, because the value of result is slice, slice is pass by reference
		tmp := make([]int, len(current))
		copy(tmp, current) // this copy like deep clone in javascript

		*result = append(*result, tmp) // use pointer slice for operating the same result slice to be append during the process
	} else {
		// DFS thinking
		/*
			subject: [2,3,6,7]
			current: [2] --> [2,2] --> [2,2,2] --> [2,2,2,2] --> [2,2,2] --> [2,2,2,3] --> [2,2,2] --> [2,2,2,6] --> [2,2,2] --> [2,2,2,7] --> [2,2,2]
			current: [2,2,3], the following numbers of 3 will be return because remain < 0
			current: [2,3] --> [2,3,3], the following numbers of 3 will be return because remain < 0
			current: [2,6], the following numbers of 6 will be return because remain < 0
			current: [3] --> [3,3] --> [3,3,3], the following numbers of 3 will be return because remain < 0
			current: [3] --> [3,6], the following numbers of 6 will be return because remain < 0
			current: [6] --> [6,6], the following numbers of 6 will be return because remain < 0
			current: [7]
		*/
		for i := startIndex; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backTrack(result, current, candidates, remain-candidates[i], i) // convergence: remain-candidates[i]
			current = current[:len(current)-1]
		}
	}
}

func Test_combinationSum1(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	type expected struct {
		result [][]int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			expected: expected{
				result: [][]int{{2, 2, 3}, {7}},
			},
		},
		{
			name: "2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			expected: expected{
				result: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
			},
		},
		{
			name: "3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			expected: expected{
				result: [][]int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			combinationSum1(tc.args.candidates, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
