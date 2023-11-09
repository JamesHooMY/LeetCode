package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subarray-sum-equals-k/description/

// method 1 brute force
// 1) calculate the sum of every subarray of each nums[i]
// 2) if sum == k, then count++
// TC = O(N^2), SC = O(1)
func subarraySum1(nums []int, k int) int {
	n := len(nums)
	count := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// method 2 prefix sum https://yuminlee2.medium.com/leetcode-560-subarray-sum-equals-k-9eb688e43534
// 1) initialize sumCountMap with initial value {0: 1}, this is very important !!!
// 2) calculate the prefix sum of nums, and store the sum to sumCountMap
// 3) if sum-k exist in sumCountMap, then count += sumCountMap[sum-k]
// 4) sumCountMap[sum]++ is very important !!!
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	count := 0
	sum := 0
	sumCountMap := map[int]int{
		0: 1, // this initial value is necessary !!!
	} // key: sum, value: count

	/*
		example: nums = [1, 1, 1], k = 2

		i = 0, sum = 1, count = 0, sumCountMap = {0: 1, 1: 1}
		i = 1, sum = 2, count = 1, sumCountMap = {0: 1, 1: 1, 2: 1}
		i = 2, sum = 3, count = 2, sumCountMap = {0: 1, 1: 1, 2: 1, 3: 1}
	*/
	for i := 0; i < n; i++ {
		sum += nums[i]

		if value, exist := sumCountMap[sum-k]; exist {
			count += value
		}

		sumCountMap[sum]++
	}

	return count
}

func Test_subarraySum1(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			subarraySum1(tc.args.nums, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_subarraySum2(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			subarraySum2(tc.args.nums, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
