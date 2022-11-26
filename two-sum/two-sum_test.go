package twosum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/two-sum/

type SortedNum struct {
	Index int
	Val   int
}

func twoSum(nums []int, target int) []int {

	snums := make([]SortedNum, len(nums))
	for i, v := range nums {
		snums[i] = SortedNum{Index: i, Val: v}
	}
	sort.Slice(snums, func(i, j int) bool {
		return snums[i].Val < snums[j].Val
	})

	largeIndex := len(snums) - 1
	smallIndex := 0
	sum := snums[largeIndex].Val + snums[smallIndex].Val
	for sum != target {
		if sum > target {
			largeIndex--
			smallIndex = 0
		} else {
			smallIndex++
		}
		sum = snums[largeIndex].Val + snums[smallIndex].Val
	}
	return []int{snums[smallIndex].Index, snums[largeIndex].Index}
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
