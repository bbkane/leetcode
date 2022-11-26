package twosum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/two-sum/
// I ended up looking this up. Turns out hashmaps are the way to go :)

func twoSum(nums []int, target int) []int {

	addendToIndex := make(map[int]int)

	for i, addend := range nums {
		otherAddend := target - addend
		if otherIndex, exists := addendToIndex[otherAddend]; exists {
			return []int{otherIndex, i}
		}
		addendToIndex[addend] = i
	}
	panic("should have reached a solution by now")
	// return nil
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "negative",
			args: args{
				nums:   []int{-1, -2, -3, -4, -5},
				target: -8,
			},
			expected: []int{2, 4},
		},
		{
			name: "time",
			args: args{
				nums:   []int{3, 2, 3},
				target: 6,
			},
			expected: []int{0, 2},
		},
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.args.nums, tt.args.target)
			require.Equal(t, tt.expected, actual)
		})
	}
}
