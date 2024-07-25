package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

// method 1 two pointers
// 1) use two pointers i and j
// 2) i is the slow pointer, j is the fast pointer
// 3) if nums[i] == nums[j], move j to the next
// 4) if nums[i] != nums[j], move i to the next, and assign nums[j] to nums[i]
// 5) return i+1
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}

	i := 0
	for j := 1; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func Test_removeDuplicates1(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result int
		nums   []int
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
				nums: []int{1, 1, 2},
			},
			expected: expected{
				result: 2,
				nums:   []int{1, 2, 2},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			expected: expected{
				result: 5,
				nums:   []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1, 1},
			},
			expected: expected{
				result: 1,
				nums:   []int{1, 1, 1},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 2, 3},
			},
			expected: expected{
				result: 3,
				nums:   []int{1, 2, 3},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, removeDuplicates1(tc.args.nums), tc.name)
	}
}