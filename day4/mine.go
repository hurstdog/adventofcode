// mine will run through all md5 hashes created by the string bgvyzdsv with
// trailing numbers to find the lowest number that provides an md5 with 5
// leading 0's.
// 5 0's answer: 254575. md5(bgvyzdsv254575) = 000004b30d481662b9cb0c105f6549b2
// 6 0's answer: 1038736. md5(bgvyzdsv1038736) = 000000b1b64bf5eb55aad89986126953
package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const INPUT = "bgvyzdsv"
const MAXNUM = 1 << 62
const PREFIX = "00000"

func main() {
	for i := int64(0); i < MAXNUM; i++ {
		r := testInt(i)
		if r != "" {
			fmt.Println(r)
			break
		}
	}
}

func testInt(i int64) string {
	t := fmt.Sprintf("%s%d", INPUT, i)
	sum := fmt.Sprintf("%032x", md5.Sum([]byte(t)))
	if strings.HasPrefix(sum, PREFIX) {
		return fmt.Sprintf("%d: md5(%v) = %v\n", i, t, sum)
	}
	return ""
}
