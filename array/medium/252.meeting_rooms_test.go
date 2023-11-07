package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/qq_29051413/article/details/108709370

// method 1
// 1) sort the intervals by start time
// 2) compare the end time of current meeting with the start time of next meeting
// TC = O(NlogN), SC = O(logN)
// * this is the best solution for me currently
func canAttendMeetings1(intervals [][]int) bool {
	if len(intervals) == 0 || len(intervals) == 1 {
		return true
	}

	// TC = O(NlogN), SC = O(logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// * endTime can be remove, because we can use intervals[i-1][1] to replace it
	/*
		endTime := intervals[0][1]
		for i := 1; i < len(intervals); i++ {
			if intervals[i][0] < endTime {
				return false
			}
			endTime = intervals[i][1]
		}
	*/

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}

func Test_canAttendMeetings1(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	type expected struct {
		result bool
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
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{
					{13, 15},
					{1, 13},
				},
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				canAttendMeetings1(tc.args.intervals),
			)
		})
	}
}
