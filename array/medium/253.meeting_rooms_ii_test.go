package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/tzh_linux/article/details/103821483

// method 1 minHeap
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

		/*
			|-------------------|
			a		            b
				|------|
				c	   d
						    |------|
						    e	   f

			intervals: [[a, b], [c, d], [e, f]]
			minHeap: [b]

			1) c < b, push d to minHeap: [b, d], sort minHeap: [d, b]
			2) e > d, pop d from minHeap: [b], push f to minHeap: [b, f], sort minHeap: [b, f]
			3) i == len(intervals), return len(minHeap) = 2
		*/

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

// method 2 array
// 1) sort the intervals by start time
// 2) use two arrays to store the start time and end time of each meeting
// 3) use two pointers to iterate the start time and end time arrays
// 4) if the start time of current meeting is bigger than the end time of current meeting, then pop the end time of current meeting
// 5) push the end time of current meeting to the end time array
// TC = O(NlogN), SC = O(N)
func minMeetingRooms2(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	startTimes := make([]int, 0, len(intervals))
	endTimes := make([]int, 0, len(intervals))

	for _, interval := range intervals {
		startTimes = append(startTimes, interval[0])
		endTimes = append(endTimes, interval[1])
	}

	// sort the start time and end time arrays
	// TC = O(NlogN), SC = O(N)
	sort.Slice(startTimes, func(i, j int) bool {
		return startTimes[i] < startTimes[j]
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i] < endTimes[j]
	})

	startIdx, endIdx := 0, 0
	rooms := 0

	for startIdx < len(startTimes) {
		/*
			|-------------------|
			a		            b
				|------|
				c	   d
						    |------|
						    e	   f

			startTimes: [a, c, e]
			endTimes:   [d, b, f]
			rooms: 0

			1) a < d, rooms++ then rooms = 1, startIdx++ then startIdx = 1
			2) c < d, rooms++ then rooms = 2, startIdx++ then startIdx = 2
			3) e > d, endIdx++ then endIdx = 1, rooms = 2, startIdx++ then startIdx = 3
			4) startIdx == len(startTimes), return rooms = 2

		*/
		if startTimes[startIdx] >= endTimes[endIdx] {
			endIdx++
		} else {
			// if startTimes[startIdx] < endTimes[endIdx], then the meetings are overlapped, then we need a new room
			rooms++
		}

		startIdx++
	}

	return rooms
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

func Test_minMeetingRooms2(t *testing.T) {
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
				minMeetingRooms2(tc.args.intervals),
			)
		})
	}
}

// benchmark
func Benchmark_minMeetingRooms1(b *testing.B) {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	for i := 0; i < b.N; i++ {
		minMeetingRooms1(intervals)
	}
}

func Benchmark_minMeetingRooms2(b *testing.B) {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	for i := 0; i < b.N; i++ {
		minMeetingRooms2(intervals)
	}
}