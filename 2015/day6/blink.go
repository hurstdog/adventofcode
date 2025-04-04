// Blink follows the instructions in a given input file to turn on and off a
// grid of lights.
// Day6 part 1 solution: 400410
// Day6 part 2 solution: 15343601
package main

import (
	"bufio"
	"fmt"
	"github.com/hurstdog/adventofcode/2015/day6/lights"
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
		cmd, err := lights.LineToCmd(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing line [%s]: %v\n", line, err)
		}
		err = lights.ApplyCmd(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error applying line [%s]: %v\n", line, err)
		}
	}
	n := lights.TotBrightness()
	fmt.Printf("After all commands, %d is the total brightness\n", n)
}
