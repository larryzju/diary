package sample

import (
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	a := [6]int{1, 2, 3}
	b := [3]int{2, 5, 6}
	MergeSortArray(a[:], b[:], 3, 3)
	if a != [6]int{1, 2, 2, 3, 5, 6} {
		t.Error(a)
	}
}
