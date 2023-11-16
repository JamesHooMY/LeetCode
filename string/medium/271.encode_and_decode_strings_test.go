package medium

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/chaochen1407/article/details/81677123?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-81677123-blog-96206568.235%5Ev38%5Epc_relevant_anti_t3_base&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-81677123-blog-96206568.235%5Ev38%5Epc_relevant_anti_t3_base&utm_relevant_index=2

// method 1 string
// encode:
// 1) use "#" to separate the length of the string and the string itself
// 2) for each string in strs, append the length of the string and the string itself to the result
// 3) finally, return the result
// TC = O(N), SC = O(N)
// decode:
// 1) use "#" to separate the length of the string and the string itself
// 2) for each iteration, get the length of the string, then get the string itself, and append the string to the result
// 3) finally, return the result
// TC = O(N), SC = O(N)
type (
	codec1 struct{}
	Codec  interface {
		Encode(strs []string) string
		Decode(str string) []string
	}
)

func Constructor1() Codec {
	return &codec1{}
}

// Encodes a list of strings to a single string.
func (this *codec1) Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		// result += fmt.Sprintf("%d#", len(str)) + str
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

// Decodes a single string to a list of strings.
func (this *codec1) Decode(str string) []string {
	result := []string{}
	for len(str) > 0 {
		index := 0
		for str[index] != '#' {
			index++
		}

		length, _ := strconv.Atoi(str[:index])
		result = append(result, str[index+1:index+1+length])
		// don't forget to update str !!!
		str = str[index+1+length:]
	}

	return result
}

// method 2 buffer bytes
// encode:
// the logic thinking is the same as method 1, but use buffer bytes to improve the performance
// TC = O(N), SC = O(N)
// decode:
// the logic thinking is the same as method 1
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
type codec2 struct{}

func Constructor2() Codec {
	return &codec2{}
}

// Encodes a list of strings to a single string.
func (this *codec2) Encode(strs []string) string {
	var buffer bytes.Buffer

	for _, str := range strs {
		buffer.WriteString(strconv.Itoa(len(str)))
		buffer.WriteString("#")
		buffer.WriteString(str)
	}

	return buffer.String()
}

// Decodes a single string to a list of strings.
func (this *codec2) Decode(str string) []string {
	result := []string{}

	for len(str) > 0 {
		index := 0
		for str[index] != '#' {
			index++
		}

		length, _ := strconv.Atoi(str[:index])
		result = append(result, str[index+1:index+1+length])
		str = str[index+1+length:]
	}

	return result
}

func Test_Constructor1_Encode(t *testing.T) {
	type args struct {
		strs []string
	}
	type expected struct {
		result string
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
				strs: []string{"hello", "world"},
			},
			expected: expected{
				result: "5#hello5#world",
			},
		},
		{
			name: "2",
			args: args{
				strs: []string{"hello", "world", "123"},
			},
			expected: expected{
				result: "5#hello5#world3#123",
			},
		},
		{
			name: "3",
			args: args{
				strs: []string{},
			},
			expected: expected{
				result: "",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			Constructor1().Encode(tc.args.strs),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_Constructor1_Decode(t *testing.T) {
	type args struct {
		str string
	}
	type expected struct {
		result []string
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
				str: "5#hello5#world",
			},
			expected: expected{
				result: []string{"hello", "world"},
			},
		},
		{
			name: "2",
			args: args{
				str: "5#hello5#world3#123",
			},
			expected: expected{
				result: []string{"hello", "world", "123"},
			},
		},
		{
			name: "3",
			args: args{
				str: "",
			},
			expected: expected{
				result: []string{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			Constructor1().Decode(tc.args.str),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_Constructor2_Encode(t *testing.T) {
	type args struct {
		strs []string
	}
	type expected struct {
		result string
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
				strs: []string{"hello", "world"},
			},
			expected: expected{
				result: "5#hello5#world",
			},
		},
		{
			name: "2",
			args: args{
				strs: []string{"hello", "world", "123"},
			},
			expected: expected{
				result: "5#hello5#world3#123",
			},
		},
		{
			name: "3",
			args: args{
				strs: []string{},
			},
			expected: expected{
				result: "",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			Constructor2().Encode(tc.args.strs),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_Constructor2_Decode(t *testing.T) {
	type args struct {
		str string
	}
	type expected struct {
		result []string
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
				str: "5#hello5#world",
			},
			expected: expected{
				result: []string{"hello", "world"},
			},
		},
		{
			name: "2",
			args: args{
				str: "5#hello5#world3#123",
			},
			expected: expected{
				result: []string{"hello", "world", "123"},
			},
		},
		{
			name: "3",
			args: args{
				str: "",
			},
			expected: expected{
				result: []string{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			Constructor2().Decode(tc.args.str),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_Constructor1_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor1().Encode([]string{"hello", "world"})
	}
}

func Benchmark_Constructor1_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor1().Decode("5#hello5#world")
	}
}

func Benchmark_Constructor2_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor2().Encode([]string{"hello", "world"})
	}
}

func Benchmark_Constructor2_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor2().Decode("5#hello5#world")
	}
}
