package medium

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum-closest/description/

// method 1 prefix sum
// 1) sort the array
// 2) use two pointer (leftIdx, rightIdx) to scan the array
// 3) use the closestSum to store the closest sum of three numbers
// 4) compare the current sum with closestSum, if the current sum is closer than closestSum, then update the closestSum
// TC = O(N^2), SC = O(logN)
// * this is the best solution for me currently
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums) // TC = O(NlogN), SC = O(logN)

	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		leftIdx := i + 1
		rightIdx := len(nums) - 1

		for leftIdx < rightIdx {
			curSum := nums[i] + nums[leftIdx] + nums[rightIdx]

			if curSum == target {
				return curSum
			}

			if abs(curSum-target) < abs(closestSum-target) {
				closestSum = curSum
			}

			if curSum < target {
				leftIdx++
			} else {
				rightIdx--
			}
		}
	}

	return closestSum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func Test_threeSumClosest1(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
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
			threeSumClosest1(tc.args.nums, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
