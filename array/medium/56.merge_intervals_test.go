package medium

import (
	"fmt"
	"sort"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-intervals/description/

// method 1
// 1) sort the intervals by the start value
// 2) use one for loop, to scan the intervals
// 3) 1st, overlap condition of sorted array: currentInterval[0] <= previousInterval[1]
// TC = O(NlogN), SC = O(N)
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

// method 2 stack easy to understand
// 1) sort the intervals by the start value
// 2) use one for loop, to scan the intervals
// 3) overlap condition of sorted array: currentInterval[0] <= topInterval[1]
// TC = O(NlogN), SC = O(N)
// * this is the best solution for me currently
func merge2(intervals [][]int) [][]int {
	// TC = O(NlogN), SC = O(logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// SC = O(N)
	stackIntervals := [][]int{intervals[0]}

	// TC = O(N)
	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		topInterval := stackIntervals[len(stackIntervals)-1]

		if currentInterval[0] <= topInterval[1] {
			topInterval[0] = util.Min(topInterval[0], currentInterval[0])
			topInterval[1] = util.Max(topInterval[1], currentInterval[1])

			// pop the top of stack
			stackIntervals = stackIntervals[:len(stackIntervals)-1]

			// push the new top of stack
			stackIntervals = append(stackIntervals, topInterval)

			continue
		}

		stackIntervals = append(stackIntervals, currentInterval)
	}

	return stackIntervals
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

func Test_merge2(t *testing.T) {
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
			merge2(tc.args.intervals),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_merge1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge1([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}

func Benchmark_merge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge2([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}