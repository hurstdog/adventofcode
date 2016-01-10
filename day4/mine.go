// mine will run through all md5 hashes created by the string bgvyzdsv with
// trailing numbers to find the lowest number that provides an md5 with 5
// leading 0's.
package main

import (
	"crypto/md5"
	"fmt"
)

const INPUT = "bgvyzdsv"
const MAXNUM = 1 << 31
const PREFIX = "000000"

func main() {
	for i := 0; i < MAXNUM; i++ {
		t := fmt.Sprintf("%s%d", INPUT, i)
		sum := md5.Sum([]byte(t))
		p := string(sum[0:5])
		if p == PREFIX {
			fmt.Printf("%d: md5(%v) = %v\n", i, t, sum)
			break
		}
	}
}
