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

func TestHandleDirections(t *testing.T) {
	ResetPoints()
	err := handleDirections(">>^<vv")
	if err != nil {
		t.Errorf("%v", err)
	}
	checkPointVal(Point{0, 0}, 1, t)
	checkPointVal(Point{1, 0}, 2, t)
	checkPointVal(Point{2, 0}, 1, t)
	checkPointVal(Point{2, 1}, 1, t)
	checkPointVal(Point{1, 1}, 1, t)
	checkPointVal(Point{1, -1}, 1, t)
}

func checkPointVal(p Point, v int, t *testing.T) {
	if pointsSeen[p] != v {
		t.Errorf("Point %v count: Expected %d, got %d\n", p, v, pointsSeen[p])
	}
}

func TestAtLeastOne(t *testing.T) {
	c, err := AtLeastOne("")
	if err != nil {
		t.Errorf("%v", err)
	}
	if c != 1 {
		t.Errorf("Base case: expected 1, got %d", c)
	}
	c, err = AtLeastOne(">>^<vv")
	if err != nil {
		t.Errorf("%v", err)
	}
	if c != 6 {
		t.Errorf("Expected count 6, got count %d", c)
	}
}

func TestAtLeastOneRobo(t *testing.T) {
	testRoboCount("", 1, t)
	testRoboCount("^v", 3, t)
	testRoboCount("^>v<", 3, t)
	testRoboCount("^v^v^v^v^v", 11, t)
}

func testRoboCount(dirs string, expected int, t *testing.T) {
	c, err := AtLeastOneRobo(dirs)
	if err != nil {
		t.Errorf("%v", err)
	}
	if c != expected {
		t.Errorf("Expected count %d, got count %d", expected, c)
	}
}
