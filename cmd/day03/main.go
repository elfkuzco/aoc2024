package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/elfkzuco/aoc2024/day03"
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

	contents, err := io.ReadAll(fh)
	if err != nil {
		log.Fatal(err)
	}
	strContents := string(contents)

	sum := day03.CorruptMemoryComputer(strContents)
	controlledSum := day03.ControlledCorruptMemoryComputer(strContents)

	fmt.Printf("CorruptMemoryComputer Sum: %d\n", sum)
	fmt.Printf("ControlledCorruptMemoryComputer Sum: %d\n", controlledSum)

}
