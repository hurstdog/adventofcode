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

// Processes the case of "x OP y -> z", returning an error on malformed input.
func handleOp(line string) error {
	tok := strings.Split(line, " ")
	if len(tok) != 5 {
		return fmt.Errorf("Bad token count: expected 5, got %v, from [%v].\n",
			len(tok), strings.Join(tok, " "))
	}
	op := tok[1]
	if op != "OR" && op != "AND" && op != "RSHIFT" && op != "LSHIFT" {
		e := "Expected op AND|OR|LSHIFT|RSHIFT with 4 tokens, got [%v]\n"
		return fmt.Errorf(e, line)
	}
	x := tok[0]
	_, ok := C[x]
	if !ok {
		return fmt.Errorf("Token %v not yet defined, used in [%v]\n", x, line)
	}
	y := tok[2]
	_, ok = C[y]
	if !ok {
		return fmt.Errorf("Token %v not yet defined, used in [%v]\n", y, line)
	}
	z := tok[4]
	_, ok = C[z]
	if ok {
		return fmt.Errorf("Line [%v] assigns to a token that already exists.\n", line)
	}

	// Now we get to actually doing the calculations
	if op == "AND" {
		C[z] = C[x] & C[y]
	} else if op == "OR" {
		C[z] = C[x] | C[y]
	}
	return nil
}

// Processes the case of "NOT x -> y", returning an error on malformed input.
func handleNot(line string) error {
	tok := strings.Split(line, " ")
	if len(tok) != 4 {
		return fmt.Errorf("Bad token count: expected 4, got %v, from [%v].\n",
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
		return fmt.Errorf("Bad token count: expected 3, got %v, from [%v].\n",
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
