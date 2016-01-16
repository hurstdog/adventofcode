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
	buf := bufio.NewScanner(f)
	c := 0
	for {
		if !buf.Scan() {
			break
		}
		t := buf.Text()
		fmt.Printf("Read line: [%s]\n", string(t))
		c++
		if c > 10 {
			os.Exit(0)
		}
	}
}
