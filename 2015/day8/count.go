// Count counts the characters in the day8 puzzle.
// Day 8 part 1: Literals: 6202, Memory: 4860; difference: 1342
// Day 8 part 2: 
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hurstdog/adventofcode/day8/counter"
	"os"
)

const INPUT = "input.txt"

func main() {
	flag.Parse()
	f, err := os.Open(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", INPUT, err)
		os.Exit(1)
	}
	buf := bufio.NewScanner(f)
	for {
		if !buf.Scan() {
			break
		}
		err := counter.AddLine(buf.Text())
    if err != nil {
      fmt.Printf("ERROR: %v", err)
      break
    }
	}
  fmt.Printf("Literals: %v, Memory: %v\n", counter.Literals(), counter.Mem())
}
