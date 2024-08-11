package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-search/description/

// method 1 binary search
// 1) set left and right pointers
// 2) iterate while left <= right
// 3) calculate mid
// 4) if nums[mid] == target, return mid
// 5) if nums[mid] < target, move left pointer to mid+1
// 6) if nums[mid] > target, move right pointer to mid-1
// 7) return -1
// TC: O(logN), SC: O(1)
// * this is the best solution for me currently
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1

	/*
		left <= right is the key point

		if left == right, mid = left = right, and nums[mid] == target, return mid

		example: nums = [1], target = 1
	*/
	for left <= right {
		mid := left + (right-left)/2 // * this is the best way to avoid overflow

		// find the target
		if nums[mid] == target {
			return mid
		}

		// move left or right pointer
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Test_search1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	type expected struct {
		result int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			expected: expected{
				result: -1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			search1(tc.args.nums, tc.args.target),
			tc.name,
		)
	}
}
