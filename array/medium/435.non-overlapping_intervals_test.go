package medium

import (
	"sort"
	"testing"

	"leetcode/array/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/non-overlapping-intervals/

// method 1
// 1) sort the intervals by start time
// 2) compare the end of current interval with the start of next interval
// 3) if overlap, then count++, and update the end to the min end from comparing current end and next end
// 4) if not overlap, then update the end to the end of next interval
// TC = O(NlogN), SC = O(logN)
// * this is the best solution for me currently
func eraseOverlapIntervals1(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	// TC = O(NlogN), SC = O(logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
			// * this is the key point, if overlap, then we should keep the min end, this can keep as much as possible intervals without overlap during the whole process
			end = util.Min(end, intervals[i][1])
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

func Test_eraseOverlapIntervals1(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	type expected struct {
		result int
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
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{1, 3},
				},
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{
					{1, 2},
					{1, 2},
					{1, 2},
				},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "4",
			args: args{
				intervals: [][]int{
					{1, 100},
					{11, 22},
					{1, 11},
					{2, 12},
				},
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				eraseOverlapIntervals1(tc.args.intervals),
			)
		})
	}
}
