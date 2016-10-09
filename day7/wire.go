// Wire follows the instructions in a given input file to calculate a wiring
// diagram.
// Day 7 part 1 solution:
// Day 7 part 2 solution:
package main

import (
	"bufio"
	"fmt"
	"github.com/hurstdog/adventofcode/day7/wires"
	"os"
)

const INPUT = "input.txt"

func main() {
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
		wires.AddLine(buf.Text())
	}
	err = wires.DefineValue("a")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	} else {
		fmt.Printf("After all commands, wire 'a' gets value %v.\n", wires.C["a"])
	}
}
