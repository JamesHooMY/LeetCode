package medium

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	commonUtil "leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/task-scheduler/

type MaxHeap []int // store frequency of each task

func (h *MaxHeap) Len() int { return len(*h) }

// max heap
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 1 use slice like hash table to store frequency of each task and turn it to a max heap
// 1) create a slice like hash table to store frequency of each task
// 2) turn the slice to a max heap
// 3) pop the max heap until it is empty
// TC: O(N), SC: O(1), N is the length of tasks
func leastInterval1(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	pq := &MaxHeap{}
	for _, f := range freq {
		if f > 0 {
			*pq = append(*pq, f)
		}
	}
	heap.Init(pq) // TC: O(N), SC: O(1)

	time := 0
	// TC: O(NlogM), SC: O(1), M is the total types of tasks
	// * but here the max of M is 26, which is a constant, so we can estimate TC: O(N)
	for pq.Len() > 0 {
		// k is the number of tasks can be executed in one interval
		k := n + 1

		tmp := []int{}
		for k > 0 && pq.Len() > 0 {
			f := heap.Pop(pq).(int)
			if f > 1 {
				tmp = append(tmp, f-1)
			}
			k--
			time++
		}

		// push back the tasks which can not be executed in this interval
		for _, f := range tmp {
			heap.Push(pq, f)
		}

		// if pq is empty, it means we have finished all tasks
		if pq.Len() > 0 {
			// the remain k is the idle time that we don't have enough tasks to execute in this interval
			time += k
		}
	}

	return time
}

// method 2 use slice like hash table to store frequency of each task and calculate the idle time
// 1) create a slice like hash table to store frequency of each task
// 2) sort the slice in descending order
// 3) calculate the idle time
// TC: O(N), SC: O(1), N is the length of tasks
func leastInterval2(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	sort.Slice(freq, func(i, j int) bool { return freq[i] > freq[j] })

	maxFreq := freq[0]
	maxFreqInterval := maxFreq - 1  // the interval between maxFreq tasks
	idleTime := maxFreqInterval * n // the idle time in the interval between maxFreq tasks
	for i := 1; i < len(freq) && idleTime > 0; i++ {
		// maxFreqInterval is the interval between maxFreq tasks, each interval can execute every different task once
		idleTime -= commonUtil.Min(maxFreqInterval, freq[i])
	}

	if idleTime < 0 {
		idleTime = 0
	}

	return len(tasks) + idleTime
}

// method 3 combine method 1 and method 2
// 1) create a slice like hash table to store frequency of each task
// 2) turn the slice to a max heap as priority queue
// 3) pop the max heap until it is empty, and calculate the idle time
// TC: O(N), SC: O(1), N is the length of tasks
// * this is the best solution for me currently
func leastInterval3(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	pq := &MaxHeap{}
	for _, f := range freq {
		if f > 0 {
			*pq = append(*pq, f)
		}
	}
	heap.Init(pq) // TC: O(N), SC: O(1)

	maxFreq := heap.Pop(pq).(int)
	maxFreqInterval := maxFreq - 1  // the interval between maxFreq tasks
	idleTime := maxFreqInterval * n // the idle time in the interval between maxFreq tasks
	for pq.Len() > 0 && idleTime > 0 {
		// maxFreqInterval is the interval between maxFreq tasks, each interval can execute every different task once
		idleTime -= commonUtil.Min(maxFreqInterval, heap.Pop(pq).(int))
	}

	if idleTime < 0 {
		idleTime = 0
	}

	return len(tasks) + idleTime
}

func Test_leastInterval1(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
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
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			expected: expected{
				result: 8,
			},
		},
		{
			name: "2",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     0,
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     2,
			},
			expected: expected{
				result: 16,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			leastInterval1(tc.args.tasks, tc.args.n),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_leastInterval2(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
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
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			expected: expected{
				result: 8,
			},
		},
		{
			name: "2",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     0,
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     2,
			},
			expected: expected{
				result: 16,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			leastInterval2(tc.args.tasks, tc.args.n),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_leastInterval3(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
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
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			expected: expected{
				result: 8,
			},
		},
		{
			name: "2",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     0,
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     2,
			},
			expected: expected{
				result: 16,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			leastInterval3(tc.args.tasks, tc.args.n),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_leastInterval1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leastInterval1([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	}
}

func Benchmark_leastInterval2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leastInterval2([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	}
}

func Benchmark_leastInterval3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leastInterval3([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	}
}
