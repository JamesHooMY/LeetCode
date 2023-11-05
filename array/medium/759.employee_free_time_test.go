package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/badguy_gao/article/details/86985347

// method 1
// 1) combine all intervals into one array
// 2) sort the intervals by start time
// 3) compare the end time of current interval with the start time of next interval, if overlap, then update the end time to the max end time
// 4) if not overlap, then add a new interval with the end time of current interval and the start time of next interval into the result
// TC = O(NlogN), SC = O(N)
// * this is the best solution for me currently
func employeeFreeTime1(intervals [][][]int) [][]int {
	combinedIntervals := [][]int{}
	// TC = O(N*M), SC = O(N*M), N is the number of intervals, M is the number of intervals in each interval
	for _, interval := range intervals {
		for _, i := range interval {
			combinedIntervals = append(combinedIntervals, i)
		}
	}

	// TC = O(NlogN), SC = O(logN)
	sort.Slice(combinedIntervals, func(i, j int) bool {
		return combinedIntervals[i][0] < combinedIntervals[j][0]
	})

	result := [][]int{}
	end := combinedIntervals[0][1]
	// TC = O(N), SC = O(N)
	for i := 1; i < len(combinedIntervals); i++ {
		if combinedIntervals[i][0] <= end {
			// * this is the key point, if overlap, then we should keep the max end, this is different from 435.non-overlapping_intervals_test.go
			end = max(end, combinedIntervals[i][1])
		} else {
			// * this is the key point, if not overlap, then we should add the free time which is the interval between end and the start of next interval into the result
			result = append(result, []int{end, combinedIntervals[i][0]})
			end = combinedIntervals[i][1]
		}
	}

	return result
}

// method 2
// 1) combine all intervals into one array
// 2) sort the intervals by start time
// 3) merge the intervals
// 4) add the free time into the result
// TC = O(NlogN), SC = O(N)
// this solution is think by myself, but it is not the best solution, one more step to merge the intervals then add the free time into the result
func employeeFreeTime2(intervals [][][]int) [][]int {
	combinedIntervals := [][]int{}
	// TC = O(N*M), SC = O(N*M), N is the number of intervals, M is the number of intervals in each interval
	for _, interval := range intervals {
		for _, i := range interval {
			combinedIntervals = append(combinedIntervals, i)
		}
	}

	// TC = O(NlogN), SC = O(logN)
	sort.Slice(combinedIntervals, func(i, j int) bool {
		return combinedIntervals[i][0] < combinedIntervals[j][0]
	})

	mergedIntervals := [][]int{combinedIntervals[0]}
	// TC = O(N), SC = O(N)
	for i := 1; i < len(combinedIntervals); i++ {
		currentInterval := combinedIntervals[i]
		previousInterval := mergedIntervals[len(mergedIntervals)-1]

		if currentInterval[0] <= previousInterval[1] {
			mergedIntervals[len(mergedIntervals)-1][0] = previousInterval[0]
			mergedIntervals[len(mergedIntervals)-1][1] = currentInterval[1]
		} else {
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}

	// TC = O(N), SC = O(N)
	result := [][]int{}
	for i := 1; i < len(mergedIntervals); i++ {
		result = append(result, []int{mergedIntervals[i-1][1], mergedIntervals[i][0]})
	}
	return result
}

func Test_employeeFreeTime1(t *testing.T) {
	type args struct {
		intervals [][][]int
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
				intervals: [][][]int{
					{{1, 2}, {5, 6}},
					{{1, 3}},
					{{4, 10}},
				},
			},
			expected: expected{
				result: [][]int{{3, 4}},
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {6, 7}},
					{{2, 4}},
					{{2, 5}, {9, 12}},
				},
			},
			expected: expected{
				result: [][]int{{5, 6}, {7, 9}},
			},
		},
		{
			name: "3",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {9, 12}},
					{{2, 4}},
					{{6, 8}},
				},
			},
			expected: expected{
				result: [][]int{{4, 6}, {8, 9}},
			},
		},
		{
			name: "4",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {9, 12}},
					{{2, 4}},
					{{6, 8}},
					{{5, 7}},
				},
			},
			expected: expected{
				result: [][]int{{4, 5}, {8, 9}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				employeeFreeTime1(tc.args.intervals),
			)
		})
	}
}

func Test_employeeFreeTime2(t *testing.T) {
	type args struct {
		intervals [][][]int
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
				intervals: [][][]int{
					{{1, 2}, {5, 6}},
					{{1, 3}},
					{{4, 10}},
				},
			},
			expected: expected{
				result: [][]int{{3, 4}},
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {6, 7}},
					{{2, 4}},
					{{2, 5}, {9, 12}},
				},
			},
			expected: expected{
				result: [][]int{{5, 6}, {7, 9}},
			},
		},
		{
			name: "3",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {9, 12}},
					{{2, 4}},
					{{6, 8}},
				},
			},
			expected: expected{
				result: [][]int{{4, 6}, {8, 9}},
			},
		},
		{
			name: "4",
			args: args{
				intervals: [][][]int{
					{{1, 3}, {9, 12}},
					{{2, 4}},
					{{6, 8}},
					{{5, 7}},
				},
			},
			expected: expected{
				result: [][]int{{4, 5}, {8, 9}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				employeeFreeTime2(tc.args.intervals),
			)
		})
	}
}
