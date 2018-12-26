// Nice reads through the input and prints the number of strings that are nice.
// Output: Nice strings: 258
package main

import (
	"bufio"
	"fmt"
	"github.com/hurstdog/adventofcode/2015/day5/nice"
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
	ch2 := make(chan bool)
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		lines++
		line := buf.Text()
		go func() {
			ch <- nice.Nice(line)
			ch2 <- nice.Nice2(line)
		}()
	}

	var c, c2 int
	for i := 0; i < lines; i++ {
		r := <-ch
		if r {
			c++
		}
		r = <-ch2
		if r {
			c2++
		}
	}

	fmt.Printf("Nice strings: %d\n", c)
	fmt.Printf("Nice2 strings: %d\n", c2)
}
