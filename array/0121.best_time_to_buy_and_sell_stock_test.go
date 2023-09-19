package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// method 1 Brute Force
// 1) use two for loop, and profit store the max profit
// TC = O(N^2), SC = (O)1
func maxProfit1(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}

	return maxProfit
}

// method 2 One pass
// 1) use one for loop, and minPrice store the min price, maxProfit store (prices[i] - minPrice)
// 2) during the for loop scanning process, check each "prices[i]" and make sure whether it is less than minPrice, if it is less than minPrice, update minPrice to prices[i]
// 3) during the for loop scanning process, check each "prices[i]", if it is greater than minPrice, and (prices[i] - minPrice) is greater than maxProfit, update maxProfit to (prices[i] - minPrice)
// TC = O(N), SC = (O)1
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue // skip the following code, directly turn into next i++ process
		}

		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func Test_maxProfit1(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "3",
			args: args{
				prices: []int{1, 2},
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
			maxProfit1(tc.args.prices),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "3",
			args: args{
				prices: []int{1, 2},
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
			maxProfit2(tc.args.prices),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
