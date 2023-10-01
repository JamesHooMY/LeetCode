package medium

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-intervals/description/

// method 1
// 1) sort the intervals by the start value
// 2) use one for loop, to scan the intervals
// 3) 1st, overlap condition of sorted array: currentInterval[0] <= previousInterval[1]
// TC = O(NlogN), SC = O(N)
// * this is the best solution for me currently
func merge1(intervals [][]int) [][]int {
	// TC = O(NlogN), SC = O(logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// SC = O(N)
	mergedIntervals := [][]int{intervals[0]}

	// TC = O(N)
	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		previousInterval := mergedIntervals[len(mergedIntervals)-1]

		if currentInterval[0] <= previousInterval[1] {
			mergedIntervals[len(mergedIntervals)-1][0] = previousInterval[0]
			mergedIntervals[len(mergedIntervals)-1][1] = currentInterval[1]
		} else {
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}

	return mergedIntervals
}

func Test_merge1(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			expected: expected{
				result: [][]int{{1, 6}, {8, 10}, {15, 18}},
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			expected: expected{
				result: [][]int{{1, 5}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			merge1(tc.args.intervals),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
