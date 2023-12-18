package medium

import (
	"fmt"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/solutions/

// method 1
// 1) use three for loop to scan the nums
// 2) use leftMaxList to store the max value from 0 to i-1
// 3) use rightMaxList to store the max value from n-1 to i+1
// 4) use result to store the max value of (leftMaxList[i]-nums[i])*rightMaxList[i]
// TC = O(3N) ≈ O(N), SC = O(2N) ≈ O(N)
func maximumTripletValue1(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	leftMaxList := make([]int, n)
	leftMaxList[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMaxList[i] = util.Max(leftMaxList[i-1], nums[i-1])
	}

	rightMaxList := make([]int, n)
	rightMaxList[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMaxList[i] = util.Max(rightMaxList[i+1], nums[i+1])
	}

	result := int64(0)
	for i := 1; i < n-1; i++ {
		// * leftMaxList[i] store the max value before nums[i]
		// * rightMaxList[i] store the max value after nums[i]
		if nums[i] < leftMaxList[i] {
			result = util.Max(result, int64((leftMaxList[i]-nums[i])*rightMaxList[i]))
		}
	}

	return result
}

func Test_maximumTripletValue1(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result int64
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
				nums: []int{12, 6, 1, 2, 7},
			},
			expected: expected{
				result: 77,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 10, 3, 4, 19},
			},
			expected: expected{
				result: 133,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3},
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
			maximumTripletValue1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
