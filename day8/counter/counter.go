// Counter is a package for solving day8 of 2015 adventofcode.
/*
    Basically, this package contains a function for working over a stream of
    character and keeping two counters: one of the total number of memory
    characters seen (a count of the raw string seen), and one for the number of
    actual characters seen (skipping escape values).

    Example:
    Sequence    Literal Chars   Mem chars
    ""          2               0
    "abc"       5               3
    "aaa\"aaa"  10              7
    "\x27"      6               1
*/
package counter

import "fmt"

type Count struct {
  literal int
  mem int
}

var seenCount Count

// Reset clears the counts.
func Reset() {
  seenCount.literal = 0
  seenCount.mem = 0
}

// Literals returns the number of literal characters seen.
func Literals() int {
  return seenCount.literal
}

// Mem returns the number of memory characters seen.
func Mem() int {
  return seenCount.mem
}

// AddLine processes a single line (ending in newline) and adds its char counts.
func AddLine(line string) error {
  var l, m int

  var backslash, hex bool
  hexCount := 0
  for i := 0; i < len(line); i++ {
    l++
    c := string(line[i])
    if hex {
      if backslash {
        return fmt.Errorf("[%v](char %v) Didn't expect \\ when parsing hex.\n", line, i)
      }
      hexCount--
      if hexCount == 0 {
        hex = false
        m++
      }
    } else if backslash {
      backslash = false
      if c == "x" {
        hex = true
        hexCount = 2
      } else if c == "\\" || c == "\"" {
        m++
      } else {
        return fmt.Errorf("Expected \\, \", or x after backslash, got '%v'\n", c)
      }
    } else  {
      // Normal character, possibly a new backslash escape.
      if c == "\\" {
        backslash = true
      } else if c != "\"" {
        // Skip "
        m++
      }
    }
  }

  seenCount.literal = seenCount.literal + l
  seenCount.mem = seenCount.mem + m
  return nil
}
