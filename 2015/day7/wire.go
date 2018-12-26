// Wire follows the instructions in a given input file to calculate a wiring
// diagram.
// Day 7 part 1 solution: 3176
// Day 7 part 2 solution: 14710
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hurstdog/adventofcode/day7/wires"
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
		wires.AddLine(buf.Text())
	}
	// Part1
	err = wires.DefineValue("a")
	aval := wires.C["a"]
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	} else {
		fmt.Printf("After all commands, wire 'a' gets value %v.\n", aval)
	}
	// Part2
	wires.Reset()
	wires.C["b"] = aval // override the value to the results of part1
	err = wires.DefineValue("a")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	} else {
		fmt.Printf("For part 2, wire 'a' gets value %v.\n", wires.C["a"])
	}
}
