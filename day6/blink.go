// Blink follows the instructions in a given input file to turn on and off a
// grid of lights.
package main

import (
	"bufio"
	"fmt"
	"github.com/hurstdog/adventofcode/day6/lights"
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
	c := 0
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		cmd, err := lights.LineToCmd(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing line [%s]: %v\n", line, err)
		}
		fmt.Printf("%s -> %v\n", line, cmd)
		c++
		if c > 10 {
			os.Exit(0)
		}
	}
}
