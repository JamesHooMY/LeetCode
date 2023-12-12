package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-operations-to-collect-elements/submissions/1063523384/

// method 1
// 1) use one for loop and one map
// 2) numMap is used to store the number which is less or equal to k, the length of numMap is used to check whether we have found all the numbers which is less or equal to k
// 3) count is used to count the number of operations
// TC = O(N), SC = O(N)
// * this method is the best solution for me currently
func minOperations1(nums []int, k int) int {
    count := 0
    numMap := map[int]bool{}

    for i:=len(nums)-1; i>=0; i-- {
        if nums[i] <= k && len(numMap) <= k && !numMap[nums[i]] {
            numMap[nums[i]] = true
        }
        count++

        if len(numMap) == k {
            break
        }
    }

    return count
}

func Test_minOperations1(t *testing.T) {
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
				nums: []int{3, 1, 5, 4, 2},
				k:    2,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 1, 5, 4, 2},
				k:    5,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 2, 5, 3, 1},
				k:    3,
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
			minOperations1(tc.args.nums, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
