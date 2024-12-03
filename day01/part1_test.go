package day01

import (
	"fmt"
	"testing"
)

func TestDistanceBetweenSlices(t *testing.T) {
	var tests = []struct {
		l1   []int
		l2   []int
		want int
	}{
		{
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			11,
		},
	}

	for _, test := range tests {
		got := Distance(test.l1, test.l2)
		descr := fmt.Sprintf("Distance(%v, %v)", test.l1, test.l2)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}

}
