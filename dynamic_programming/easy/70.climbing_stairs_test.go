package easy

import (
	"testing"
)

// method 1 recursion, DFS bottom-up
// 1) use recursion to calculate the number of ways to climb to the top
// 2) if n == 0 || n == 1 || n == 2, return n
// 3) return climbStairs1(n-1) + climbStairs1(n-2)
// TC = O(2^N), SC = O(N)
func climbStairs1(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

// method 2 dynamic programming
// 1) use an array to store the number of ways to climb to the top
// 2) if n == 0 || n == 1 || n == 2, return n
// 3) use a for loop to calculate the number of ways to climb to the top
// 4) return dp[n]
// TC = O(N), SC = O(N)
func climbStairs2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// method 3 dynamic programming with constant space (named as rolling array or rolling variables)
// 1) use two variables a and b to store the number of ways to climb to the top
// 2) if n == 0 || n == 1 || n == 2, return n
// 3) use a for loop to calculate the number of ways to climb to the top
// 4) return b
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func climbStairs3(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func Test_climbStairs1(t *testing.T) {
	type args struct {
		n int
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
				n: 2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				n: 3,
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				n: 4,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "4",
			args: args{
				n: 20,
			},
			expected: expected{
				result: 10946,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs1(tt.args.n); got != tt.expected.result {
				t.Errorf("climbStairs1() = %v, want %v", got, tt.expected.result)
			}
		})
	}
}

func Test_climbStairs2(t *testing.T) {
	type args struct {
		n int
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
				n: 2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				n: 3,
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				n: 4,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "4",
			args: args{
				n: 20,
			},
			expected: expected{
				result: 10946,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); got != tt.expected.result {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.expected.result)
			}
		})
	}
}

func Test_climbStairs3(t *testing.T) {
	type args struct {
		n int
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
				n: 2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				n: 3,
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				n: 4,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "4",
			args: args{
				n: 20,
			},
			expected: expected{
				result: 10946,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3(tt.args.n); got != tt.expected.result {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.expected.result)
			}
		})
	}
}

// benchmark
func Benchmark_climbStairs1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs1(20)
	}
}

func Benchmark_climbStairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs2(20)
	}
}

func Benchmark_climbStairs3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs3(20)
	}
}