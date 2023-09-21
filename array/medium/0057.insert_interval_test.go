package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/insert-interval/description/

// method 1
// 1) use a "new array" to store the new intervals list
// 2) use one for loop, to scan the intervals
// 3) 1st, if the intervals[i][1] < newIntervals[0], mean that the "max" of the intervals[i] is lower than the newIntervals min value, then append intervals[i] into the "new array".
// 4) 2nd, if the intervals[i][0] < newIntervals[1], mean that the intervals[i] and newIntervals are "overlaping", so we need to update the newIntervals min and max. Comparing intervals[i][0] and newIntervals[0] to get the min value; intervals[i][1] and newIntervals[1] to get the max value.
// 5) 3th, append the newIntervals to the "new array"
// 6) finally, the following range in intervals, intervals[i][0] lower value is bigger than the newIntervals[1] max value, we can directly append them to the "new array"
// TC = O(N), SC = O(N)
func insert(intervals [][]int, newInterval []int) [][]int {
    newIntervalsList := [][]int{}

    i, n := 0, len(intervals)
    // 1st, find the ranges in intervals, which max value are lower than the min value of the newIntervals
    for i < n && intervals[i][1] < newInterval[0] {
        newIntervalsList = append(newIntervalsList, intervals[i])
        i++
    }

    // 2nd, find "overlapping" between the range of intervals with newIntervals, then update newIntervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }

    // 3th
    newIntervalsList = append(newIntervalsList, newInterval)

    // finally
    newIntervalsList = append(newIntervalsList, intervals[i:]...)
    // for i < n {
    //     newIntervalsList = append(newIntervalsList, intervals[i])
    //     i++
    // }


    return newIntervalsList
}

func min(a int, b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
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
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			expected: expected{
				result: [][]int{{1, 5}, {6, 9}},
			},
		},
		{
			name: "2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			expected: expected{
				result: [][]int{{1, 2}, {3, 10}, {12, 16}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			insert(tc.args.intervals, tc.args.newInterval),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
