package day03

import "testing"

func TestControlledCorruptMemoryComputer(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
	}

	for _, test := range tests {
		got := ControlledCorruptMemoryComputer(test.input)
		if got != test.want {
			t.Errorf("ControlledCorruptMemoryComputer(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
