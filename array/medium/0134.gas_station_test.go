package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/gas-station/

// method 1 greedy algorithm
// 1) use one for loop, to scan the gas
// 2) use two variable, currentGas and startStation
// 3) if currentGas < 0, then we need to change the startStation to the next station
// 4) finally, if totalGas < totalCost, then we can't complete the circuit
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func canCompleteCircuit1(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0

	currentGas := 0
	startStation := 0

	// the greedy algorithm is to find out the posible startStation, and the currentGas >= 0
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]

		// if currentGas < 0, then we need to change the startStation to the next station
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			startStation = i + 1
			currentGas = 0
		}
	}

	// if totalGas < totalCost, then we can't complete the circuit
	if totalGas < totalCost {
		return -1
	}

	return startStation
}

func Test_canCompleteCircuit1(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
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
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			expected: expected{
				result: -1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			canCompleteCircuit1(tc.args.gas, tc.args.cost),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
