package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/search-insert-position/

// method 1 binary search
// 1) set left and right pointers
// 2) iterate while left <= right
// 3) calculate mid
// 4) if nums[mid] == target, return mid
// 5) if nums[mid] < target, move left pointer to mid+1
// 6) if nums[mid] > target, move right pointer to mid-1
// 7) return left
// TC: O(logN), SC: O(1)
// * this is the best solution for me currently
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	/*
		[1, 3, 5, 6], target = 2
		left = 0, right = 3, mid = 1, nums[mid] = 3, nums[mid] > target, right = mid-1 = 0
		left = 0, right = 0, mid = 0, nums[mid] = 1, nums[mid] < target, left = mid+1 = 1
		left = 1, right = 0, return left = 1

		[1, 3, 5, 6], target = 7
		left = 0, right = 3, mid = 1, nums[mid] = 3, nums[mid] < target, left = mid+1 = 2
		left = 2, right = 3, mid = 2, nums[mid] = 5, nums[mid] < target, left = mid+1 = 3
		left = 3, right = 3, mid = 3, nums[mid] = 6, nums[mid] < target, left = mid+1 = 4
		left = 4, right = 3, return left = 44

		* left always points to the first element that is greater than or equal to target !!!
		right always points to the last element that is less than target
	*/

	return left
}

func Test_searchInsert1(t *testing.T) {
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
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "4",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "5",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			expected: expected{
				result: 0,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			searchInsert1(tc.args.nums, tc.args.target),
			tc.name,
		)
	}
}
