package day02

import "testing"

func TestReportIsSafeWithDampener(t *testing.T) {
	var tests = []struct {
		report []int
		want   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		got := IsSafeWithDampener(test.report)
		if got != test.want {
			t.Errorf("IsSafeWithDampener(%v) = %v, want %v", test.report, got, test.want)
		}
	}
}
