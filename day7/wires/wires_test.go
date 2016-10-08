// Tests for the Wires package
package wires

import (
	"fmt"
	"testing"
)

func TestRunLine(t *testing.T) {
	expectValue("123 -> x", "x", 123, t)
	exp := fmt.Errorf("Line [123 -> x] assigns to a token that already exists.\n")
	expectErr("123 -> x", exp, t)
}

func expectValue(s string, k string, v int, t *testing.T) {
	err := RunLine(s)
	if err != nil {
		t.Error(err)
	}
	if C[k] != v {
		t.Errorf("Expecting %v = %v, got %v\n", k, v, C[k])
	}
}

func expectErr(s string, e error, t *testing.T) {
	err := RunLine(s)
	if err.Error() != e.Error() {
		t.Errorf("Expected error [%v], got error [%v]\n", e, err)
	}
}
