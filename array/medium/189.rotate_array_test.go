package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rotate-array/description/

// method 1 use reverse function to rotate the nums, this is more easy to understand
// 1) use reverse to rotate the nums
// 2) reverse the numbers from 0 to k-1
// 3) reverse the numbers from k to n-1
// TC = O(2N) ≈ O(N), SC = O(1)
func rotate1(nums []int, k int) {
	n := len(nums)

	// reverse all the nums
	reverse(nums, 0, n-1)

	// reverse the numbers from 0 to k-1
	reverse(nums, 0, k-1)

	// reverse the numbers from k to n-1
	reverse(nums, k, n-1)
}

func reverse(nums []int, startIndex, endIndex int) {
	for startIndex < endIndex {
		nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
		startIndex++
		endIndex--
	}
}

// method 2 use extra array to store the result
// 1) use extra array to store the result
// 2) the new index of nums[i] is (i+k)%len(nums)
// 3) copy the newArray to nums
// TC = O(N), SC = O(N)
func rotate2(nums []int, k int) {
	newArray := make([]int, len(nums))
	n := len(nums)

	for i := range nums {
		newArray[(i+k)%n] = nums[i] // (i+k)%len(nums) is the new index of nums[i]
	}

	// nums = newArray // * this is not working, because the nums is pass by value, so we need to use copy function to copy the newArray to nums
	copy(nums, newArray)
}

// method 3 use cyclic replacement https://zhuanlan.zhihu.com/p/412049933
// 1) use cyclic replacement to rotate the nums
// 2) use count to count the number of the replaced numbers, if count == n, then we can stop the loop
// 3) the new index of nums[i] is (i+k)%len(nums)
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func rotate3(nums []int, k int) {
	n := len(nums)

	count := 0 // use count to count the number of the replaced numbers, if count == n, then we can stop the loop
	for startIndex := 0; count < n; startIndex++ {
		currentIndex := startIndex
		currentValue := nums[currentIndex]

		// this loop will not cause TC = O(N^2), because the currentIndex will be the startIndex, so the loop will be break
		for {
			nextIndex := (currentIndex + k) % n
			nextValue := nums[nextIndex]

			nums[nextIndex] = currentValue

			currentIndex = nextIndex
			currentValue = nextValue

			count++

			// if startIndex == currentIndex, then we need to break the loop, because this means the current index is the start index
			if startIndex == currentIndex {
				break
			}
		}
	}
}

func Test_rotate1(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: expected{
				result: []int{3, 99, -1, -100},
			},
		},
	}

	for _, tc := range testCases {
		rotate1(tc.args.nums, tc.args.k)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_rotate2(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: expected{
				result: []int{3, 99, -1, -100},
			},
		},
	}

	for _, tc := range testCases {
		rotate2(tc.args.nums, tc.args.k)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_rotate3(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: expected{
				result: []int{3, 99, -1, -100},
			},
		},
	}

	for _, tc := range testCases {
		rotate3(tc.args.nums, tc.args.k)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
