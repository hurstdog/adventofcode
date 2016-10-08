// Wires is a package for processing wire schema from day 7
package wires

import (
	"fmt"
	"strconv"
	"strings"
)

const MASK16 = 0xFFFF

// Types of connections:
// const -> x
// x [AND|OR|LSHIFT|RSHIFT] y -> z
// [NOT] x -> y

// Need to be able to parse a line and turn that into a set of up to two ops
// Need a lookup table of each op. map[string]int
// We'll update each value as we process through the file.

// Circuit
var C map[string]int = make(map[string]int)

// A common signature for all line processing functions
type linefunc func(string) error

func RunLine(line string) error {
	var err error
	// num tokens is the number of spaces + 1
	tokl := strings.Count(line, " ") + 1
	if tokl < 3 {
		return fmt.Errorf("Line [%v] has less than 3 tokens.\n", line)
	} else if tokl == 3 {
		return handleAssignment(line)
	} else if tokl == 4 {
		return handleNot(line)
	} else {
		return fmt.Errorf("Line [%v] has too many tokens.\n", line)
	}

	return err
}

// Processes the case of "NOT x -> y", returning an error on malformed input.
func handleNot(line string) error {
	tok := strings.Split(line, " ")
	if len(tok) != 4 {
		return fmt.Errorf("Too many tokens, expected 4, got %v, from [%v].\n",
			len(tok), strings.Join(tok, " "))
	}
	if tok[0] != "NOT" {
		return fmt.Errorf("Expected /^NOT.*/ with 4 tokens, got /%v/\n", line)
	}
	k := tok[1]
	v, ok := C[k]
	if !ok {
		return fmt.Errorf("Token %v not yet defined, used in [%v]\n", k, line)
	}
	if tok[2] != "->" {
		return fmt.Errorf("4 tokens implies '->' is the third, it isn't: [%v].\n",
			strings.Join(tok, " "))
	}
	id := tok[3]
	_, ok = C[id]
	if ok {
		return fmt.Errorf("Line [%v] assigns to a token that already exists.\n",
			strings.Join(tok, " "))
	}
	C[id] = ^v & MASK16
	return nil
}

// Processes the case of "N -> x", returning an error on malformed input.
func handleAssignment(line string) error {
	tok := strings.Split(line, " ")
	if len(tok) != 3 {
		return fmt.Errorf("Too many tokens, expected 3, got %v, from [%v].\n",
			len(tok), strings.Join(tok, " "))
	}
	if tok[1] != "->" {
		return fmt.Errorf("3 tokens implies '->' is the second, it isn't: [%v].\n",
			strings.Join(tok, " "))
	}
	val, err := strconv.Atoi(tok[0])
	if err != nil {
		return err
	}
	id := tok[2]
	_, ok := C[id]
	if ok {
		return fmt.Errorf("Line [%v] assigns to a token that already exists.\n",
			strings.Join(tok, " "))
	}
	C[id] = val
	return nil
}
