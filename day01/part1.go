package day01

import (
	"cmp"
	"slices"
)

// Compute the sum of the difference between elements of a and b.
// The slices are sorted in place before summing the differences of a[i] and b[i].
// Assumes that both slices are of equal length.
func Distance(a, b []int) int {
	cmpFunc := func(a, b int) int {
		return cmp.Compare(a, b)
	}
	slices.SortFunc(a, cmpFunc)
	slices.SortFunc(b, cmpFunc)
	return computeDistance(a, b, 0)
}

func computeDistance(a, b []int, acc int) int {
	if len(a) == 0 && len(b) == 0 {
		return acc
	}
	diff := 0
	if a[0] > b[0] {
		diff = a[0] - b[0]
	} else {
		diff = b[0] - a[0]
	}
	return computeDistance(a[1:], b[1:], acc+diff)
}
