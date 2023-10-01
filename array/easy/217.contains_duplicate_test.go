package easy

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate/description/

// method 1
// use one for loop, one map
// 1) first for loop, use map to store key with each num and and value with initialize struct, https://geektutu.com/post/hpg-empty-struct.html
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func containsDuplicate1(nums []int) bool {
	numMap := map[int]struct{}{}

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return true
		}

		numMap[num] = struct{}{}
	}

	return false
}

// method 2
// sort the nums, and use one for loop and one tmp variable
// 1) use sort.Ints(nums) to sort the nums, TC = O(N*logN), SC = O(logN), https://blog.csdn.net/nieling3/article/details/118102350
// 2) use one for loop to check each num, tmp variable store the previous num, if nums[i] == tmp, return true
// TC = O(N*logN), SC = O(logN)
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums) // TC = O(N*logN), SC = O(logN)

	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == tmp {
			return true
		}

		tmp = nums[i]
	}

	return false
}

func Test_containsDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result bool
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
				nums: []int{1, 2, 3, 1},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			containsDuplicate1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_containsDuplicate2(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result bool
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
				nums: []int{1, 2, 3, 1},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			containsDuplicate2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
