// Tests for the parens package.
package parens

import "testing"

var testTable map[string]int = make(map[string]int)

func init() {
	testTable["(())"] = 0
	testTable["()()"] = 0
	testTable["((("] = 3
	testTable["(()(()("] = 3
	testTable["))((((("] = 3
	testTable["))("] = -1
	testTable["())"] = -1
	testTable[")))"] = -3
	testTable[")())())"] = -3
}

func TestParens(t *testing.T) {
	for s, c := range testTable {
		if Count([]byte(s)) != c {
			t.Errorf("count %s: Got %d, expected %d\n", s, Count([]byte(s)), c)
		}
	}
}
