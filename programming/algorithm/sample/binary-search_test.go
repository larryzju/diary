package sample

import (
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	cases := []struct {
		input, expect int
	}{
		{0, 4},
		{1, 5},
		{2, 6},
		{3, -1},
		{4, 0},
		{5, 1},
		{6, 2},
		{7, 3},
		{8, -1},
	}

	for _, v := range cases {
		answer := SearchInRotatedSortedArray(nums, v.input)
		if answer != v.expect {
			t.Error(v.input, v.expect, answer)
		}
	}
}
