package thekweakestrowsinamatrix

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func kWeakestRows(mat [][]int, k int) []int {
	var rs rows
	for i, r := range mat {
		soldierCount := 0
		for _, s := range r {
			if s == 0 {
				break
			}
			soldierCount++
		}
		rsRow := row{Index: i, SoldierCount: soldierCount}
		rs = append(rs, rsRow)
	}
	sort.Sort(rs)
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = rs[i].Index
	}
	return ret
}

type row struct {
	Index        int
	SoldierCount int
}

type rows []row

func (r rows) Len() int {
	return len(r)
}

func (r rows) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r rows) Less(i, j int) bool {
	if r[i].SoldierCount == r[j].SoldierCount {
		return r[i].Index < r[j].Index
	}
	return r[i].SoldierCount < r[j].SoldierCount
}

func Test_kWeakestRows(t *testing.T) {
	tests := []struct {
		name        string
		mat         [][]int
		k           int
		expectedRes []int
	}{
		{
			name: "1",
			mat: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:           3,
			expectedRes: []int{2, 0, 3},
		},
		{

			name: "2",
			mat: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			k:           2,
			expectedRes: []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRes := kWeakestRows(tt.mat, tt.k)
			require.Equal(t, tt.expectedRes, actualRes)
		})
	}
}
