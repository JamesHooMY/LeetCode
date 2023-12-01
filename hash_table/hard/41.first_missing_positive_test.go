package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/first-missing-positive/description/

// method 1 hash table
// 1) use a hash table to store the positive numbers
// 2) use a for loop to scan the numbers
// 3) if the number is positive, then add the number to the hash table
// 4) use a for loop to find the first missing positive number
// 5) return the first missing positive number
// TC = O(N), SC = O(N)
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}

	// use a hash table to store the positive numbers
	hashTable := make(map[int]struct{})

	// 1st iteration collect the positive integer map
	for _, v := range nums {
		// if the number is positive, then add the number to the hash table
		if v > 0 {
			hashTable[v] = struct{}{} // key: number, value: struct{}
		}
	}

	// 2nd iteration make sure the missing lowest positive integer
	for i := 1; i <= n; i++ {
		if _, ok := hashTable[i]; !ok {
			return i
		}
	}

	return n + 1
}

// method 2 cyclic sort
// 1) use cyclic sort to put the positive number to the right place
// 2) find the first missing positive number
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}

	// 1st iteration put the positive number to the right place
	for i := 0; i < n; i++ {
		// * for loop until the number is in the right place
		// * nums[i] != i+1, make sure the number is not in the right place
		// * nums[i] != nums[nums[i]-1], make sure the numbers are same, reduce same number swap
		for nums[i] > 0 && nums[i] <= n && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			// swap nums[i] and nums[nums[i]-1]
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 2nd iteration find the first missing positive number
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func Test_firstMissingPositive1(t *testing.T) {
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
				nums: []int{1, 2, 0},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			firstMissingPositive1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_firstMissingPositive2(t *testing.T) {
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
				nums: []int{1, 2, 0},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			firstMissingPositive2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
