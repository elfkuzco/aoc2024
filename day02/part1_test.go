package day02

import "testing"

func TestRunningDifference(t *testing.T) {
	var tests = []struct {
		entries []int
		want    []int
	}{
		{
			[]int{7, 6, 4, 2, 1},
			[]int{1, 2, 2, 1},
		},
		{
			[]int{1, 2, 7, 8, 9},
			[]int{-1, -5, -1, -1},
		},
	}

	for _, test := range tests {
		got := runningDifference(test.entries)
		if !equal(got, test.want) {
			t.Errorf("runningDifference(%v) = %v, want %v", test.entries, got, test.want)
		}
	}
}

func TestSignsAreTheSame(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want bool
	}{
		{1, 1, true},
		{1, -1, false},
		{-2, -3, true},
	}

	for _, test := range tests {
		got := signsAreTheSame(test.a, test.b)
		if got != test.want {
			t.Errorf("signsAreTheSame(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}

func TestReportIsSafe(t *testing.T) {
	var tests = []struct {
		report []int
		want   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		got := IsSafe(test.report)
		if got != test.want {
			t.Errorf("IsSafe(%v) = %v, want %v", test.report, got, test.want)
		}
	}
}

func TestDifferencesAreSafe(t *testing.T) {
	var tests = []struct {
		differences []int
		want        bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{-1, -5, -1, -1}, false},
		{[]int{2, 1, 4, 1}, false},
		{[]int{-2, 1, -2, -1}, false},
		{[]int{2, 2, 0, 3}, false},
		{[]int{-2, -3, -1, -2}, true},
	}

	for _, test := range tests {
		got := differencesAreSafe(test.differences, 0)
		if got != test.want {
			t.Errorf("differencesAreSafe(%v) = %v, want %v", test.differences, got, test.want)
		}
	}
}
