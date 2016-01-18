// Tests for the Lights package
package lights

import "testing"

var testData map[string]Command = make(map[string]Command)

func init() {
	testData["turn on 0,0 through 999,999"] =
		Command{ON, Point{0, 0, OFF}, Point{999, 999, OFF}}
	testData["turn off 0,0 through 999,999"] =
		Command{OFF, Point{0, 0, OFF}, Point{999, 999, OFF}}
	testData["toggle 0,0 through 999,999"] =
		Command{TOGGLE, Point{0, 0, OFF}, Point{999, 999, OFF}}
	testData["toggle 888,888 through 888,888"] =
		Command{TOGGLE, Point{888, 888, OFF}, Point{888, 888, OFF}}
}

func TestSetLight(t *testing.T) {
	p := Point{0, 0, OFF}
	v, err := getLight(p)
	expectValue(v, OFF, err, t)

	// OFF -> ON
	err = setLight(p, ON)
	v, err = getLight(p)
	expectValue(v, ON, err, t)

	// ON -> OFF
	err = setLight(p, OFF)
	v, err = getLight(p)
	expectValue(v, OFF, err, t)

	// OFF -> ON
	err = setLight(p, TOGGLE)
	v, err = getLight(p)
	expectValue(v, ON, err, t)

	// ON -> OFF
	err = setLight(p, TOGGLE)
	v, err = getLight(p)
	expectValue(v, OFF, err, t)
}

func expectValue(v int, exp int, err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expecting %v value == %v, got error: %v\n", v, exp, err)
	}
	if v != exp {
		t.Errorf("Point %v, expected value %v\n", v, exp)
	}
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
