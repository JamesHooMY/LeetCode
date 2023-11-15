package easy

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/chaochen1407/article/details/81677123?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-81677123-blog-96206568.235%5Ev38%5Epc_relevant_anti_t3_base&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-81677123-blog-96206568.235%5Ev38%5Epc_relevant_anti_t3_base&utm_relevant_index=2

type codec struct{}

type Codec interface {
	Encode(strs []string) string
	Decode(str string) []string
}

func Constructor() Codec {
	return &codec{}
}

// Encodes a list of strings to a single string.
func (this *codec) Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		// result += fmt.Sprintf("%d#", len(str)) + str
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

// Decodes a single string to a list of strings.
func (this *codec) Decode(str string) []string {
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

func Test_Encode(t *testing.T) {
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
			Constructor().Encode(tc.args.strs),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_Decode(t *testing.T) {
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
			Constructor().Decode(tc.args.str),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
