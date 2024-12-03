package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/elfkzuco/aoc2024/day01"
)

func readInput(filename string) ([]int, []int, error) {
	var err error

	fh, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	a, b := []int{}, []int{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			return nil, nil, fmt.Errorf("each line must contain 2 fields, got %d instead", len(line))
		}
		l1, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, nil, err
		}
		l2, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, nil, err
		}
		a = append(a, l1)
		b = append(b, l2)
	}

	if err = scanner.Err(); err != nil {
		return nil, nil, err
	}
	return a, b, nil
}

func main() {
	var filename string
	flag.StringVar(&filename, "i", "input.txt", "Input filename")
	flag.Parse()

	a, b, err := readInput(filename)
	if err != nil {
		log.Fatalf("error while reading input from %s: %v\n", filename, err)
	}
	fmt.Printf("Distance: %d\n", day01.Distance(slices.Clone(a), slices.Clone(b)))
	fmt.Printf("Similarity Score: %d\n", day01.SimilarityScore(a, b))

}
