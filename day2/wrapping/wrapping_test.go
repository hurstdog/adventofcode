// Tests for the wrapping package.
package wrapping

import "testing"

var lwhTable map[string][]int = make(map[string][]int)
var areaTable map[string]int = make(map[string]int)

func init() {
	lwhTable["1x1x1"] = []int{1, 1, 1}
	lwhTable["15x127x3"] = []int{15, 127, 3}
	areaTable["1x1x1"] = 7
	areaTable["2x3x4"] = 58
	areaTable["1x1x10"] = 43
}

func TestPaperNeeded(t *testing.T) {
	for s, c := range areaTable {
		area, err := PaperNeeded(s)
		if err != nil {
			t.Errorf("PaperNeeded %s: Got error %v", s, err)
		}
		if area != c {
			t.Errorf("PaperNeeded %s: Got %v, expected %v\n", s, area, c)
		}
	}
}

func TestToLWH(t *testing.T) {
	for s, c := range lwhTable {
		lwh, err := toLWH(s)
		if err != nil {
			t.Errorf("toLWH %s: Got error %v", s, err)
		}
		if !equal(lwh, c) {
			t.Errorf("toLWH %s: Got %v, expected %v\n", s, lwh, c)
		}
	}
}

func equal(a []int, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
