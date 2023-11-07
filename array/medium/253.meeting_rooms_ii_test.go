package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/tzh_linux/article/details/103821483

// method 1
// 1) sort the intervals by start time
// 2) use minHeap to store the end time of each meeting
// 3) if the start time of current meeting is bigger than the minHeap top, then pop the minHeap top
// 4) push the end time of current meeting to minHeap
// 5) sort the minHeap by end time
// TC = O(NlogN), SC = O(logN)
// * this is the best solution for me currently
func minMeetingRooms1(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// sort the intervals by start time
	// TC = O(NlogN), SC = O(logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minHeap := []int{intervals[0][1]} // store the end time of each meeting

	for i := 1; i < len(intervals); i++ {
		// if the start time of current meeting is bigger than the minHeap top, then pop the minHeap top
		if intervals[i][0] >= minHeap[0] {
			minHeap = minHeap[1:]
		}

		// push the end time of current meeting to minHeap
		minHeap = append(minHeap, intervals[i][1])

		// sort the minHeap by end time
		// * this is not the ordinary heap sort, but this is easy to implement
		// TC = O(NlogN), SC = O(logN)
		sort.Slice(minHeap, func(i, j int) bool {
			return minHeap[i] < minHeap[j]
		})
	}

	return len(minHeap)
}

func Test_minMeetingRooms1(t *testing.T) {
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
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				minMeetingRooms1(tc.args.intervals),
			)
		})
	}
}
