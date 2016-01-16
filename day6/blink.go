// Blink follows the instructions in a given input file to turn on and off a
// grid of lights.
package main

import (
	"bufio"
	"fmt"
	"os"
)

const INPUT = "input.txt"

func main() {
	f, err := os.Open(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", INPUT, err)
		os.Exit(1)
	}
	buf := bufio.NewReader(f)
	c := 0
	for line := range buf.ReadLine() {
		fmt.Printf("Read line: [%s]\n", string(line))
		c++
		if c > 10 {
			os.Exit(0)
		}
	}
}
