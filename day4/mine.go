// mine will run through all md5 hashes created by the string bgvyzdsv with
// trailing numbers to find the lowest number that provides an md5 with 5
// leading 0's.
package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const INPUT = "bgvyzdsv"
const MAXNUM = 1 << 62
const CHECK = 3
const PREFIX = "000000"

func main() {
	ch := make(chan string)
	for i := int64(0); i < MAXNUM; i++ {
		go testInt(i, ch)
	}

	// Print the first three results and quit
	for i := 0; i < CHECK; i++ {
		fmt.Println(<-ch)
	}
}

func testInt(i int64, ch chan<- string) {
	t := fmt.Sprintf("%s%d", INPUT, i)
	sum := fmt.Sprintf("%032x", md5.Sum([]byte(t)))
	if i%10000 == 0 {
		fmt.Printf("Test: %d: md5(%v) = %v\n", i, t, sum)
	}
	if strings.HasPrefix(sum, PREFIX) {
		ch <- fmt.Sprintf("%d: md5(%v) = %v\n", i, t, sum)
	}
}
