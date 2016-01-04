// Tests for the nav package.
package nav

import (
	"testing"
)

var locTable map[string][]int = make(map[string][]int)

func TestUpdateLoc(t *testing.T) {
	var p Point // 0, 0
	p = testUpdate(p, ">", Point{1, 0}, t)
	p = testUpdate(p, ">", Point{2, 0}, t)
	p = testUpdate(p, "^", Point{2, 1}, t)
	p = testUpdate(p, "<", Point{1, 1}, t)
	p = testUpdate(p, "v", Point{1, 0}, t)
	p = testUpdate(p, "v", Point{1, -1}, t)
}

func testUpdate(p Point, d string, e Point, t *testing.T) Point {
	res, err := updateLoc(p, d)
	if err != nil {
		t.Errorf("%v", err)
	}
	if res.x != e.x || res.y != e.y {
		t.Errorf("updateLoc(%v, %s), expected %v, got %v\n", p, d, e, res)
	}
	return res
}
