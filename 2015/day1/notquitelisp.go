// notquitelisp
// http://adventofcode.com/day/1
package main

import (
	"fmt"
	"github.com/hurstdog/adventofcode/2015/day1/parens"
	"io/ioutil"
	"os"
)

const INPUT = "input.txt"

func main() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", INPUT, err)
	}
	fmt.Println(parens.Position(data))
}
