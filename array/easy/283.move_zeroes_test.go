package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/move-zeroes/description/

// method 1 brute force
// 1) use two for loop, first for loop scan each num in nums, if found num is 0, then use second for loop to swap the zero num to the last
// TC = O(N^2), SC = O(1)
func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// swap the zero num to the last
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

// method 2 one pointer
// 1) use one pointer zeroIndex to record the index of zero
// 2) first for loop, if nums[i] != 0, then swap nums[i] and nums[zeroIndex], and zeroIndex++
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func moveZeroes2(nums []int) {
	zeroIndex := 0

	// i must start from 0, because we need to check each num in nums
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}

func Test_moveZeroes1(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{0, 1, 0, 3, 12},
			},
			expected: expected{
				result: []int{1, 3, 12, 0, 0},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0},
			},
			expected: expected{
				result: []int{0},
			},
		},
	}

	for _, tc := range testCases {
		moveZeroes1(tc.args.nums)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_moveZeroes2(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{0, 1, 0, 3, 12},
			},
			expected: expected{
				result: []int{1, 3, 12, 0, 0},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0},
			},
			expected: expected{
				result: []int{0},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 0, 3, 0, 4, 0},
			},
			expected: expected{
				result: []int{1, 3, 4, 0, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		moveZeroes2(tc.args.nums)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
