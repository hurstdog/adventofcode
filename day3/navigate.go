// calc_wrapping
// http://adventofcode.com/day/2
package main

import (
	"fmt"
	"github.com/hurstdog/adventofcode/day2/wrapping"
	"io/ioutil"
	"os"
	"strings"
)

const INPUT = "input.txt"

func main() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", INPUT, err)
		os.Exit(1)
	}
	var total int
	var ribbon int
	for _, line := range strings.Split(string(data), "\n") {
		cur, err := wrapping.PaperNeeded(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line %s: %v\n", line, err)
			continue
		}
		total += cur
		rib, err := wrapping.RibbonNeeded(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line %s: %v\n", line, err)
			continue
		}
		ribbon += rib
	}
	fmt.Printf("Total wrapping paper needed: %d sq ft\n", total)
	fmt.Printf("Total ribbon needed: %d ft\n", ribbon)
}
