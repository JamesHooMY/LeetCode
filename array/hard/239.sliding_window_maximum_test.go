package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sliding-window-maximum/description/

// method 1 (brute force)
// 1) use a loop to iterate the nums
// 2) use a loop to iterate the k elements after the current element
// 3) find the max element in the k elements
// 4) add the max element into the result
// TC = O(N*k), SC = O(N)
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	} else if k == 1 {
		return nums
	}

	result := []int{}
	// TC = O(N), SC = O(N)
	for i := 0; i < n-k+1; i++ {
		max := nums[i]
		// TC = O(k), SC = O(1)
		for j := i + 1; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
	}

	return result
}

// method 2 (deque)
// 1) use a deque to store the index of the elements
// 2) use a loop to iterate the nums
// 3) if the deque is not empty and the index of the first element in the deque is out of the window, then remove the first element
// 4) if the deque is not empty and the value of the last element in the deque is less than the current element, then remove the last element
// 5) add the index of the current element into the deque
// 6) if the index of the first element in the deque is in the window, then add the value of the first element into the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	} else if k == 1 {
		return nums
	}

	// store the max value of each window during the iteration
	result := make([]int, 0, n-k+1)
	/*
		result := make([]int, n-k+1)
		resultIndex := 0
	*/

	// store the index of the max value of each window during the iteration
	deque := []int{}

	// TC = O(N), SC = O(N)
	for i := 0; i < n; i++ {
		// make sure the deque[0] whether it was out of current window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// make sure to eliminate the items in deque are lower than nums[i], start from last item, keep index of the deque[0] is the index of highest value in current window
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// add current index into deque
		deque = append(deque, i)

		 // make sure the reading numbers already fit the window size for comparison
		if i >= k-1 {
			result = append(result, nums[deque[0]])
			/*
				result[resultIndex] = nums[deque[0]]
				resultIndex++
			*/
		}
	}

	return result
}

func Test_maxSlidingWindow1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type expected struct {
		result []int
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
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{3, 3, 5, 5, 6, 7},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    3,
			},
			expected: expected{
				result: []int{3, 3, 2, 5},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    1,
			},
			expected: expected{
				result: []int{1, 3, 1, 2, 0, 5},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				maxSlidingWindow1(tc.args.nums, tc.args.k),
			)
		})
	}
}

func Test_maxSlidingWindow2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type expected struct {
		result []int
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
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{3, 3, 5, 5, 6, 7},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    3,
			},
			expected: expected{
				result: []int{3, 3, 2, 5},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    1,
			},
			expected: expected{
				result: []int{1, 3, 1, 2, 0, 5},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				maxSlidingWindow2(tc.args.nums, tc.args.k),
			)
		})
	}
}

func Benchmark_maxSlidingWindow1(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	for i := 0; i < b.N; i++ {
		maxSlidingWindow1(nums, k)
	}
}

func Benchmark_maxSlidingWindow2(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	for i := 0; i < b.N; i++ {
		maxSlidingWindow2(nums, k)
	}
}
