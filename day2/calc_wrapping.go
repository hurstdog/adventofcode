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
	for _, line := range strings.Split(string(data), "\n") {
		cur, err := wrapping.PaperNeeded(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line %s: %v\n", line, err)
			continue
		}
		fmt.Printf("%s -> %d. cur == %d\n", line, cur, total)
		total += cur
	}
	fmt.Printf("Total wrapping paper needed: %d sq ft\n", total)
}
