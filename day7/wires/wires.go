// Wires is a package for processing wire schema from day 7
package wires

import (
	"fmt"
	"strconv"
	"strings"
)

// Types of connections:
// const -> x
// x [AND|OR|LSHIFT|RSHIFT] y -> z
// [NOT] x -> y

// Need to be able to parse a line and turn that into a set of up to two ops
// Need a lookup table of each op. map[string]int
// We'll update each value as we process through the file.

// Circuit
var C map[string]int = make(map[string]int)

func RunLine(line string) error {
	var err error
	// Tokenize
	tok := strings.Split(line, " ")
	tokl := len(tok)
	if tokl < 3 {
		return fmt.Errorf("Line [%v] has less than 3 tokens.\n", line)
	} else if tokl == 3 {
		if tok[1] != "->" {
			return fmt.Errorf("3 tokens implies 2 == ->, it doesn't: [%v].\n", line)
		}
		// const -> x
		val, err := strconv.Atoi(tok[0])
		if err != nil {
			return err
		}
		id := tok[2]
		_, ok := C[id]
		if ok {
			return fmt.Errorf("Line [%v] assigns to a token that already exists.\n", line)
		}
		C[id] = val
	} else {
		return fmt.Errorf("Line [%v] has too many tokens.\n", line)
	}

	return err
}
