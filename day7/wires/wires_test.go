// Tests for the Wires package
package wires

import (
	"fmt"
	"testing"
)

func resetC() {
	for k, _ := range C {
		delete(C, k)
	}
}

func TestResetC(t *testing.T) {
	expectValue(RunLine, "123 -> x", "x", 123, t)
	resetC()
	expectValue(RunLine, "123 -> x", "x", 123, t)
}

func TestHandleAssignment(t *testing.T) {
	resetC()
	expectValue(handleAssignment, "123 -> x", "x", 123, t)
	exp := fmt.Errorf("Line [123 -> x] assigns to a token that already exists.\n")
	expectErr(handleAssignment, "123 -> x", exp, t)
}

func TestHandleNot(t *testing.T) {
	resetC()

	exp := fmt.Errorf("Expected /^NOT.*/ with 4 tokens, got /ASD x -> h/\n")
	expectErr(handleNot, "ASD x -> h", exp, t)

	exp = fmt.Errorf("Token x not yet defined, used in [NOT x -> h]\n")
	expectErr(handleNot, "NOT x -> h", exp, t)

	handleAssignment("123 -> x")
	handleAssignment("123 -> y")
	exp = fmt.Errorf("Line [NOT x -> y] assigns to a token that already exists.\n")
	expectErr(handleNot, "NOT x -> y", exp, t)

	expectValue(handleNot, "NOT x -> h", "h", 65412, t)
}

func expectValue(fn linefunc, s string, k string, v int, t *testing.T) {
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
