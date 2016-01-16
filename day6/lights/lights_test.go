// Tests for the Lights package
package lights

import "testing"

var testData map[string]Command = make(map[string]Command)

func init() {
	testData["turn on 0,0 through 999,999"] =
		Command{ON, Point{0, 0}, Point{999, 999}}
	testData["turn off 0,0 through 999,999"] =
		Command{OFF, Point{0, 0}, Point{999, 999}}
	testData["toggle 0,0 through 999,999"] =
		Command{TOGGLE, Point{0, 0}, Point{999, 999}}
	testData["toggle 888,888 through 888,888"] =
		Command{TOGGLE, Point{888, 888}, Point{888, 888}}
}

func TestLineToCmd(t *testing.T) {
	for cmd, exp := range testData {
		testLineToCmd(cmd, exp, t)
	}
}

func testLineToCmd(cmd string, exp Command, t *testing.T) {
	res, err := LineToCmd(cmd)
	if err != nil {
		t.Errorf("%s gave error: %v\n", cmd, err)
	}
	if res != exp {
		t.Errorf("%s: expected %v, got %v\n", cmd, exp, res)
	}
}
