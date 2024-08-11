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
func firstBadVersion1(n int) int {
	left, right := 1, n

	firstBadVersion := n
	for left <= right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid - 1
			firstBadVersion = mid
		} else {
			left = mid + 1
		}
	}

	return firstBadVersion
}

func isBadVersion(version int) bool {
	badVersion := []int{1, 4, 5}
	for _, v := range badVersion {
		if version == v {
			return true
		}
	}

	return false
}

func Test_firstBadVersion1(t *testing.T) {
	type args struct {
		n int
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
			name: "test 1",
			args: args{
				n: 5,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "test 2",
			args: args{
				n: 1,
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected.result, firstBadVersion1(tc.args.n))
		})
	}
}
