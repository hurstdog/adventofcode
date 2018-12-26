// navigate
// http://adventofcode.com/day/3
package main

import (
	"fmt"
	"github.com/hurstdog/adventofcode/day3/nav"
	"io/ioutil"
	"os"
)

const INPUT = "input.txt"

func main() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", INPUT, err)
		os.Exit(1)
	}
	dirs := string(data)
	c, err := nav.AtLeastOne(dirs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing directions: %v", err)
	}
	fmt.Printf("Total points visited: %d\n", c)

	c, err = nav.AtLeastOneRobo(dirs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing directions: %v", err)
	}
	fmt.Printf("Total robo points visited: %d\n", c)
}
