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
		line := buf.Text()
		err := wires.RunLine(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing line [%s]: %v\n", line, err)
		}
	}
	a := wires.C["a"]
	fmt.Printf("After all commands, wire 'a' gets value %v.\n", a)
}
