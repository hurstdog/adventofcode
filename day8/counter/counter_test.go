// Tests for the Counter package
package counter

import (
	"testing"
)

func TestAll(t *testing.T) {
  var test string
  Reset()
  expectEmpty(t)
  test = "\"\\x27\""
  expectAndReset(test, test, 6, 1, t)
  test = "\"\\x27\\x27\""
  expectAndReset(test, test, 10, 2, t)
  test = ""
  expectAndReset(test, test, 0, 0, t)
  test = "\"\""
  expectAndReset(test, test, 2, 0, t)
  test = "\"abc\""
  expectAndReset(test, test, 5, 3, t)
  test = "\"aaa\\\"aaa\""
  expectAndReset(test, test, 10, 7, t)
}

func expectEmpty(t *testing.T) {
  expect("expectNull", "", 0, 0, t)
}

func expectAndReset(desc string, s string, l int, m int, t *testing.T) {
  expect(desc, s, l, m, t)
  Reset()
}

func expect(desc string, s string, l int, m int, t *testing.T) {
  AddLine(s)
	if Literals() != l {
    t.Errorf("%v: Expected Literals() == %v, got %v\n", desc, l, Literals())
  }
	if Mem() != m {
    t.Errorf("%v: Expected Mem() == %v, got %v\n", desc, m, Mem())
  }
}

