package day03

import (
	"regexp"
	"strconv"
)

var Parser = regexp.MustCompile(`mul\((?P<multiplicand>\d+),(?P<multiplier>\d+)\)`)

func CorruptMemoryComputer(input string) int {
	matches := Parser.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		return 0
	}

	// Build up a 2-dimensional slice where each slice contains the
	// multipler and the multiplicand
	tokens := make([][2]int, len(matches))
	for i, match := range matches {
		multiplier, _ := strconv.Atoi(match[1])
		multiplicand, _ := strconv.Atoi(match[2])
		tokens[i] = [2]int{multiplier, multiplicand}
	}
	return sumTokensMultiplier(tokens, 0)
}

func sumTokensMultiplier(tokens [][2]int, acc int) int {
	if len(tokens) == 0 {
		return acc
	}
	return sumTokensMultiplier(tokens[1:], acc+(tokens[0][0]*tokens[0][1]))
}
