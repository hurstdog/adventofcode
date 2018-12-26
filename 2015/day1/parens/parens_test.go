// Tests for the parens package.
package parens

import "testing"

var countTable map[string]int = make(map[string]int)
var posTable map[string]int = make(map[string]int)

func init() {
	countTable["(())"] = 0
	countTable["()()"] = 0
	countTable["((("] = 3
	countTable["(()(()("] = 3
	countTable["))((((("] = 3
	countTable["))("] = -1
	countTable["())"] = -1
	countTable[")))"] = -3
	countTable[")())())"] = -3

	posTable[")"] = 1
	posTable["()())"] = 5
}

func TestCount(t *testing.T) {
	for s, c := range countTable {
		if Count([]byte(s)) != c {
			t.Errorf("count %s: Got %d, expected %d\n", s, Count([]byte(s)), c)
		}
	}
}

func TestPosition(t *testing.T) {
	for s, c := range posTable {
		if Position([]byte(s)) != c {
			t.Errorf("pos %s: Got %d, expected %d\n", s, Count([]byte(s)), c)
		}
	}
}
