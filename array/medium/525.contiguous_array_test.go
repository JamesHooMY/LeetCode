package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contiguous-array/description/

// method 1 prefix sum
// 1) replace 0 to -1
// 2) if sum == 0, then the maxLen will be i+1, because the sum from 0 to i is 0
// 2) find the same prefix sum; if not exist, then put the sum to sumMap; the sumMap only store the first same sum
// 3) maxLen = i - index in sumMap
// TC = O(N), SC = O(N)
func findMaxLength1(nums []int) int {
	maxLen := 0
	sum := 0
	sumMap := map[int]int{} // key: sum, value: index

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}

		// if the first same sum is 0, then the maxLen will be i+1
		if sum == 0 {
			maxLen = i + 1
		}

		if index, exist := sumMap[sum]; exist {
			if maxLen < i-index {
				maxLen = i - index
			}
		} else {
			sumMap[sum] = i
		}
	}

	return maxLen
}

// method 2 prefix sum https://www.youtube.com/watch?v=uYuSLvjEyjQ
// 1) initialize sumMap with initial value {0: -1}
// logic thinking is same as method 1, but this method is more easy to understand for me
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func findMaxLength2(nums []int) int {
	maxLen := 0
	sum := 0

	// initial {0: -1} due to the first same sum is 0, then the maxLen will be i+1
	// for example, nums = [0, 1], then sum = 0, maxLen = i - (-1) = 2
	sumMap := map[int]int{
		0: -1, // this initial value is necessary !!!
	} // key: sum, value: index

	/*
		example: nums = [0, 1, 0, 1]

		i = 0, sum = -1, maxLen = 0, sumMap = {0: -1, -1: 0}
		i = 1, sum = 0, maxLen = 2, sumMap = {0: -1, -1: 0}
		i = 2, sum = -1, maxLen = 2, sumMap = {0: -1, -1: 0}
		i = 3, sum = 0, maxLen = 4, sumMap = {0: -1, -1: 0}
	*/
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}

		if index, exist := sumMap[sum]; exist {
			if maxLen < i-index {
				maxLen = i - index
			}
		} else {
			sumMap[sum] = i
		}
	}

	return maxLen
}

func Test_findMaxLength1(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{0, 1},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 1, 0},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 1, 0, 1},
			},
			expected: expected{
				result: 4,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findMaxLength1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_findMaxLength2(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{0, 1},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 1, 0},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 1, 0, 1},
			},
			expected: expected{
				result: 4,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findMaxLength2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
