// Tests for the Wires package
package wires

import (
	"fmt"
	"testing"
)

var testFileR []string = make([]string, 7)

func init() {
	testFileR = []string{
		"y RSHIFT 2 -> g",
		"x LSHIFT 2 -> f",
		"x OR y -> e",
		"x AND y -> d",
		"456 -> y",
		"123 -> x",
	}
}

func TestResetC(t *testing.T) {
	Load(testFileR)
	if Input[3] != testFileR[3] {
		t.Errorf("Load failed: index 3 mismatch: [%v] != [%v]\n",
			Input[3], testFileR[3])
	}
	Reset()
	if len(Input) != 0 {
		t.Errorf("Reset failed: expected length 0, got %v\n", len(Input))
	}
}

func TestDefineValue(t *testing.T) {
	Load(testFileR)
	err := DefineValue("g")
	if err != nil {
		t.Error(err)
	}
	if C["g"] != 114 {
		t.Error("Expected g=114, got x=%v\n", C["x"])
	}
}

func TestHandleAssignment(t *testing.T) {
	Load(testFileR)
	expectValue(handleAssignment, "123 -> x", "x", 123, t)
	exp := fmt.Errorf("Line [123 -> x] assigns to a token that already exists.\n")
	expectErr(handleAssignment, "123 -> x", exp, t)
	expectValue(handleAssignment, "x -> y", "y", 123, t)
}

func TestHandleNot(t *testing.T) {
	Load(testFileR)

	exp := fmt.Errorf("Expected /^NOT.*/ with 4 tokens, got /ASD x -> h/\n")
	expectErr(handleNot, "ASD x -> h", exp, t)

	handleAssignment("123 -> x")
	handleAssignment("123 -> y")
	exp = fmt.Errorf("Line [NOT x -> y] assigns to a token that already exists.\n")
	expectErr(handleNot, "NOT x -> y", exp, t)

	expectValue(handleNot, "NOT x -> h", "h", 65412, t)
}

func TestHandleOpErrors(t *testing.T) {
	Load(testFileR)

	e := "Expected op AND|OR|LSHIFT|RSHIFT with 4 tokens, got [x A y -> z]\n"
	exp := fmt.Errorf(e)
	expectErr(handleOp, "x A y -> z", exp, t)

	handleAssignment("123 -> x")
	handleAssignment("123 -> y")
	handleAssignment("123 -> k")
	exp = fmt.Errorf("Line [x AND y -> k] assigns to a token that already exists.\n")
	expectErr(handleOp, "x AND y -> k", exp, t)
}

func TestHandleOps(t *testing.T) {
	Load(testFileR)

	handleAssignment("123 -> x")
	handleAssignment("456 -> y")
	expectValue(handleOp, "x AND y -> d", "d", 72, t)
	expectValue(handleOp, "x OR y -> e", "e", 507, t)
	expectValue(handleOp, "x LSHIFT 2 -> f", "f", 492, t)
	expectValue(handleOp, "y RSHIFT 2 -> g", "g", 114, t)
}

func expectValue(fn linefunc, s string, k string, v uint16, t *testing.T) {
	err := fn(s)
	if err != nil {
		t.Error(err)
	}
	if C[k] != v {
		t.Errorf("Expecting %v = %v, got %v\n", k, v, C[k])
	}
}

func expectErr(fn linefunc, s string, e error, t *testing.T) {
	err := fn(s)
	if err.Error() != e.Error() {
		t.Errorf("Expected error [%v], got error [%v]\n", e, err)
	}
}
