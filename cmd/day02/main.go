package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elfkzuco/aoc2024/day02"
)

func main() {
	var filename string
	flag.StringVar(&filename, "i", "input.txt", "Input filename")
	flag.Parse()

	var err error

	fh, err := os.Open(filename)

	if err != nil {
		log.Fatalf("error while reading input from %s: %v\n", filename, err)
	}
	defer fh.Close()

	safeReports := 0
	safeReportsWithDampener := 0

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		report := make([]int, len(line))

		for i, level := range line {
			v, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("each level must be an integer, got %s", level)
			}
			report[i] = v
		}
		if day02.IsSafe(report) {
			safeReports += 1
		}

		if day02.IsSafeWithDampener(report) {
			safeReportsWithDampener += 1
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeReports)
	fmt.Printf("Safe Reports With Dampener: %d\n", safeReportsWithDampener)

}
