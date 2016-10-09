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

// Circuit. var -> value
var C map[string]uint16 = make(map[string]uint16)

// Input file, kept around for reference.
var Input []string

// Lines that have been seen, to prevent double-processing.
var Seen map[string]bool = make(map[string]bool)

// A common signature for all line processing functions
type linefunc func(string) error

// Clears the Circuit, Seen, and input File
func Reset() {
	for k, _ := range C {
		delete(C, k)
	}
	for k, _ := range Seen {
		delete(Seen, k)
	}
	Input = nil
}

// Loads the given slice into our tracker slice.
func Load(lines []string) {
	Reset()
	Input = make([]string, len(lines))
	copy(Input, lines)
}

// Adds a single line into our tracker slice
func AddLine(line string) {
	Input = append(Input, line)
}

// This loops over all of the lines in the file and processes the line that
// defines the variable in x next.
// Coupled with skipping lines we've already processed, this allows us to go
// through the lines in any order and effectively depth-first search for only
// the values we need.
func DefineValue(x string) error {
	match := "-> " + x
	for _, v := range Input {
		if strings.HasSuffix(v, match) {
			//fmt.Printf("Running line [%v] to define %v.\n", v, x)
			err := RunLine(v)
			if err != nil {
				return err
			}
			//fmt.Printf("%v = %v.\n", x, C[x])
			break
		}
	}
	return nil
}

func RunLine(line string) error {
	// Skip lines we've already seen.
	_, ok := Seen[line]
	if ok {
		return nil
	}

	Seen[line] = true

	var err error
	// num tokens is the number of spaces + 1
	tokl := strings.Count(line, " ") + 1
	if tokl < 3 {
		return fmt.Errorf("Line [%v] has less than 3 tokens.\n", line)
	} else if tokl == 3 {
		return handleAssignment(line)
	} else if tokl == 4 {
		return handleNot(line)
	} else if tokl == 5 {
		return handleOp(line)
	} else {
		return fmt.Errorf("Line [%v] has too many tokens.\n", line)
	}

	return err
}

// Processes the case of "x OP y -> z", returning an error on malformed input.
func handleOp(line string) error {
	var err error
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
	xval, err := resolve(x)
	if err != nil {
		return err
	}
	y := tok[2]
	yval, err := resolve(y)
	if err != nil {
		return err
	}
	z := tok[4]
	_, ok := C[z]
	if ok {
		return fmt.Errorf("Line [%v] assigns to a token that already exists.\n", line)
	}

	// Now we get to actually doing the calculations
	if op == "AND" {
		C[z] = xval & yval
	} else if op == "OR" {
		C[z] = xval | yval
	} else {
		ynum, _ := strconv.Atoi(y)
		if op == "LSHIFT" {
			C[z] = uint16((C[x] << uint(ynum)) & MASK16)
		} else if op == "RSHIFT" {
			C[z] = uint16((C[x] >> uint(ynum)) & MASK16)
		}
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
	v, err := resolve(k)
	if err != nil {
		return err
	}
	if tok[2] != "->" {
		return fmt.Errorf("4 tokens implies '->' is the third, it isn't: [%v].\n",
			strings.Join(tok, " "))
	}
	id := tok[3]
	_, ok := C[id]
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
	val16 := uint16(val)
	// If it's not a number, then it's a variable. Fetch it or define it.
	if err != nil {
		fetched, ok := C[tok[0]]
		if ok {
			val16 = fetched
		} else {
			err = DefineValue(tok[0])
			if err != nil {
				return err
			}
			val16 = C[tok[0]]
		}
	}
	id := tok[2]
	_, ok := C[id]
	if ok {
		return fmt.Errorf("Line [%v] assigns to a token that already exists.\n",
			strings.Join(tok, " "))
	}
	C[id] = val16
	return nil
}

// resolve returns the string as a number, if it's a number. Otherwise it
// finds the definition of the variable and returns that.
func resolve(s string) (uint16, error) {
	// If it exists, return it
	val, ok := C[s]
	if ok {
		return val, nil
	}

	// If it's a number, return that.
	num, err := strconv.Atoi(s)
	if err == nil {
		return uint16(num), nil
	}

	// better define it
	err = DefineValue(s)
	if err != nil {
		return 0, err
	}
	return C[s], nil
}
