package day03

import "testing"

func TestCorruptMemoryComputer(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
	}

	for _, test := range tests {
		got := CorruptMemoryComputer(test.input)
		if got != test.want {
			t.Errorf("CorruptMemoryComputer(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
