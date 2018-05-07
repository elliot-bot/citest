package view

import (
	"testing"

	"citest/view/util"
)

func TestFindSubStr(t *testing.T) {
	var tests = []struct {
		str    string
		sbuStr string
		length int
	}{
		{"aabcceddabc", "dabc", 4},
		{"aabcceddabccabcde", "cabcde", 6},
	}

	for _, test := range tests {
		if subStr, length := util.FindSubStr(test.str); test.str != subStr && test.length != length {
			t.Errorf("want %s, %d, but get %s, %d", test.str, test.length, subStr, length)
		}
	}

}

func TestMergeIntSclice(t *testing.T) {
	var tests = []struct {
		a1, a2 []int
		merged []int
	}{
		{a1: []int{1, 3, 4}, a2: []int{1, 1, 2, 5}, merged: []int{1, 1, 1, 2, 3, 4, 5}},
		{a1: []int{7, 2, 1}, a2: []int{5, 3, 4, 0}, merged: []int{0, 1, 2, 3, 4, 5, 7}},
	}

	for _, test := range tests {
		if merged := util.MergeIntSclice(test.a1, test.a2); !IntSliceEqualBCE(merged, test.merged) {
			t.Error("want:", test.merged, " but was: ", merged)
		}
	}
}

func TestIsAbleToLast(t *testing.T) {
	var tests = []struct {
		a        []int
		isAbleTo bool
	}{
		{[]int{1, 4, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{nil, true},
		{[]int{0}, true},
	}

	for _, test := range tests {
		if isAbleTo := util.IsAbleToLast(test.a); isAbleTo != test.isAbleTo {
			t.Error(test, "want: ", test.isAbleTo, " but was: ", isAbleTo)
		}
	}
}

func IntSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
