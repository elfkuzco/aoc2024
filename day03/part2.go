package day03

import (
	"regexp"
	"strconv"
)

var ControlledParser = regexp.MustCompile(`mul\((?P<multiplicand>\d+),(?P<multiplier>\d+)\)|don't\(\)|do\(\)`)

func ControlledCorruptMemoryComputer(input string) int {
	matches := ControlledParser.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		return 0
	}

	enabled := true

	// Build up a 2-dimensional slice where each slice contains the
	// multipler and the multiplicand
	tokens := [][2]int{}
	for _, match := range matches {
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		}
		if enabled {
			multiplier, _ := strconv.Atoi(match[1])
			multiplicand, _ := strconv.Atoi(match[2])
			tokens = append(tokens, [2]int{multiplier, multiplicand})
		}
	}
	return sumTokensMultiplier(tokens, 0)
}
