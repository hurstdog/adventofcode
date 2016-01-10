// Nice reads through the input and prints the number of strings that are nice.
// Output: Nice strings: 258
package main

import (
	"bufio"
	"fmt"
	"github.com/hurstdog/adventofcode/day5/nice"
	"os"
)

const INPUT = "input.txt"

func main() {
	f, err := os.Open(INPUT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v", INPUT, err)
	}

	var lines int
	ch := make(chan bool)
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		lines++
		go func() {
			ch <- nice.Nice(buf.Text())
		}()
	}

	var c int
	for i := 0; i < lines; i++ {
		r := <-ch
		if r {
			c++
		}
	}

	fmt.Printf("Nice strings: %d\n", c)
}
